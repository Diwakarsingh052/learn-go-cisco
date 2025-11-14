package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDoubleHandler(t *testing.T) {
	tt := []struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "OK",
			queryParam:     "2",
			expectedStatus: 200,
			expectedBody:   "double of number: 2 : 4",
		},
		{
			name:           "empty",
			queryParam:     "",
			expectedStatus: 400,
			expectedBody:   "missing query parameter",
		},
	}
	// we are running the test without running a actual http server
	// ranging over the test cases
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			// ResponseRecorder is an implementation of [http.ResponseWriter] that
			// records its mutations for later inspection in tests.
			w := httptest.NewRecorder()

			r := httptest.NewRequest(http.MethodGet, "/double?v="+tc.queryParam, nil)

			// calling the handler function directly
			// we are not running the server to call the function
			doubleHandler(w, r)

			require.Equal(t, tc.expectedStatus, w.Code)
			// trim the whitespace from the response body
			body := strings.TrimSpace(w.Body.String())

			require.Equal(t, tc.expectedBody, body)

		})
	}
}
