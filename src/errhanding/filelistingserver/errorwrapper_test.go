package main

import (
	"net/http"
	"testing"
	"net/http/httptest"
	"os"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error  {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string  {
	return e.Message()
}

func (e testingUserError) Message() string  {
	return string(e)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error  {
	return testingUserError("user error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error  {
	return os.ErrNotExist
}

func errNotPermission(writer http.ResponseWriter, request *http.Request) error  {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter, request *http.Request) error  {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter, requst *http.Request) error  {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct{
		h appHandler
		code int
		message string
} {
		{errPanic, 500, "Internal Server Error"},
		{errUserError, 400, "user error"},
		{errNotFound, 404, "Not Found"},
		{errNotPermission, 403, "Forbidden"},
		{errUnknown, 500, "Internal Server Error"},
		{noError, 200, "no error"},
}

func TestErrWrapper(t *testing.T)  {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()   //假的response
		request := httptest.NewRequest(http.MethodGet, "http://www.imooc.com", nil)  //假的request
		f(response, request)

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T)  {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)
		verifyResponse(resp, tt.code, tt.message, t)

	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMessage string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode || body != expectedMessage {
		t.Errorf("expected (%d, %s); got (%d, %s)", expectedCode, expectedMessage, resp.StatusCode, body)
	}
}