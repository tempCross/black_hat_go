package main

import (
	"crypto/x509/pkix"
	"crypto/x509"
	"crypto/rsa"
    "crypto/rand"
    "log"
    "net"
    "math/big"
    "time"
    "errors"
    "fmt"
    "encoding/pem"
    "net/http/httptest"
    "net/http"
    "crypto/tls"
   )
//helper function to create template with serial number and other requires fields

func CertTemplate()(*x509.Certificate, error) {
	// generate a random serial number (a real cert authority would have some logic behind this)
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, errors.New("failed to generate serial number: " +err.Error())

	} 

	tmpl := x509.Certificate{
		SerialNumber:			serialNumber,
		Subject:				pkix.Name{Organization: []string{"Yhat, Inc."}},
		SignatureAlgorithm: 	x509.SHA256WithRSA,
		NotBefore:				time.Now(),
		NotAfter:				time.Now().Add(time.Hour), // valid for an hour
		BasicConstraintsValid:  true,
	}
	return &tmpl, nil
}

func CreateCert(template, parent *x509.Certificate, pub interface{}, parentPriv interface{})(
	cert *x509.Certificate, certPEM []byte, err error){

	certDER, err := x509.CreateCertificate(rand.Reader, template, parent, pub, parentPriv)
	if err != nil {
		return
	}
	// parse the resulting certificate so we can use it again
	cert, err = x509.ParseCertificate(certDER)
	if err != nil {
		return
	}
	// PEM encode the certificate (this is a standard TLS encoding)
	b := pem.Block{Type: "CERTIFICATE", Bytes: certDER}
	certPEM = pem.EncodeToMemory(&b)
	return	
}
func main(){

// generate a new key-pair
rootKey, err := rsa.GenerateKey(rand.Reader, 2048)
if err != nil {
	log.Fatalf("generating random key: %v", err)
}

rootCertTmpl, err := CertTemplate()
if err != nil {
	log.Fatalf("creating cert template: %v", err)
}
// describe what the certificate will be used for
rootCertTmpl.IsCA = true
rootCertTmpl.KeyUsage = x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature
rootCertTmpl.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth}
rootCertTmpl.IPAddresses = []net.IP{net.ParseIP("10.0.0.121")}

rootCert, rootCertPEM, err := CreateCert(rootCertTmpl, rootCertTmpl, &rootKey.PublicKey, rootKey)
if err != nil {
	log.Fatalf("error creating cert: %v", err)
}
fmt.Printf("%s\n", rootCertPEM)
fmt.Printf("%#x\n", rootCert.Signature) // more ugly binary

// PEM encode the private key
rootKeyPEM := pem.EncodeToMemory(&pem.Block{
	Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(rootKey),
})

// Create a TLS cert using the private key and certificate
rootTLSCert, err := tls.X509KeyPair(rootCertPEM, rootKeyPEM)
if err != nil {
	log.Fatalf("invalid key pair: %v", err)
}

ok := func(w http.ResponseWriter, r *http.Request) {w.Write([]byte("Hi"))}
s := httptest.NewUnstartedServer(http.HandlerFunc(ok))

// Configure the server to present the certificate we created
s.TLS = &tls.Config{
	Certificates: []tls.Certificate{rootTLSCert},
}
// make a HTTPS request to the server
s.StartTLS()
_, err = http.Get(s.URL)
s.Close()

fmt.Println(err)
// http: TLS handshake error from
}
