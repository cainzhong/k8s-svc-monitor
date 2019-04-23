package k8s

import "testing"

func TestGetK8sSvc(t *testing.T) {
	namespace := "itsma1"
	GetK8sSvc(namespace)
}
