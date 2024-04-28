package multicluster

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (m *gatedClient) GroupVersionKindFor(obj runtime.Object) (schema.GroupVersionKind, error) {
	//TODO implement me
	panic("implement me")
}

func (m *gatedClient) IsObjectNamespaced(obj runtime.Object) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (in *remoteClusterClient) GroupVersionKindFor(obj runtime.Object) (schema.GroupVersionKind, error) {
	//TODO implement me
	panic("implement me")
}

func (in *remoteClusterClient) IsObjectNamespaced(obj runtime.Object) (bool, error) {
	//TODO implement me
	panic("implement me")
}
