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

func CreateElasticsearch(clientKubedb *kubedb.Clientset, ns string) (*kubedbV1.Elasticsearch, error)  {
	// Elasticsearch
	//templatePath := "templates/tpl_elasticsearch.yaml"
	templatePath := "templates/tpl_cluster_elasticsearch.yaml"
	// 获取模板 TODO getTemplate
	templateContent, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(templateContent))
	// 模板解析 TODO parseTemplate

	var elasticsearch kubedbV1.Elasticsearch
	var jsonContent []byte
	jsonContent, err = yaml.YAMLToJSON(templateContent)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(jsonContent))
	err = json.Unmarshal(jsonContent, &elasticsearch)
	if err != nil {
		return nil, err
	}

	var elasticsearchDatabase *kubedbV1.Elasticsearch
	elasticsearchInterface := clientKubedb.KubedbV1alpha1().Elasticsearches(ns)
	elasticsearchDatabase, err = elasticsearchInterface.Create(&elasticsearch)
	return elasticsearchDatabase, err
}

func GetElasticsearch(clientKubedb *kubedb.Clientset, ns string, name string) (*kubedbV1.Elasticsearch, error) {
	elasticsearchInterface := clientKubedb.KubedbV1alpha1().Elasticsearches(ns)
	var elasticsearchDatabase *kubedbV1.Elasticsearch
	elasticsearchDatabase, err := elasticsearchInterface.Get(name, metav1.GetOptions{})
	return elasticsearchDatabase, err
}

func UpdateElasticsearch(clientKubedb *kubedb.Clientset, ns string, elasticsearch *kubedbV1.Elasticsearch) (*kubedbV1.Elasticsearch, error) {
	elasticsearchInterface := clientKubedb.KubedbV1alpha1().Elasticsearches(ns)
	elasticsearchDatabase, err := elasticsearchInterface.Update(elasticsearch)
	return elasticsearchDatabase, err
}

func DeleteElasticsearch(clientKubedb *kubedb.Clientset, ns string, name string) error {
	elasticsearchInterface := clientKubedb.KubedbV1alpha1().Elasticsearches(ns)
	err := elasticsearchInterface.Delete(name, nil)
	return err
}

func ListElasticsearchVersions(clientKubedb *kubedb.Clientset) (*catalogV1.ElasticsearchVersionList, error) {
	elasticsearchVersionInterface := clientKubedb.CatalogV1alpha1().ElasticsearchVersions()
	elasticsearchVersionList, err := elasticsearchVersionInterface.List(metav1.ListOptions{})
	return elasticsearchVersionList, err
}
