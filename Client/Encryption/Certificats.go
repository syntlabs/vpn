package main

import (
	"fmt",
	"net"
)

func getKey(pkX string) []byte {

	validFileFormat(pkX, ".pem")

	keyX, xerr := os.ReadFile(os.Getenv(pkX))
	if xerr != nil {
		raise(Err.cyph.Key)
	}

	return keyX
}

func getCert(certPath string) []byte {

	validFileFormat(certPath, ".pem")

	cert, err := os.ReadFile(os.Getenv(certPath))
	if err != nil {
		raise(Err.cyph.Cert)
	}

	return cert
}

func certnpool() (tls.Certificate, *x509.CertPool) {

	key := getKey(keypath)
	cert := getCert(certpath)
	tlsCert, err := tls.X509KeyPair(cert, key)
	if err != nil {
		raise(Err.cyph.PairCert)
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)

	return tlsCert, certPool
}

func exchangeCerts () {
	
}