package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListNamespace(clientk8s *kubernetes.Clientset) (*v1.NamespaceList, error) {
	namespaceListRes, err := clientk8s.CoreV1().Namespaces().List(metav1.ListOptions{})
	return namespaceListRes, err
}

func GetNamespace(clientk8s *kubernetes.Clientset, ns string) (*v1.Namespace, error) {
	namespaceRes, err := clientk8s.CoreV1().Namespaces().Get(ns, metav1.GetOptions{})
	return namespaceRes, err
}

func AddNamespace(clientk8s *kubernetes.Clientset, ns string) (*v1.Namespace, error) {
	namespaceReq := v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns}}
	namespaceRes, err := clientk8s.CoreV1().Namespaces().Create(&namespaceReq)
	return namespaceRes, err
}
