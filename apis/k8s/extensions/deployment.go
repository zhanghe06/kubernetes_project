package extensions

import (
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListDeployment(clientk8s *kubernetes.Clientset, ns string) (*v1beta1.DeploymentList, error) {
	deploymentListRes, err := clientk8s.ExtensionsV1beta1().Deployments(ns).List(metav1.ListOptions{})
	return deploymentListRes, err
}

func GetDeployment(clientk8s *kubernetes.Clientset, ns string, name string) (*v1beta1.Deployment, error) {
	deploymentRes, err := clientk8s.ExtensionsV1beta1().Deployments(ns).Get(name, metav1.GetOptions{})
	return deploymentRes, err
}

func UpdateDeployment(clientk8s *kubernetes.Clientset, ns string, deployment *v1beta1.Deployment) (*v1beta1.Deployment, error) {
	deploymentRes, err := clientk8s.ExtensionsV1beta1().Deployments(ns).Update(deployment)
	return deploymentRes, err
}
