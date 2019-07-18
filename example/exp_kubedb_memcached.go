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
	//serviceName := "memcached-quickstart"
	serviceName := "memcached-cluster"

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

	// create memcached
	//memcachedDatabase, err := apiKubedb.CreateMemcached(clientKubedb, ns)
	//if err != nil {
	//	panic(err.Error())
	//}
	//utils.ShowYaml(memcachedDatabase)

	// get memcached
	// kubedb get mc -n demo memcached-quickstart -o yaml
	memcachedDatabase, err := apiKubedb.GetMemcached(clientKubedb, ns, serviceName)
	if err != nil {
		panic(err.Error())
	}
	utils.ShowYaml(memcachedDatabase)

	// update memcached
	/*
	 * terminationPolicy
	 * -----------------
	 * DoNotTerminate	禁止删除
	 * Pause (Default)	删除实例、保留数据、保留密码
	 * Delete			删除实例、删除数据、保留密码
	 * WipeOut			删除实例、删除数据、删除密码
	 */
	memcachedDatabase.Spec.TerminationPolicy = "WipeOut"
	memcachedDatabaseNew, err := apiKubedb.UpdateMemcached(clientKubedb, ns, memcachedDatabase)
	if err != nil {
		panic(err.Error())
	}
	utils.ShowYaml(memcachedDatabaseNew)

	// delete memcached
	// kubedb delete mc memcached-cluster -n demo
	err = apiKubedb.DeleteMemcached(clientKubedb, ns, serviceName)
	if err != nil {
		panic(err.Error())
	}
}
