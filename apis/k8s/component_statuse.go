package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListComponentStatus(clientk8s *kubernetes.Clientset) (*v1.ComponentStatusList, error) {
	componentStatusListRes, err := clientk8s.CoreV1().ComponentStatuses().List(metav1.ListOptions{})
	return componentStatusListRes, err
}

func GetComponentStatus(clientk8s *kubernetes.Clientset, name string) (*v1.ComponentStatus, error) {
	componentStatusRes, err := clientk8s.CoreV1().ComponentStatuses().Get(name, metav1.GetOptions{})
	return componentStatusRes, err
}

func UpdateComponentStatus(clientk8s *kubernetes.Clientset, componentStatus *v1.ComponentStatus) (*v1.ComponentStatus, error) {
	componentStatusRes, err := clientk8s.CoreV1().ComponentStatuses().Update(componentStatus)
	return componentStatusRes, err
}
