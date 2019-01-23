package gmcredentials

//
//import (
//	"crypto/tls"
//	"crypto/x509"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"testing"
//)
//
//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
//}
//
//func TestHTTPSServer(t *testing.T) {
//	//cert, err := gmtls.LoadX509KeyPair(ca, cakey)
//	//if err != nil {
//	//	t.Fatal(err)
//	//}
//	certPool := x509.NewCertPool()
//	caCert, err := ioutil.ReadFile(ca)
//	if err != nil {
//		t.Fatal(err)
//	}
//	certPool.AppendCertsFromPEM(caCert)
//
//	defaultTLS := &tls.Config{
//		ClientCAs: certPool,
//		//ClientAuth: tls.RequireAndVerifyClientCert,	//证书双向认证
//		ClientAuth: tls.RequestClientCert, //证书单向认证
//	}
//
//	server := &http.Server{
//		Addr:       ":8888",
//		Handler:    nil,
//		TLSConfig:  defaultTLS,
//	}
//	http.HandleFunc("/", handler)
//	defer server.Close()
//
//	//server := &http.Server{
//	//	Addr:       ":8888",
//	//	Handler:    nil,
//	//}
//	//http.HandleFunc("/", handler)
//	defer server.Close()
//
//	//err = server.ListenAndServe()
//	err = server.ListenAndServeTLS(
//		ca,
//		cakey)
//	if err != nil {
//		t.Fatal("Server exit. ListenAndServeTLS error: "+err.Error())
//		return
//	}
//}
