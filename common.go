package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

var (
	port, token, base string
	tls               bool
)

// init sets up defaults for optional command line arguments
func init() {
	const (
		defaultPort  = ":8080"
		usagePort    = "http service address"
		defaultToken = ""
		usageToken   = "http auth token"
		defaultBase  = "/var/www"
		usageBase    = "REPREPRO_BASE_DIR"
		defaultTls   = false
		usageTls     = "TLS boolean flag"
	)
	flag.StringVar(&port, "port", defaultPort, usagePort)
	flag.StringVar(&port, "p", defaultPort, usagePort+" (shorthand)")
	flag.StringVar(&token, "token", defaultToken, usageToken)
	flag.StringVar(&token, "t", defaultToken, usageToken+" (shorthand)")
	flag.StringVar(&base, "base", defaultBase, usageBase)
	flag.StringVar(&base, "b", defaultBase, usageBase+" (shorthand)")
	flag.BoolVar(&tls, "ssl", defaultTls, usageTls)
	flag.BoolVar(&tls, "s", defaultTls, usageTls+" (shorthand)")
}

// doMarshall returns a JSON object from any structure
func doMarshall(m interface{}) []byte {
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Println(err)
	}
	return b
}

// setCommand generates a string to be handled by the shell function
func setCommand(arg string) string {
	return "reprepro -b " + base + " " + arg + " "
}

// serve sets up the main http listener based on the tls command line argument
func serve() {
	switch tls {
	case false:
		err := http.ListenAndServe(port, nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	case true:
		err := http.ListenAndServeTLS(port, "cert.pem", "key.pem", nil)
		if err != nil {
			log.Fatal("ListenAndServeTLS: ", err)
		}
	}
}
