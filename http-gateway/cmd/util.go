package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"os"
)

func getCertPath() string {
	path := os.Getenv("CERT_DIR_PATH")
	if path == "" {
		panic("NO <CERT_DIR_PATH> ENV VARIABLE")
	}
	return path
}

func loadClientTLSCredentials() (credentials.TransportCredentials, error) {
	pemServerCA, err := ioutil.ReadFile(fmt.Sprintf("%s/cert/ca-cert.pem", getCertPath()))
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	clientCert, err := tls.LoadX509KeyPair(fmt.Sprintf("%s/cert/client-cert.pem", getCertPath()), fmt.Sprintf("%s/cert/client-key.pem", getCertPath()))
	if err != nil {
		return nil, err
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}
