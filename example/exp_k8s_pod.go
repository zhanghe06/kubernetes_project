package main

import (
	"encoding/json"
	"flag"
	"path/filepath"

	"github.com/zhanghe06/kubernetes_project/utils"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	// Verify operator installation
	ns := "kube-system"
	ls := "app=kubedb"
	pods, err := clientK8s.CoreV1().Pods(ns).List(metav1.ListOptions{LabelSelector: ls})
	if err != nil {
		panic(err.Error())
	}

	podsIndentBytes, err := json.MarshalIndent(pods, "", "\t")
	if err != nil {
		panic(err.Error())
	}
	utils.ShowJson(podsIndentBytes)
}
