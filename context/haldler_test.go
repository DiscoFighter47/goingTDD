package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type storeSpy struct {
	response string
	cancled  bool
}

func (spy *storeSpy) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return spy.response
}

func (spy *storeSpy) Cancle() {
	spy.cancled = true
}

func TestHandler(t *testing.T) {
	t.Run("Fetch data", func(t *testing.T) {
		data := "Hello, world"
		spy := &storeSpy{
			response: data,
		}
		svr := NewHandler(spy)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)

		if spy.cancled {
			t.Errorf("request was cancled")
		}
		if response.Body.String() != data {
			t.Errorf("got '%s' want '%s'", response.Body.String(), data)
		}
	})

	t.Run("Cancle request", func(t *testing.T) {
		data := "Hello, world"
		spy := &storeSpy{
			response: data,
		}
		svr := NewHandler(spy)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancleCtx, cancle := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Microsecond, cancle)
		request = request.WithContext(cancleCtx)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)

		if !spy.cancled {
			t.Errorf("request wasn't not cancled")
		}
	})
}
