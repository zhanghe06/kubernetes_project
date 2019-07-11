package main

import (
	"flag"
	"path/filepath"

	kubedbV1 "github.com/kubedb/apimachinery/apis/kubedb/v1alpha1"
	kubedb "github.com/kubedb/apimachinery/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

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
	clientkubedb, err := kubedb.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	ns := ""

	var mysql kubedbV1.MySQL
	mySQLInterface := clientkubedb.KubedbV1alpha1().MySQLs(ns)
	err = mySQLInterface.Delete(mysql.Name, &metav1.DeleteOptions{})
	_, err = mySQLInterface.Create(&mysql)
}
