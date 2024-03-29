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

	// kubectl get elasticsearchversions
	elasticsearchVersionList, err := apiKubedb.ListElasticsearchVersions(clientKubedb)
	if err != nil {
		panic(err.Error())
	}
	utils.ShowJson(elasticsearchVersionList)

	// kubectl get memcachedversions
	memcachedVersionList, err := apiKubedb.ListMemcachedVersions(clientKubedb)
	if err != nil {
		panic(err.Error())
	}
	utils.ShowJson(memcachedVersionList)

	// kubectl get mongodbversions
	mongoDBVersionList, err := apiKubedb.ListMongoDBVersions(clientKubedb)
	if err != nil {
		panic(err.Error())
	}
	utils.ShowJson(mongoDBVersionList)

	// kubectl get mysqlversions
	mySQLVersionList, err := apiKubedb.ListMySQLVersions(clientKubedb)
	if err != nil {
		panic(err.Error())
	}
	utils.ShowJson(mySQLVersionList)

	// kubectl get postgresversions
	postgresVersionList, err := apiKubedb.ListPostgresVersions(clientKubedb)
	if err != nil {
		panic(err.Error())
	}
	utils.ShowJson(postgresVersionList)

	// kubectl get redisversions
	redisVersionList, err := apiKubedb.ListRedisVersions(clientKubedb)
	if err != nil {
		panic(err.Error())
	}
	utils.ShowJson(redisVersionList)
}
