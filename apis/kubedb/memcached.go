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

func CreateMemcached(clientKubedb *kubedb.Clientset, ns string) (*kubedbV1.Memcached, error)  {
	// Memcached
	templatePath := "templates/tpl_memcached.yaml"
	// 获取模板 TODO getTemplate
	templateContent, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(templateContent))
	// 模板解析 TODO parseTemplate

	var memcached kubedbV1.Memcached
	var jsonContent []byte
	jsonContent, err = yaml.YAMLToJSON(templateContent)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(jsonContent))
	err = json.Unmarshal(jsonContent, &memcached)
	if err != nil {
		return nil, err
	}

	var memcachedDatabase *kubedbV1.Memcached
	memcachedInterface := clientKubedb.KubedbV1alpha1().Memcacheds(ns)
	memcachedDatabase, err = memcachedInterface.Create(&memcached)
	return memcachedDatabase, err
}

func DeleteMemcached(clientKubedb *kubedb.Clientset, ns string, name string) error {
	memcachedInterface := clientKubedb.KubedbV1alpha1().Memcacheds(ns)
	err := memcachedInterface.Delete(name, nil)
	return err
}

func ListMemcachedVersions(clientKubedb *kubedb.Clientset) (*catalogV1.MemcachedVersionList, error) {
	memcachedVersionInterface := clientKubedb.CatalogV1alpha1().MemcachedVersions()
	memcachedVersionList, err := memcachedVersionInterface.List(metav1.ListOptions{})
	return memcachedVersionList, err
}
