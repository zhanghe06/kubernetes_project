package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListLimitRange(clientk8s *kubernetes.Clientset, ns string) (*v1.LimitRangeList, error) {
	limitRangeListRes, err := clientk8s.CoreV1().LimitRanges(ns).List(metav1.ListOptions{})
	return limitRangeListRes, err
}

func GetLimitRange(clientk8s *kubernetes.Clientset, ns string, name string) (*v1.LimitRange, error) {
	limitRangeRes, err := clientk8s.CoreV1().LimitRanges(ns).Get(name, metav1.GetOptions{})
	return limitRangeRes, err
}

func UpdateLimitRange(clientk8s *kubernetes.Clientset, ns string, limitRange *v1.LimitRange) (*v1.LimitRange, error) {
	limitRangeRes, err := clientk8s.CoreV1().LimitRanges(ns).Update(limitRange)
	return limitRangeRes, err
}
