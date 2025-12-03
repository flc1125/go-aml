package aml

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockTargetRequest struct {
	Name string `url:"name" json:"name"`
	Age  int    `url:"age" json:"age"`
}

func TestRequest_reqRaw(t *testing.T) {
	raw := reqRaw{
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
	assert.Equal(t, "user123", v.Get("userid"))
	assert.Equal(t, "pass123", v.Get("password"))
	assert.Equal(t, "Alice", v.Get("name"))
	assert.Equal(t, "30", v.Get("age"))

	// Verify JSON Marshaling
	bytes, err := json.Marshal(raw)
	require.NoError(t, err)
	var data map[string]any
	err = json.Unmarshal(bytes, &data)
	require.NoError(t, err)
	assert.Equal(t, "user123", data["userid"])
	assert.Equal(t, "pass123", data["password"])
	assert.Equal(t, "Alice", data["name"])
	assert.Equal(t, float64(30), data["age"]) // JSON numbers are float64
}
