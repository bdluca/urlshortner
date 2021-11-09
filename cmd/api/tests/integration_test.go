package tests

import (
	"encoding/json"
	"fmt"
	"github.com/brunodeluca/gophercises/urlshort/cmd/api/app"
	"github.com/brunodeluca/gophercises/urlshort/internal/platform/environment"
	"github.com/brunodeluca/gophercises/urlshort/internal/shortener/shorturl"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var server *http.ServeMux

func TestMain(m *testing.M) {
	dependencies := app.BuildDependencies(environment.Local)
	fakeApp := app.Build(dependencies)
	server = fakeApp
	m.Run()
}

func TestRegister(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		expectedCode     int
		expResp          bool
		expectedResponse string
	}{
		{
			name:             "OK",
			input:            "https://google.com",
			expectedCode:     http.StatusOK,
			expResp:          true,
			expectedResponse: "https://google.com",
		},
		{
			name:         "invalid url",
			input:        "google.com",
			expectedCode: http.StatusInternalServerError,
			expResp:      false,
		},
		{
			name:         "invalid url",
			input:        "aaa",
			expectedCode: http.StatusInternalServerError,
			expResp:      false,
		},
	}

	for _, tt := range tests {
		target := fmt.Sprintf("/register?url=%s", tt.input)
		res := performRequest(http.MethodGet, target, "", server, nil)

		if tt.expResp {
			bb, _ := ioutil.ReadAll(res.Body)
			var body shorturl.ShortURL
			_ = json.Unmarshal(bb, &body)

			assert.Equal(t, tt.expectedResponse, body.URL)
			assert.NotEmpty(t, body.ID)
		}

		assert.Equal(t, tt.expectedCode, res.Code)
	}
}

func performRequest(method, target, body string, server *http.ServeMux, headers map[string]string) *httptest.ResponseRecorder {
	var req *http.Request
	if body == "" {
		req = httptest.NewRequest(method, target, http.NoBody)
	} else {
		payload := strings.NewReader(body)
		req = httptest.NewRequest(method, target, payload)
	}
	for headKey, headValue := range headers {
		req.Header.Add(headKey, headValue)
	}
	res := httptest.NewRecorder()

	server.ServeHTTP(res, req)
	log.Printf(res.Body.String())
	return res
}
