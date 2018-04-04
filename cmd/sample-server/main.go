// This is the test application for our basic authentication middleware sample

package main

import (
	"net/http"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
	auth "github.com/nickyborisov/basic-auth-sample/basic-auth-middleware"
)

type User struct {
	Id   string
	Name string
}

func main() {

	log.Debug("starting the server");

	md := auth.BasicAuthMiddleware{
		AuthenticateCallback:func(u, p string) error {
			if u != "nicky" {
				return auth.NewInvalidUserError(u)
			}

			if p != "borisov" {
				return auth.NewInvalidPasswordError(p)
			}

			return nil
		},
		Realm:"Basic Auth Sample",

	};

	http.HandleFunc("/", auth.CreateBasicAuthMiddlewareFunc(md,
		func(w http.ResponseWriter, r *http.Request) {

			log.Debug("new reuest received");

			u := User{Id: "US123", Name: "Nicky"}
			json.NewEncoder(w).Encode(u)
		}))
	log.Fatal(http.ListenAndServe(":8033", nil))
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}