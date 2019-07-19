package main

import (
	"flag"
	"fmt"
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
	secretName := "postgres-cluster-auth"

	secretRes, err := apiK8s.GetSecret(clientK8s, ns, secretName)
	if err != nil {
		panic(err.Error())
	}
	data := secretRes.Data
	utils.ShowJson(data) // Json 序列化时，[]byte 类型的值自动base64编码

	fmt.Println(string(data["POSTGRES_PASSWORD"])) // 明文

	data["POSTGRES_PASSWORD"] = []byte("EUq1vfI1PTRM3dXO") // 这里明文，不需要编码
	secretResNew, err := apiK8s.UpdateSecret(clientK8s, ns, secretRes)
	if err != nil {
		panic(err.Error())
	}
	utils.ShowJson(secretResNew.Data)
}
