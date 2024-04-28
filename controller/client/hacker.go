package client

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (d *delegatingClient) GroupVersionKindFor(obj runtime.Object) (schema.GroupVersionKind, error) {
	//TODO implement me
	panic("implement me")
}

func (d *delegatingClient) IsObjectNamespaced(obj runtime.Object) (bool, error) {
	//TODO implement me
	panic("implement me")
}
