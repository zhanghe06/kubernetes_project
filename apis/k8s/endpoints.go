package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListEndpoints(clientk8s *kubernetes.Clientset, ns string) (*v1.EndpointsList, error) {
	endpointsListRes, err := clientk8s.CoreV1().Endpoints(ns).List(metav1.ListOptions{})
	return endpointsListRes, err
}

func GetEndpoints(clientk8s *kubernetes.Clientset, ns string, name string) (*v1.Endpoints, error) {
	endpointsRes, err := clientk8s.CoreV1().Endpoints(ns).Get(name, metav1.GetOptions{})
	return endpointsRes, err
}

func UpdateEndpoints(clientk8s *kubernetes.Clientset, ns string, endpoints *v1.Endpoints) (*v1.Endpoints, error) {
	endpointsRes, err := clientk8s.CoreV1().Endpoints(ns).Update(endpoints)
	return endpointsRes, err
}
