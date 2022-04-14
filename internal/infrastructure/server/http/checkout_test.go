package http_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	internalHTTP "github.com/matheussbaraglini/hash-challenge/internal/infrastructure/server/http"
	"github.com/stretchr/testify/assert"
)

func TestCheckout(t *testing.T) {
	t.Run("should validate successfully", func(t *testing.T) {
		server := httptest.NewServer(internalHTTP.NewHandler())
		defer server.Close()

		URL, _ := url.Parse(server.URL)

		body := strings.NewReader(`
		{
			"products": [
				{
					"id": 1,
					"quantity": 1
				}
			]
		}`)

		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/checkout", URL), body)
		assert.NoError(t, err)

		res, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.StatusCode)

		bodyBytes, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.NotEmpty(t, bodyBytes)

		expected := `
		{
			"products": [
				{
					"id": 1,
					"quantity": 1
				}
			]
		}`

		assert.JSONEq(t, expected, string(bodyBytes))
	})
}
