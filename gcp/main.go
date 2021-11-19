package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	projectID  = "mineral-minutia-820"
	pubsubHost = "pubsub.googleapis.com:443"
)

func main() {

	ctx := context.Background()

	var tlsCfg tls.Config
	rootCAs := x509.NewCertPool()
	pem, err := ioutil.ReadFile("gcp_roots.pem")
	if err != nil {
		log.Fatalf("failed to load root CA certificates  error=%v", err)
	}
	if !rootCAs.AppendCertsFromPEM(pem) {
		log.Fatalf("no root CA certs parsed from file ")
	}
	tlsCfg.RootCAs = rootCAs
	tlsCfg.ServerName = "pubsub.googleapis.com"

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

	rpcreds := credentials.NewTLS(&tlsCfg)

	gcpcreds, err := google.FindDefaultCredentials(ctx, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		log.Fatalf("Error getting default credentials %v", err)
	}

	client, err := pubsub.NewClient(ctx, projectID,
		option.WithGRPCDialOption(grpc.WithTransportCredentials(rpcreds)),
		option.WithCredentials(gcpcreds))
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}
	topics := client.Topics(ctx)
	for {
		topic, err := topics.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error listing topics %v", err)
		}
		log.Println(topic)
	}

}
