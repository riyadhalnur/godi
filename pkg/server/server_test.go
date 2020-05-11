package server

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/godi/pkg/server/util"
)

func TestRouteMount(t *testing.T) {
	testMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}

	testHandler := func(ctx context.Context, req *util.Request) (*util.Response, error) {
		return &util.Response{
			StatusCode: http.StatusOK,
			Body:       "ok",
		}, nil
	}

	testRoutes := []util.Route{
		util.Route{
			"test",
			"/test",
			http.MethodGet,
			testHandler,
		},
	}

	srv := Server{}
	srv.AddMiddlewares(testMiddleware)
	srv.AddRoutes(testRoutes...)

	assert.Equal(t, 1, len(srv.routers))
	assert.Equal(t, 1, len(srv.middlewares))
}

func TestHealthCheck(t *testing.T) {
	srv := Server{
		config: &Config{},
	}
	router := srv.mountRoutes()

	req, err := http.NewRequest(http.MethodGet, "/healthz", nil)
	if err != nil {
		assert.Nil(t, err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, "ok", rr.Body.String())
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestServerResponse(t *testing.T) {
	const (
		endpoint string = "/test"
	)

	t.Run("internal server error", func(t *testing.T) {
		testHandler := func(ctx context.Context, req *util.Request) (*util.Response, error) {
			err := errors.New("something went wrong in handler")
			return nil, err
		}

		testRoutes := []util.Route{
			util.Route{
				"test",
				endpoint,
				http.MethodGet,
				testHandler,
			},
		}

		srv := Server{
			config: &Config{},
		}
		srv.AddRoutes(testRoutes...)
		router := srv.mountRoutes()

		req, err := http.NewRequest(http.MethodGet, endpoint, nil)
		if err != nil {
			assert.Nil(t, err)
		}

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		body, _ := ioutil.ReadAll(rr.Body)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.JSONEq(t, `{"code":500,"message":"something went wrong in handler"}`, string(body))
	})

	t.Run("request error", func(t *testing.T) {
		testHandler := func(ctx context.Context, req *util.Request) (*util.Response, error) {
			return &util.Response{
				StatusCode: http.StatusBadRequest,
				Body:       "something in the request is wrong",
			}, nil
		}

		testRoutes := []util.Route{
			util.Route{
				"test",
				endpoint,
				http.MethodGet,
				testHandler,
			},
		}

		srv := Server{
			config: &Config{},
		}
		srv.AddRoutes(testRoutes...)
		router := srv.mountRoutes()

		req, err := http.NewRequest(http.MethodGet, endpoint, nil)
		if err != nil {
			assert.Nil(t, err)
		}

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "something in the request is wrong", rr.Body.String())
	})

	t.Run("redirects", func(t *testing.T) {
		testHandler := func(ctx context.Context, req *util.Request) (*util.Response, error) {
			return &util.Response{
				StatusCode: http.StatusMovedPermanently,
				Headers: map[string]string{
					"Location": "https://google.com",
				},
			}, nil
		}

		testRoutes := []util.Route{
			util.Route{
				"test",
				endpoint,
				http.MethodGet,
				testHandler,
			},
		}

		srv := Server{
			config: &Config{},
		}
		srv.AddRoutes(testRoutes...)
		router := srv.mountRoutes()

		r := httptest.NewServer(router)
		defer r.Close()

		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				assert.Equal(t, "google.com", req.URL.Host)
				return http.ErrUseLastResponse
			},
			Timeout: 2 * time.Second,
		}

		res, err := client.Get(r.URL + endpoint)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusMovedPermanently, res.StatusCode)
	})
}
