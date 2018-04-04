package basic_auth

import (
	"net/http"
	"fmt"
)

type httpHandlerFunc func(http.ResponseWriter, *http.Request)

type AuthenticateFunc func(user, password string) error

type BasicAuthMiddleware struct {
	AuthenticateCallback AuthenticateFunc
	Realm                string
}

func (auth *BasicAuthMiddleware) writeChallenge(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", fmt.Sprint("Basic realm=\"%s", auth.Realm))
	w.WriteHeader(401)
	w.Write([]byte("401 Unauthorized\n"))
}

func (auth *BasicAuthMiddleware) writeInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(500)
	w.Write([]byte("Internal server error\n"))
}

func (auth *BasicAuthMiddleware) handle(w http.ResponseWriter, r *http.Request, next http.Handler) {

	username, password, ok := r.BasicAuth()

	if !ok {
		auth.writeChallenge(w)
		return
	}

	if err := auth.AuthenticateCallback(username, password); err != nil {
		switch err.(type) {
		case InvalidUserError:
			auth.writeChallenge(w)
			return

		case InvalidPasswordError:
			auth.writeChallenge(w)
			return

		default:
			auth.writeInternalServerError(w)
			return
		}
	}

	next.ServeHTTP(w, r)
}

func CreateBasicAuthMiddlewareFunc(auth BasicAuthMiddleware, next httpHandlerFunc) httpHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		auth.handle(w, r, http.HandlerFunc(next))
	}
}

func CreateBasicAuthMiddlewareHandler(auth BasicAuthMiddleware, next http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			auth.handle(w, r, next)
		})
}

