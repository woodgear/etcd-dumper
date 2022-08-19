package main

import (
	b64 "encoding/base64"
	"fmt"
	"os"

	"k8s.io/kubectl/pkg/scheme"
)

type Etcd struct {
	Header Header
	Kvs    []KV
}

type Header struct {

}

type KV struct {
	Key string
}

func main() {
	dest := make([]byte, 1024)
	b64.StdEncoding.Decode(dest, []byte(os.Args[1]))
	decoder := scheme.Codecs.UniversalDeserializer()
	obj, gvk, err := decoder.Decode(dest, nil, nil)
	fmt.Print(obj)
	fmt.Print(gvk)
	fmt.Print(err)
}
