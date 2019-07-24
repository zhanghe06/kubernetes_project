package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListService(clientk8s *kubernetes.Clientset, ns string) (*v1.ServiceList, error) {
	serviceListRes, err := clientk8s.CoreV1().Services(ns).List(metav1.ListOptions{})
	return serviceListRes, err
}

func GetService(clientk8s *kubernetes.Clientset, ns string, name string) (*v1.Service, error) {
	serviceRes, err := clientk8s.CoreV1().Services(ns).Get(name, metav1.GetOptions{})
	return serviceRes, err
}

func UpdateService(clientk8s *kubernetes.Clientset, ns string, service *v1.Service) (*v1.Service, error) {
	serviceRes, err := clientk8s.CoreV1().Services(ns).Update(service)
	return serviceRes, err
}
