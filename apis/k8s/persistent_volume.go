package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListPersistentVolume(clientk8s *kubernetes.Clientset) (*v1.PersistentVolumeList, error) {
	persistentVolumeListRes, err := clientk8s.CoreV1().PersistentVolumes().List(metav1.ListOptions{})
	return persistentVolumeListRes, err
}

func GetPersistentVolume(clientk8s *kubernetes.Clientset, name string) (*v1.PersistentVolume, error) {
	persistentVolumeRes, err := clientk8s.CoreV1().PersistentVolumes().Get(name, metav1.GetOptions{})
	return persistentVolumeRes, err
}

func UpdatePersistentVolume(clientk8s *kubernetes.Clientset, persistentVolume *v1.PersistentVolume) (*v1.PersistentVolume, error) {
	persistentVolumeRes, err := clientk8s.CoreV1().PersistentVolumes().Update(persistentVolume)
	return persistentVolumeRes, err
}
