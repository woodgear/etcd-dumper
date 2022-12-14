package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/woodgear/etcd-dumper/pkg/encoding"
	"gopkg.in/yaml.v2"
)

type Etcd struct {
	Header `json:"header"`
	Kvs    []*KV `json:"kvs"`
	Count  int   `json:"count"`
}

type Header struct {
	ClusterID uint64 `json:"cluster_id" yaml:"cluster_id"`
	MemberID  uint64 `json:"member_id"`
	Revision  uint   `json:"revision"`
	RaftTerm  uint   `json:"raft_term"`
}

type KV struct {
	RawKey         string `json:"key"`
	Key            string
	CreateRevision uint   `json:"create_revision"`
	ModRevision    uint   `json:"mod_revision"`
	Version        uint   `json:"version"`
	RawValue       string `json:"value" yaml:"-"`
	Value          encoding.Value
}

func FromRawJsonByte(raw []byte) (*Etcd, error) {
	etcd := Etcd{}
	err := json.Unmarshal(raw, &etcd)
	if err != nil {
		return nil, fmt.Errorf("paser json fail err %v", err)
	}
	err = etcd.decode()
	if err != nil {
		return nil, fmt.Errorf("parse raw data fail err %v", err)
	}
	return &etcd, nil
}
func FromRawJson(rawJson string) (*Etcd, error) {
	return FromRawJsonByte([]byte(rawJson))
}

func (e *Etcd) ToYaml() (string, error) {
	yamlData, err := yaml.Marshal(e)
	if err != nil {
		return "", err
	}
	return string(yamlData), nil
}

func (e *Etcd) decode() error {
	for i := range e.Kvs {
		entry := e.Kvs[i]
		k_b64_str := entry.RawKey
		v_b64_str := entry.RawValue

		k_str_bytes, err := b64.StdEncoding.DecodeString(k_b64_str)
		k_str := string(k_str_bytes)
		if err != nil {
			return err
		}
		// fmt.Printf("rk %v k %v\n", entry.RawKey, k_str)
		v_bytes := make([]byte, len(v_b64_str)*2)
		size, err := b64.StdEncoding.Decode(v_bytes, []byte(v_b64_str))
		if err != nil {
			return err
		}
		v_bytes = v_bytes[:size]
		v, err := encoding.DecodeValue(v_bytes)
		if err != nil {
			return err
		}
		entry.Key = string(k_str)
		entry.Value = v
	}
	return nil
}

func run() error {
	file := os.Args[1]
	raw, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	e, err := FromRawJsonByte(raw)
	if err != nil {
		return err
	}
	yaml_str, err := e.ToYaml()
	if err != nil {
		return err
	}
	_ = yaml_str
	fmt.Print(yaml_str)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
