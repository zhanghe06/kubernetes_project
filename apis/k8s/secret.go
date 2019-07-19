package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetSecret(clientk8s *kubernetes.Clientset, ns string, name string) (*v1.Secret, error) {
	secretRes, err := clientk8s.CoreV1().Secrets(ns).Get(name, metav1.GetOptions{})
	return secretRes, err
}

func UpdateSecret(clientk8s *kubernetes.Clientset, ns string, secret *v1.Secret) (*v1.Secret, error) {
	secretRes, err := clientk8s.CoreV1().Secrets(ns).Update(secret)
	return secretRes, err
}
