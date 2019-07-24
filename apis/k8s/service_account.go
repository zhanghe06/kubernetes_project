package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListServiceAccount(clientk8s *kubernetes.Clientset, ns string) (*v1.ServiceAccountList, error) {
	serviceListRes, err := clientk8s.CoreV1().ServiceAccounts(ns).List(metav1.ListOptions{})
	return serviceListRes, err
}

func GetServiceAccount(clientk8s *kubernetes.Clientset, ns string, name string) (*v1.ServiceAccount, error) {
	serviceAccountRes, err := clientk8s.CoreV1().ServiceAccounts(ns).Get(name, metav1.GetOptions{})
	return serviceAccountRes, err
}

func UpdateServiceAccount(clientk8s *kubernetes.Clientset, ns string, serviceAccount *v1.ServiceAccount) (*v1.ServiceAccount, error) {
	serviceAccountRes, err := clientk8s.CoreV1().ServiceAccounts(ns).Update(serviceAccount)
	return serviceAccountRes, err
}
