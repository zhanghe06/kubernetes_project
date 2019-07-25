package extensions

import (
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListDaemonSet(clientk8s *kubernetes.Clientset, ns string) (*v1beta1.DaemonSetList, error) {
	daemonSetListRes, err := clientk8s.ExtensionsV1beta1().DaemonSets(ns).List(metav1.ListOptions{})
	return daemonSetListRes, err
}

func GetDaemonSet(clientk8s *kubernetes.Clientset, ns string, name string) (*v1beta1.DaemonSet, error) {
	daemonSetRes, err := clientk8s.ExtensionsV1beta1().DaemonSets(ns).Get(name, metav1.GetOptions{})
	return daemonSetRes, err
}

func UpdateDaemonSet(clientk8s *kubernetes.Clientset, ns string, daemonSet *v1beta1.DaemonSet) (*v1beta1.DaemonSet, error) {
	daemonSetRes, err := clientk8s.ExtensionsV1beta1().DaemonSets(ns).Update(daemonSet)
	return daemonSetRes, err
}
