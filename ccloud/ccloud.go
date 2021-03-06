package main

import (
	"flag"
	"github.com/xlvector/caspercloud"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(6)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	port := flag.String("port", "8000", "port number")
	flag.Parse()

	service := caspercloud.NewCasperServer()
	http.Handle("/submit", service)
	l, e := net.Listen("tcp", ":"+*port)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
