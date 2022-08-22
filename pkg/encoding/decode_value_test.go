package encoding

import (
	b64 "encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {

	// value of /registry/gateway.networking.k8s.io/gatewayclasses/alb-dev
	rawb64 := "eyJhcGlWZXJzaW9uIjoiZ2F0ZXdheS5uZXR3b3JraW5nLms4cy5pby92MWFscGhhMiIsImtpbmQiOiJHYXRld2F5Q2xhc3MiLCJtZXRhZGF0YSI6eyJhbm5vdGF0aW9ucyI6eyJtZXRhLmhlbG0uc2gvcmVsZWFzZS1uYW1lIjoiYWxiLWRldiIsIm1ldGEuaGVsbS5zaC9yZWxlYXNlLW5hbWVzcGFjZSI6ImNwYWFzLXN5c3RlbSJ9LCJjcmVhdGlvblRpbWVzdGFtcCI6IjIwMjItMDgtMTlUMDc6MzU6MTNaIiwiZ2VuZXJhdGlvbiI6MSwibGFiZWxzIjp7ImFsYjIuY3BhYXMuaW8vZ2F0ZXdheWNsYXNzIjoiYWxiMiIsImFwcC5rdWJlcm5ldGVzLmlvL21hbmFnZWQtYnkiOiJIZWxtIn0sIm1hbmFnZWRGaWVsZHMiOlt7ImFwaVZlcnNpb24iOiJnYXRld2F5Lm5ldHdvcmtpbmcuazhzLmlvL3YxYWxwaGEyIiwiZmllbGRzVHlwZSI6IkZpZWxkc1YxIiwiZmllbGRzVjEiOnsiZjptZXRhZGF0YSI6eyJmOmFubm90YXRpb25zIjp7Ii4iOnt9LCJmOm1ldGEuaGVsbS5zaC9yZWxlYXNlLW5hbWUiOnt9LCJmOm1ldGEuaGVsbS5zaC9yZWxlYXNlLW5hbWVzcGFjZSI6e319LCJmOmxhYmVscyI6eyIuIjp7fSwiZjphbGIyLmNwYWFzLmlvL2dhdGV3YXljbGFzcyI6e30sImY6YXBwLmt1YmVybmV0ZXMuaW8vbWFuYWdlZC1ieSI6e319fSwiZjpzcGVjIjp7Ii4iOnt9LCJmOmNvbnRyb2xsZXJOYW1lIjp7fSwiZjpwYXJhbWV0ZXJzUmVmIjp7Ii4iOnt9LCJmOmdyb3VwIjp7fSwiZjpraW5kIjp7fSwiZjpuYW1lIjp7fSwiZjpuYW1lc3BhY2UiOnt9fX19LCJtYW5hZ2VyIjoiaGVsbSIsIm9wZXJhdGlvbiI6IlVwZGF0ZSIsInRpbWUiOiIyMDIyLTA4LTE5VDA3OjM1OjEzWiJ9XSwibmFtZSI6ImFsYi1kZXYiLCJ1aWQiOiIyNWMyMmVhMS0zNWI2LTRiNGYtOTgyNi0wNDgyMjJkMDk3ZmUifSwic3BlYyI6eyJjb250cm9sbGVyTmFtZSI6ImFsYjIuZ2F0ZXdheS5jcGFhcy5pby9hbGItZGV2IiwicGFyYW1ldGVyc1JlZiI6eyJncm91cCI6ImNyZC5hbGF1ZGEuaW8iLCJraW5kIjoiQUxCMiIsIm5hbWUiOiJhbGItZGV2IiwibmFtZXNwYWNlIjoiY3BhYXMtc3lzdGVtIn19fQo="
	raw := make([]byte, len(rawb64)*2)
	size, err := b64.StdEncoding.Decode(raw, []byte(rawb64))
	assert.NoError(t, err)
	raw = raw[:size]
	t.Logf("yaml %s\n", string(raw))
	obj, err := decodeViaUnstruct(raw)
	assert.NoError(t, err)
	t.Logf("obj %v", obj.GetObjectKind().GroupVersionKind().Group)
	t.Logf("obj %v", obj.GetObjectKind().GroupVersionKind().Kind)
	t.Logf("obj %v", obj.GetObjectKind().GroupVersionKind().Version)
	assert.Equal(t, obj.GetObjectKind().GroupVersionKind().Group, "gateway.networking.k8s.io")
	assert.Equal(t, obj.GetObjectKind().GroupVersionKind().Kind, "GatewayClass")
	assert.Equal(t, obj.GetObjectKind().GroupVersionKind().Version, "v1alpha2")
}
