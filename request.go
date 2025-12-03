package aml

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type RequestOption func(*http.Request) error

func WithRequestHeader(name, value string) RequestOption {
	return func(req *http.Request) error {
		req.Header.Set(name, value)
		return nil
	}
}

func WithRequestHeaders(headers map[string]string) RequestOption {
	return func(req *http.Request) error {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
		return nil
	}
}

func WithRequestHeaderFunc(fn func(http.Header)) RequestOption {
	return func(req *http.Request) error {
		fn(req.Header)
		return nil
	}
}

func WithRequestUserAgent(userAgent string) RequestOption {
	return func(req *http.Request) error {
		req.Header.Set("User-Agent", userAgent)
		return nil
	}
}

type reqRaw[T any] struct {
	UserID   string `json:"userid" url:"userid"`
	Password string `json:"password" url:"password"`
	Target   T      `json:"-" url:"-"`
}

var (
	_ json.Marshaler = (*reqRaw[any])(nil)
	_ query.Encoder  = (*reqRaw[any])(nil)
)

func (r reqRaw[T]) MarshalJSON() ([]byte, error) {
	// 把 Target 转成 map[string]any
	tBytes, err := json.Marshal(r.Target)
	if err != nil {
		return nil, err
	}

	var tMap map[string]any
	if err := json.Unmarshal(tBytes, &tMap); err != nil {
		return nil, err
	}

	// 合并字段
	out := make(map[string]any, len(tMap)+2)
	for k, v := range tMap {
		out[k] = v
	}

	out["userid"] = r.UserID
	out["password"] = r.Password

	return json.Marshal(out)
}

func (r reqRaw[T]) EncodeValues(key string, v *url.Values) error {
	// Encode the Target field
	nv, err := query.Values(r.Target)
	if err != nil {
		return err
	}

	// Merge target values into v
	for k, vs := range nv {
		for _, val := range vs {
			v.Add(k, val)
		}
	}

	// Add authentication fields
	v.Add("userid", r.UserID)
	v.Add("password", r.Password)

	return nil
}
