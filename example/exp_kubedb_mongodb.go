package main

import (
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	"path/filepath"

	apiK8s "github.com/zhanghe06/kubernetes_project/apis/k8s"
	apiKubedb "github.com/zhanghe06/kubernetes_project/apis/kubedb"
	"github.com/zhanghe06/kubernetes_project/utils"

	kubedb "github.com/kubedb/apimachinery/client/clientset/versioned"
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

	ns := "demo"
	serviceName := "mongodb-quickstart"
	//serviceName := "mongodb-cluster"

	// create the k8s clientset
	clientK8s, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// get namespace
	namespaceRes, err := apiK8s.GetNamespace(clientK8s, ns)
	if err != nil {
		if errors.IsNotFound(err) {
			//create namespace
			namespaceRes, err = apiK8s.AddNamespace(clientK8s, ns)
			if err != nil {
				panic(err.Error())
			}
		} else {
			panic(err.Error())
		}
	}
	fmt.Println(namespaceRes.Name)

	// create the kubedb clientset
	clientKubedb, err := kubedb.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// create mongodb
	//mongodbDatabase, err := apiKubedb.CreateMongodb(clientKubedb, ns)
	//if err != nil {
	//	panic(err.Error())
	//}
	//utils.ShowYaml(mongodbDatabase)

	// get mongodb
	// kubedb get mg -n demo mongodb-quickstart -o yaml
	mongodbDatabase, err := apiKubedb.GetMongodb(clientKubedb, ns, serviceName)
	if err != nil {
		panic(err.Error())
	}
	utils.ShowYaml(mongodbDatabase)

	// update mongodb
	/*
	 * terminationPolicy
	 * -----------------
	 * DoNotTerminate	禁止删除
	 * Pause (Default)	删除实例、保留数据、保留密码
	 * Delete			删除实例、删除数据、保留密码
	 * WipeOut			删除实例、删除数据、删除密码
	 */
	mongodbDatabase.Spec.TerminationPolicy = "WipeOut"
	mongodbDatabaseNew, err := apiKubedb.UpdateMongodb(clientKubedb, ns, mongodbDatabase)
	if err != nil {
		panic(err.Error())
	}
	utils.ShowYaml(mongodbDatabaseNew)

	// delete mongodb
	// kubedb delete mg mongodb-cluster -n demo
	err = apiKubedb.DeleteMongodb(clientKubedb, ns, serviceName)
	if err != nil {
		panic(err.Error())
	}
}
