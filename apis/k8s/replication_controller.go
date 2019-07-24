package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListReplicationController(clientk8s *kubernetes.Clientset, ns string) (*v1.ReplicationControllerList, error) {
	podListRes, err := clientk8s.CoreV1().ReplicationControllers(ns).List(metav1.ListOptions{})
	return podListRes, err
}

func GetReplicationController(clientk8s *kubernetes.Clientset, ns string, name string) (*v1.ReplicationController, error) {
	podRes, err := clientk8s.CoreV1().ReplicationControllers(ns).Get(name, metav1.GetOptions{})
	return podRes, err
}

func UpdateReplicationController(clientk8s *kubernetes.Clientset, ns string, pod *v1.ReplicationController) (*v1.ReplicationController, error) {
	podRes, err := clientk8s.CoreV1().ReplicationControllers(ns).Update(pod)
	return podRes, err
}
