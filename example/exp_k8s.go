package main

import (
	"encoding/json"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
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

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Verify operator installation
	ns := "kube-system"
	ls := "app=kubedb"
	pods, err := clientset.CoreV1().Pods(ns).List(metav1.ListOptions{LabelSelector: ls})
	if err != nil {
		panic(err.Error())
	}

	podsIndentBytes, err := json.MarshalIndent(pods, "", "\t")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(podsIndentBytes))
}
