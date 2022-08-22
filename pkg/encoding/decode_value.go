package encoding

import (
	"fmt"

	"gopkg.in/yaml.v2"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	clientsetscheme "k8s.io/client-go/kubernetes/scheme"
	agi "k8s.io/kube-aggregator/pkg/apis/apiregistration/install"
)

var Scheme = clientsetscheme.Scheme
var Codecs = clientsetscheme.Codecs
var ParameterCodec = clientsetscheme.ParameterCodec

type Value runtime.Object

func init() {
	agi.Install(Scheme)
}

func decodeViaUnstruct(raw []byte) (Value, error) {
	u := map[string]interface{}{}
	err := yaml.Unmarshal(raw, &u)
	if err != nil {
		return nil, err
	}
	obj := &unstructured.Unstructured{Object: u}
	// ignore managedfields
	a, err := meta.Accessor(obj)
	if err != nil {
		return nil, err
	}
	// The object is not a `metav1.Object`, ignore it.
	a.SetManagedFields(nil)
	return obj, nil
}

func decodeViaCodec(raw []byte) (Value, error) {
	object, _, err := Codecs.UniversalDeserializer().Decode(raw, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("decode value to object fail err: %v", err)
	}
	// ignore managedfields
	a, err := meta.Accessor(object)
	if err == nil {
		// The object is not a `metav1.Object`, ignore it.
		a.SetManagedFields(nil)
	}
	return object, nil
}

func DecodeValue(raw []byte) (Value, error) {
	obj, err := decodeViaCodec(raw)
	if err == nil {
		return obj, nil
	}
	obj, err = decodeViaUnstruct(raw)
	return obj, err
}
