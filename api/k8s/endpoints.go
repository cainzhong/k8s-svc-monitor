package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetAllEndpoints(clientset *kubernetes.Clientset, namespace string) (*v1.EndpointsList, error) {
	endpointsList, err := clientset.CoreV1().Endpoints(namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return endpointsList, nil
}

func GetEndpoints(clientset *kubernetes.Clientset, namespace, name string) (*v1.Endpoints, error) {
	endpointsList, err := clientset.CoreV1().Endpoints(namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	endpointsArr := endpointsList.Items
	for i := 0; i < len(endpointsArr); i++ {
		endpointsName := endpointsArr[i].Name
		if endpointsName == name {
			return &endpointsArr[i], nil
		}
	}
	return nil, nil
}
