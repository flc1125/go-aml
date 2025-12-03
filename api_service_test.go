package aml

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_QueryOver(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/amlupload/api/queryOver", r.URL.Path)

		_, _ = w.Write(loadData(t, "testdata/query-over.json"))
	}))

	resp, _, err := client.QueryOver(t.Context())
	require.NoError(t, err)
	assert.Equal(t, 11, resp.Remain)
	assert.Equal(t, 22, resp.Over)
}
