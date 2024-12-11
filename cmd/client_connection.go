package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func getClientConn(c commonConfig) (*grpc.ClientConn, error) {

	address := fmt.Sprintf("%s:%d", c.server, c.port)

	if c.insecure {
		return grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	tlsConf := &tls.Config{}
	if len(c.cacert) > 0 {
		certPool := x509.NewCertPool()
		ca, err := os.ReadFile(c.cacert)
		if err != nil {
			return nil, fmt.Errorf("could not read ca certificate: %s", err)
		}
		// Append the certificates from the CA
		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			return nil, errors.New("failed to append ca certs")
		}
		tlsConf.RootCAs = certPool
	}

	if c.mtls {

		certificate, err := tls.LoadX509KeyPair(c.cert, c.key)
		if err != nil {
			return nil, fmt.Errorf("could not load client key pair: %s", err)
		}

		tlsConf.ServerName = c.server
		tlsConf.Certificates = []tls.Certificate{certificate}
	}

	return grpc.NewClient(address, grpc.WithTransportCredentials(credentials.NewTLS(tlsConf)))
}
