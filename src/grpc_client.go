package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/salrashid123/grpc_keylog/echo"

	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const ()

var (
	conn *grpc.ClientConn
)

func main() {

	address := flag.String("host", "localhost:50051", "host:port of gRPC server")
	insecure := flag.Bool("insecure", false, "connect without TLS")
	tlsCert := flag.String("tlsCert", "", "tls Certificate")
	serverName := flag.String("servername", "grpc.domain.com", "CACert for server")

	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "INFO")
	flag.Parse()

	var err error
	var conn *grpc.ClientConn
	if *insecure == true {
		conn, err = grpc.Dial(*address, grpc.WithInsecure())
	} else {

		var tlsCfg tls.Config
		rootCAs := x509.NewCertPool()
		pem, err := ioutil.ReadFile(*tlsCert)
		if err != nil {
			log.Fatalf("failed to load root CA certificates  error=%v", err)
		}
		if !rootCAs.AppendCertsFromPEM(pem) {
			log.Fatalf("no root CA certs parsed from file ")
		}
		tlsCfg.RootCAs = rootCAs
		tlsCfg.ServerName = *serverName

		// https://developer.mozilla.org/en-US/docs/Mozilla/Projects/NSS/Key_Log_Format
		sslKeyLogfile := os.Getenv("SSLKEYLOGFILE")
		if sslKeyLogfile != "" {
			var w *os.File
			w, err := os.OpenFile(sslKeyLogfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
			if err != nil {
				log.Fatalf("Could not create keylogger: ", err)
			}
			tlsCfg.KeyLogWriter = w
		}

		ce := credentials.NewTLS(&tlsCfg)
		conn, err = grpc.Dial(*address, grpc.WithTransportCredentials(ce))
	}
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	c := echo.NewEchoServerClient(conn)
	/// UNARY
	for i := 0; i < 5; i++ {
		r, err := c.SayHelloUnary(ctx, &echo.EchoRequest{Name: fmt.Sprintf("Unary Request %d", i)})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		time.Sleep(1 * time.Second)
		log.Printf("Unary Response: %v [%v]", i, r)
	}

	// ********CLIENT

	cstream, err := c.SayHelloClientStream(context.Background())

	if err != nil {
		log.Fatalf("%v.SayHelloClientStream(_) = _, %v", c, err)
	}

	for i := 1; i < 5; i++ {
		if err := cstream.Send(&echo.EchoRequest{Name: fmt.Sprintf("client stream RPC %d ", i)}); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("%v.Send(%v) = %v", cstream, i, err)
		}
	}

	creply, err := cstream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", cstream, err, nil)
	}
	log.Printf(" Got SayHelloClientStream  [%s]", creply.Message)

	/// ***** SERVER
	stream, err := c.SayHelloServerStream(ctx, &echo.EchoRequest{Name: "Stream RPC msg"})
	if err != nil {
		log.Fatalf("SayHelloStream(_) = _, %v", err)
	}
	for {
		m, err := stream.Recv()
		if err == io.EOF {
			t := stream.Trailer()
			log.Println("Stream Trailer: ", t)
			break
		}
		if err != nil {
			log.Fatalf("SayHelloStream(_) = _, %v", err)
		}

		log.Printf("Message: [%s]", m.Message)
	}

	/// ********** BIDI

	done := make(chan bool)
	stream, err = c.SayHelloBiDiStream(context.Background())
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}
	ctx = stream.Context()

	go func() {
		for i := 1; i <= 10; i++ {
			req := echo.EchoRequest{Name: "Bidirectional CLient RPC msg "}
			if err := stream.SendMsg(&req); err != nil {
				log.Fatalf("can not send %v", err)
			}
		}
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			log.Printf("Response: [%s] ", resp.Message)
		}
	}()

	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()

	<-done

}
