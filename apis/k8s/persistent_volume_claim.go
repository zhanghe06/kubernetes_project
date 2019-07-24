package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListPersistentVolumeClaim(clientk8s *kubernetes.Clientset, ns string) (*v1.PersistentVolumeClaimList, error) {
	persistentVolumeClaimListRes, err := clientk8s.CoreV1().PersistentVolumeClaims(ns).List(metav1.ListOptions{})
	return persistentVolumeClaimListRes, err
}

func GetPersistentVolumeClaim(clientk8s *kubernetes.Clientset, ns string, name string) (*v1.PersistentVolumeClaim, error) {
	persistentVolumeClaimRes, err := clientk8s.CoreV1().PersistentVolumeClaims(ns).Get(name, metav1.GetOptions{})
	return persistentVolumeClaimRes, err
}

func UpdatePersistentVolumeClaim(clientk8s *kubernetes.Clientset, ns string, persistentVolumeClaim *v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error) {
	persistentVolumeClaimRes, err := clientk8s.CoreV1().PersistentVolumeClaims(ns).Update(persistentVolumeClaim)
	return persistentVolumeClaimRes, err
}
