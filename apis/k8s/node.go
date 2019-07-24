package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListNode(clientk8s *kubernetes.Clientset) (*v1.NodeList, error) {
	nodeListRes, err := clientk8s.CoreV1().Nodes().List(metav1.ListOptions{})
	return nodeListRes, err
}

func GetNode(clientk8s *kubernetes.Clientset, name string) (*v1.Node, error) {
	nodeRes, err := clientk8s.CoreV1().Nodes().Get(name, metav1.GetOptions{})
	return nodeRes, err
}

func UpdateNode(clientk8s *kubernetes.Clientset, node *v1.Node) (*v1.Node, error) {
	nodeRes, err := clientk8s.CoreV1().Nodes().Update(node)
	return nodeRes, err
}
