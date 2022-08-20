package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDump(t *testing.T) {
	data := `
	{
    "header": {
        "cluster_id": 14358680983224840443,
        "member_id": 1033796535975940116,
        "revision": 959834,
        "raft_term": 6
    },
    "kvs": [
        {
            "key": "L3JlZ2lzdHJ5L2FwaXJlZ2lzdHJhdGlvbi5rOHMuaW8vYXBpc2VydmljZXMvdjEu",
            "create_revision": 9,
            "mod_revision": 9,
            "version": 1,
            "value": "eyJraW5kIjoiQVBJU2VydmljZSIsImFwaVZlcnNpb24iOiJhcGlyZWdpc3RyYXRpb24uazhzLmlvL3YxIiwibWV0YWRhdGEiOnsibmFtZSI6InYxLiIsInVpZCI6ImVlOGNiMDg0LTBiM2MtNDUyZC05MzM4LWRjMzNlM2U3NDc0NiIsImNyZWF0aW9uVGltZXN0YW1wIjoiMjAyMi0wOC0yMFQwMzoxMjozNloiLCJsYWJlbHMiOnsia3ViZS1hZ2dyZWdhdG9yLmt1YmVybmV0ZXMuaW8vYXV0b21hbmFnZWQiOiJvbnN0YXJ0In0sIm1hbmFnZWRGaWVsZHMiOlt7Im1hbmFnZXIiOiJrdWJlLWFwaXNlcnZlciIsIm9wZXJhdGlvbiI6IlVwZGF0ZSIsImFwaVZlcnNpb24iOiJhcGlyZWdpc3RyYXRpb24uazhzLmlvL3YxIiwidGltZSI6IjIwMjItMDgtMjBUMDM6MTI6MzZaIiwiZmllbGRzVHlwZSI6IkZpZWxkc1YxIiwiZmllbGRzVjEiOnsiZjptZXRhZGF0YSI6eyJmOmxhYmVscyI6eyIuIjp7fSwiZjprdWJlLWFnZ3JlZ2F0b3Iua3ViZXJuZXRlcy5pby9hdXRvbWFuYWdlZCI6e319fSwiZjpzcGVjIjp7ImY6Z3JvdXBQcmlvcml0eU1pbmltdW0iOnt9LCJmOnZlcnNpb24iOnt9LCJmOnZlcnNpb25Qcmlvcml0eSI6e319fX1dfSwic3BlYyI6eyJ2ZXJzaW9uIjoidjEiLCJncm91cFByaW9yaXR5TWluaW11bSI6MTgwMDAsInZlcnNpb25Qcmlvcml0eSI6MX0sInN0YXR1cyI6eyJjb25kaXRpb25zIjpbeyJ0eXBlIjoiQXZhaWxhYmxlIiwic3RhdHVzIjoiVHJ1ZSIsImxhc3RUcmFuc2l0aW9uVGltZSI6IjIwMjItMDgtMjBUMDM6MTI6MzZaIiwicmVhc29uIjoiTG9jYWwiLCJtZXNzYWdlIjoiTG9jYWwgQVBJU2VydmljZXMgYXJlIGFsd2F5cyBhdmFpbGFibGUifV19fQo="
        },
        {
            "key": "L3JlZ2lzdHJ5L3Jhbmdlcy9zZXJ2aWNlaXBz",
            "create_revision": 4,
            "mod_revision": 947841,
            "version": 4,
            "value": "azhzAAohChFzdG9yYWdlLms4cy5pby92MRIMU3RvcmFnZUNsYXNzEp0GCuUFCghzdGFuZGFyZBIAGgAiACokZDA5YjE5OWMtYWZlZC00Y2JkLWIwODctODI2YWU0YjBmMTg4MgA4AEIICKiigZgGEABivAIKMGt1YmVjdGwua3ViZXJuZXRlcy5pby9sYXN0LWFwcGxpZWQtY29uZmlndXJhdGlvbhKHAnsiYXBpVmVyc2lvbiI6InN0b3JhZ2UuazhzLmlvL3YxIiwia2luZCI6IlN0b3JhZ2VDbGFzcyIsIm1ldGFkYXRhIjp7ImFubm90YXRpb25zIjp7InN0b3JhZ2VjbGFzcy5rdWJlcm5ldGVzLmlvL2lzLWRlZmF1bHQtY2xhc3MiOiJ0cnVlIn0sIm5hbWUiOiJzdGFuZGFyZCJ9LCJwcm92aXNpb25lciI6InJhbmNoZXIuaW8vbG9jYWwtcGF0aCIsInJlY2xhaW1Qb2xpY3kiOiJEZWxldGUiLCJ2b2x1bWVCaW5kaW5nTW9kZSI6IldhaXRGb3JGaXJzdENvbnN1bWVyIn0KYjMKK3N0b3JhZ2VjbGFzcy5rdWJlcm5ldGVzLmlvL2lzLWRlZmF1bHQtY2xhc3MSBHRydWV6AIoBpwIKGWt1YmVjdGwtY2xpZW50LXNpZGUtYXBwbHkSBlVwZGF0ZRoRc3RvcmFnZS5rOHMuaW8vdjEiCAioooGYBhAAMghGaWVsZHNWMTrYAQrVAXsiZjptZXRhZGF0YSI6eyJmOmFubm90YXRpb25zIjp7Ii4iOnt9LCJmOmt1YmVjdGwua3ViZXJuZXRlcy5pby9sYXN0LWFwcGxpZWQtY29uZmlndXJhdGlvbiI6e30sImY6c3RvcmFnZWNsYXNzLmt1YmVybmV0ZXMuaW8vaXMtZGVmYXVsdC1jbGFzcyI6e319fSwiZjpwcm92aXNpb25lciI6e30sImY6cmVjbGFpbVBvbGljeSI6e30sImY6dm9sdW1lQmluZGluZ01vZGUiOnt9fUIAEhVyYW5jaGVyLmlvL2xvY2FsLXBhdGgiBkRlbGV0ZToUV2FpdEZvckZpcnN0Q29uc3VtZXIaACIA"
        }
    ],
    "count": 1
}`
	e, err := FromRawJson(data)
	assert.NoError(t, err)
	assert.Equal(t, e.Kvs[0].Key, "/registry/apiregistration.k8s.io/apiservices/v1.")
	assert.Equal(t, e.Kvs[0].Value.GetObjectKind().GroupVersionKind().Kind, "APIService")
	assert.Equal(t, e.Kvs[1].Key, "/registry/ranges/serviceips")
	assert.Equal(t, e.Kvs[1].Value.GetObjectKind().GroupVersionKind().Kind, "StorageClass")
	yaml_str, err := e.ToYaml()
	assert.NoError(t, err)
	t.Log(yaml_str)
}
