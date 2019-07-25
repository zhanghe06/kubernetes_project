package main

import (
	"flag"
	"fmt"
	"path/filepath"

	apiK8s "github.com/zhanghe06/kubernetes_project/apis/k8s"
	apiK8sExt "github.com/zhanghe06/kubernetes_project/apis/k8s/extensions"
	"github.com/zhanghe06/kubernetes_project/utils"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string
	if home := utils.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	ns := "kube-system"
	deploymentName := "kubedb-operator"

	// create the k8s clientset
	clientK8s, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// get namespace
	namespaceRes, err := apiK8s.GetNamespace(clientK8s, ns)
	if err != nil {
		if errors.IsNotFound(err) {
			fmt.Printf("namespace %s not found", ns)
		}
		panic(err.Error())
	}
	fmt.Println(namespaceRes.Name)

	// get kubedb deployment
	deployment, err := apiK8sExt.GetDeployment(clientK8s, ns, deploymentName)
	if err != nil {
		panic(err.Error())
	}
	utils.ShowJson(deployment)
}
