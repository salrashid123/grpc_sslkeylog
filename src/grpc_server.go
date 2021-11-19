package main

import (
	"flag"
	"io"
	"math/rand"
	"net"
	"sync"

	"github.com/salrashid123/grpc_keylog/echo"

	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

var (
	tlsCert             = flag.String("tlsCert", "", "tls Certificate")
	tlsKey              = flag.String("tlsKey", "", "tls Key")
	grpcport            = flag.String("grpcport", "", "grpcport")
	insecure            = flag.Bool("insecure", false, "startup without TLS")
	unhealthyproability = flag.Int("unhealthyproability", 0, "percentage chance the service is unhealthy (0->100)")
)

const (
	address string = ":50051"
)

type Server struct {
	mu sync.Mutex
	// statusMap stores the serving status of the services this Server monitors.
	statusMap map[string]healthpb.HealthCheckResponse_ServingStatus
}

// NewServer returns a new Server.
func NewServer() *Server {
	return &Server{
		statusMap: make(map[string]healthpb.HealthCheckResponse_ServingStatus),
	}
}

func (s *Server) SayHelloUnary(ctx context.Context, in *echo.EchoRequest) (*echo.EchoReply, error) {
	log.Println("Got Unary Request: ")
	uid, _ := uuid.NewUUID()
	return &echo.EchoReply{Message: "SayHelloUnary Response " + uid.String()}, nil
}

func (s *Server) SayHelloServerStream(in *echo.EchoRequest, stream echo.EchoServer_SayHelloServerStreamServer) error {
	log.Println("Got SayHelloServerStream: Request ")
	for i := 0; i < 5; i++ {
		stream.Send(&echo.EchoReply{Message: "SayHelloServerStream Response"})
	}
	return nil
}

func (s Server) SayHelloBiDiStream(srv echo.EchoServer_SayHelloBiDiStreamServer) error {
	ctx := srv.Context()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}
		log.Printf("Got SayHelloBiDiStream %s", req.Name)
		resp := &echo.EchoReply{Message: "SayHelloBiDiStream Server Response"}
		if err := srv.Send(resp); err != nil {
			log.Printf("send error %v", err)
		}
	}
}

func (s Server) SayHelloClientStream(stream echo.EchoServer_SayHelloClientStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&echo.EchoReply{Message: "SayHelloClientStream  Response"})
		}
		if err != nil {
			return err
		}
		log.Printf("Got SayHelloClientStream Request: %s", req.Name)
	}
}

func (s *Server) Check(ctx context.Context, in *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if in.Service == "" {
		// return overall status
		return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
	}

	r := rand.Intn(100)
	if r <= *unhealthyproability {
		s.statusMap["echo.EchoServer"] = healthpb.HealthCheckResponse_NOT_SERVING
	} else {
		s.statusMap["echo.EchoServer"] = healthpb.HealthCheckResponse_SERVING
	}
	status, ok := s.statusMap[in.Service]
	if !ok {
		return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_UNKNOWN}, grpc.Errorf(codes.NotFound, "unknown service")
	}
	return &healthpb.HealthCheckResponse{Status: status}, nil
}

func (s *Server) Watch(in *healthpb.HealthCheckRequest, srv healthpb.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "Watch is not implemented")
}

func main() {

	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "INFO")
	flag.Parse()

	if *grpcport == "" {
		flag.Usage()
		log.Panicf("missing -grpcport flag (:50051)")
	}

	lis, err := net.Listen("tcp", *grpcport)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	sopts := []grpc.ServerOption{}
	if *insecure == false {
		if *tlsCert == "" || *tlsKey == "" {
			log.Fatalf("Must set --tlsCert and tlsKey if --insecure flags is not set")
		}
		ce, err := credentials.NewServerTLSFromFile(*tlsCert, *tlsKey)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		sopts = append(sopts, grpc.Creds(ce))
	}

	s := grpc.NewServer(sopts...)
	srv := NewServer()
	healthpb.RegisterHealthServer(s, srv)
	echo.RegisterEchoServerServer(s, srv)
	reflection.Register(s)
	log.Println("Starting Server...")
	s.Serve(lis)

}
