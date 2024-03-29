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
	//serviceName := "redis-quickstart"
	serviceName := "redis-cluster"

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

	// create redis
	//redisDatabase, err := apiKubedb.CreateRedis(clientKubedb, ns)
	//if err != nil {
	//	panic(err.Error())
	//}
	//utils.ShowYaml(redisDatabase)

	// get redis
	// kubedb get rd -n demo redis-quickstart -o yaml
	redisDatabase, err := apiKubedb.GetRedis(clientKubedb, ns, serviceName)
	if err != nil {
		panic(err.Error())
	}
	utils.ShowYaml(redisDatabase)

	// update redis
	/*
	 * terminationPolicy
	 * -----------------
	 * DoNotTerminate	禁止删除
	 * Pause (Default)	删除实例、保留数据、保留密码
	 * Delete			删除实例、删除数据、保留密码
	 * WipeOut			删除实例、删除数据、删除密码
	 */
	redisDatabase.Spec.TerminationPolicy = "WipeOut"
	redisDatabaseNew, err := apiKubedb.UpdateRedis(clientKubedb, ns, redisDatabase)
	if err != nil {
		panic(err.Error())
	}
	utils.ShowYaml(redisDatabaseNew)

	// delete redis
	// kubedb delete rd redis-cluster -n demo
	err = apiKubedb.DeleteRedis(clientKubedb, ns, serviceName)
	if err != nil {
		panic(err.Error())
	}
}
