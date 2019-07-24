package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListEvent(clientk8s *kubernetes.Clientset, ns string) (*v1.EventList, error) {
	eventListRes, err := clientk8s.CoreV1().Events(ns).List(metav1.ListOptions{})
	return eventListRes, err
}

func GetEvent(clientk8s *kubernetes.Clientset, ns string, name string) (*v1.Event, error) {
	eventRes, err := clientk8s.CoreV1().Events(ns).Get(name, metav1.GetOptions{})
	return eventRes, err
}

func UpdateEvent(clientk8s *kubernetes.Clientset, ns string, event *v1.Event) (*v1.Event, error) {
	eventRes, err := clientk8s.CoreV1().Events(ns).Update(event)
	return eventRes, err
}
