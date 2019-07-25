package extensions

import (
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListReplicaSet(clientk8s *kubernetes.Clientset, ns string) (*v1beta1.ReplicaSetList, error) {
	replicaSetListRes, err := clientk8s.ExtensionsV1beta1().ReplicaSets(ns).List(metav1.ListOptions{})
	return replicaSetListRes, err
}

func GetReplicaSet(clientk8s *kubernetes.Clientset, ns string, name string) (*v1beta1.ReplicaSet, error) {
	replicaSetRes, err := clientk8s.ExtensionsV1beta1().ReplicaSets(ns).Get(name, metav1.GetOptions{})
	return replicaSetRes, err
}

func UpdateReplicaSet(clientk8s *kubernetes.Clientset, ns string, replicaSet *v1beta1.ReplicaSet) (*v1beta1.ReplicaSet, error) {
	replicaSetRes, err := clientk8s.ExtensionsV1beta1().ReplicaSets(ns).Update(replicaSet)
	return replicaSetRes, err
}
