package main

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "knrp <gateway> <service>",
	Short:   "Knative Reverse Proxy (knrp) allows to access Knative service with/without magic DNS",
	Args:    cobra.ExactArgs(2),
	Example: "  knrp localhost:80 hello.default.127.0.0.1.sslip.io\n  knrp localhost:80 hello.default.example.com",
	RunE: func(cmd *cobra.Command, args []string) error {
		gateway, err := url.ParseRequestURI("http://" + strings.TrimPrefix(args[0], "http://"))
		if err != nil {
			return err
		}
		service := strings.TrimPrefix(args[1], "http://")

		listener, err := net.Listen("tcp", ":0")
		if err != nil {
			return err
		}

		log.Printf("knrp run at port: %d", listener.Addr().(*net.TCPAddr).Port)
		return http.Serve(listener, proxyHandler(gateway, service))
	},
}

func proxyHandler(gateway *url.URL, service string) http.HandlerFunc {
	proxy := httputil.NewSingleHostReverseProxy(gateway)
	originalDirector := proxy.Director
	proxy.Director = func(r *http.Request) {
		originalDirector(r)
		r.Host = service
	}

	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
