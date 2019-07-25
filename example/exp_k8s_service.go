package main

import (
	"flag"
	"path/filepath"

	apiK8s "github.com/zhanghe06/kubernetes_project/apis/k8s"
	"github.com/zhanghe06/kubernetes_project/utils"

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

	// create the k8s clientset
	clientK8s, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	ns := "demo"
	serviceName := "kubedb"

	serviceListRes, err := apiK8s.ListService(clientK8s, ns)
	if err != nil {
		panic(err.Error())
	}
	utils.ShowJson(serviceListRes)

	serviceRes, err := apiK8s.GetService(clientK8s, ns, serviceName)
	if err != nil {
		panic(err.Error())
	}
	utils.ShowJson(serviceRes)
}
