package aml

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

type mockTargetRequest struct {
	Name string `url:"name" json:"name"`
	Age  int    `url:"age" json:"age"`
}

func TestRequest_reqRaw(t *testing.T) {
	raw := reqRaw[mockTargetRequest]{
		UserID:   "user123",
		Password: "pass123",
		Target: mockTargetRequest{
			Name: "Alice",
			Age:  30,
		},
	}

	// Verify Query Encoding
	v := make(url.Values)
	err := raw.EncodeValues("", &v)
	require.NoError(t, err)
	spew.Dump(v)

	// Verify JSON Marshaling
	bytes, err := json.Marshal(raw)
	require.NoError(t, err)
	spew.Dump(string(bytes))
}
