package kubedb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
	catalogV1 "github.com/kubedb/apimachinery/apis/catalog/v1alpha1"
	kubedbV1 "github.com/kubedb/apimachinery/apis/kubedb/v1alpha1"
	kubedb "github.com/kubedb/apimachinery/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateRedis(clientKubedb *kubedb.Clientset, ns string) (*kubedbV1.Redis, error) {
	// Redis
	templatePath := "templates/tpl_redis.yaml"
	// 获取模板 TODO getTemplate
	templateContent, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(templateContent))
	// 模板解析 TODO parseTemplate

	var redis kubedbV1.Redis
	var jsonContent []byte
	jsonContent, err = yaml.YAMLToJSON(templateContent)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(jsonContent))
	err = json.Unmarshal(jsonContent, &redis)
	if err != nil {
		return nil, err
	}

	var redisDatabase *kubedbV1.Redis
	redisInterface := clientKubedb.KubedbV1alpha1().Redises(ns)
	redisDatabase, err = redisInterface.Create(&redis)
	return redisDatabase, err
}

func DeleteRedis(clientKubedb *kubedb.Clientset, ns string, name string) error {
	redisInterface := clientKubedb.KubedbV1alpha1().Redises(ns)
	err := redisInterface.Delete(name, nil)
	return err
}

func ListRedisVersions(clientKubedb *kubedb.Clientset) (*catalogV1.RedisVersionList, error) {
	redisVersionInterface := clientKubedb.CatalogV1alpha1().RedisVersions()
	redisVersionList, err := redisVersionInterface.List(metav1.ListOptions{})
	return redisVersionList, err
}
