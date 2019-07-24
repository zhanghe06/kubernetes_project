package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListPod(clientk8s *kubernetes.Clientset, ns string) (*v1.PodList, error) {
	podListRes, err := clientk8s.CoreV1().Pods(ns).List(metav1.ListOptions{})
	return podListRes, err
}

func GetPod(clientk8s *kubernetes.Clientset, ns string, name string) (*v1.Pod, error) {
	podRes, err := clientk8s.CoreV1().Pods(ns).Get(name, metav1.GetOptions{})
	return podRes, err
}

func UpdatePod(clientk8s *kubernetes.Clientset, ns string, pod *v1.Pod) (*v1.Pod, error) {
	podRes, err := clientk8s.CoreV1().Pods(ns).Update(pod)
	return podRes, err
}
