package basic_auth

import (
	"net/http"
	"io"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"errors"
)

const (
	TestRealName = "Test"
)

func Test_When_Missing_AuthorizationHeader_Using_Func(t *testing.T) {

	md := BasicAuthMiddleware{
		AuthenticateCallback: func(u, p string) error {
			return nil
		},
		Realm:TestRealName,
	}

	handler := CreateBasicAuthMiddlewareFunc(md, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	})

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	assert.Equal(t, resp.StatusCode, 401, "they should be equal")
}

func Test_When_AuthorizationHeader_Using_Func_Should_Pass(t *testing.T) {

	md := BasicAuthMiddleware{
		AuthenticateCallback: func(u, p string) error {
			return nil
		},
		Realm:TestRealName,
	}

	handler := CreateBasicAuthMiddlewareFunc(md, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	})

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	req.SetBasicAuth(TestUserName, TestPassword)

	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	assert.Equal(t, resp.StatusCode, 200, "they should be equal")
}

func Test_When_AuthorizationHeader_Using_Func_Returns_UnknownError_Status_Is_500(t *testing.T) {

	md := BasicAuthMiddleware{
		AuthenticateCallback: func(u, p string) error {
			return errors.New("Unknown error")
		},
		Realm:TestRealName,
	}

	handler := CreateBasicAuthMiddlewareFunc(md, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	})

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	req.SetBasicAuth(TestUserName, TestPassword)

	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	assert.Equal(t, resp.StatusCode, 500, "they should be equal")
}

func Test_When_AuthorizationHeader_Using_Func_Returns_InvalidUserError_Status_Is_401(t *testing.T) {

	md := BasicAuthMiddleware{
		AuthenticateCallback: func(u, p string) error {
			return NewInvalidUserError(u)
		},
		Realm:TestRealName,
	}

	handler := CreateBasicAuthMiddlewareFunc(md, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	})

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	req.SetBasicAuth(TestUserName, TestPassword)

	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	assert.Equal(t, resp.StatusCode, 401, "they should be equal")
}

func Test_When_AuthorizationHeader_Using_Func_Returns_InvalidPasswordError_Status_Is_401(t *testing.T) {

	md := BasicAuthMiddleware{
		AuthenticateCallback: func(u, p string) error {
			return NewInvalidPasswordError(p)
		},
		Realm:TestRealName,
	}

	handler := CreateBasicAuthMiddlewareFunc(md, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	})

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	req.SetBasicAuth(TestUserName, TestPassword)

	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	assert.Equal(t, resp.StatusCode, 401, "they should be equal")
}

func Test_When_AuthorizationHeader_Using_Func_Returns_Nill_And_Check_User_And_Password_Status_Is_200(t *testing.T) {

	md := BasicAuthMiddleware{
		AuthenticateCallback: func(u, p string) error {

			if u != TestUserName {
				return NewInvalidUserError(u)

			}

			if p != TestPassword {
				return NewInvalidPasswordError(p)
			}

			return nil
		},
		Realm:TestRealName,
	}

	handler := CreateBasicAuthMiddlewareFunc(md, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	})

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	req.SetBasicAuth(TestUserName, TestPassword)

	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	assert.Equal(t, resp.StatusCode, 200, "they should be equal")
}


func Test_When_Missing_AuthorizationHeader_Using_Handler(t *testing.T) {

	md := BasicAuthMiddleware{
		AuthenticateCallback: func(u, p string) error {
			return nil
		},
		Realm:TestRealName,
	}

	handler := CreateBasicAuthMiddlewareHandler(md, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}))

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	resp := w.Result()

	assert.Equal(t, resp.StatusCode, 401, "they should be equal")
}