package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListConfigMap(clientk8s *kubernetes.Clientset, ns string) (*v1.ConfigMapList, error) {
	configMapListRes, err := clientk8s.CoreV1().ConfigMaps(ns).List(metav1.ListOptions{})
	return configMapListRes, err
}

func GetConfigMap(clientk8s *kubernetes.Clientset, ns string, name string) (*v1.ConfigMap, error) {
	configMapRes, err := clientk8s.CoreV1().ConfigMaps(ns).Get(name, metav1.GetOptions{})
	return configMapRes, err
}

func UpdateConfigMap(clientk8s *kubernetes.Clientset, ns string, configMap *v1.ConfigMap) (*v1.ConfigMap, error) {
	configMapRes, err := clientk8s.CoreV1().ConfigMaps(ns).Update(configMap)
	return configMapRes, err
}
