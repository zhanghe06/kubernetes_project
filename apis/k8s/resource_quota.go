package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListResourceQuota(clientk8s *kubernetes.Clientset, ns string) (*v1.ResourceQuotaList, error) {
	resourceQuotaListRes, err := clientk8s.CoreV1().ResourceQuotas(ns).List(metav1.ListOptions{})
	return resourceQuotaListRes, err
}

func GetResourceQuota(clientk8s *kubernetes.Clientset, ns string, name string) (*v1.ResourceQuota, error) {
	resourceQuotaRes, err := clientk8s.CoreV1().ResourceQuotas(ns).Get(name, metav1.GetOptions{})
	return resourceQuotaRes, err
}

func UpdateResourceQuota(clientk8s *kubernetes.Clientset, ns string, resourceQuota *v1.ResourceQuota) (*v1.ResourceQuota, error) {
	resourceQuotaRes, err := clientk8s.CoreV1().ResourceQuotas(ns).Update(resourceQuota)
	return resourceQuotaRes, err
}
