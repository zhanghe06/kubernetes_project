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

func CreateMongodb(clientKubedb *kubedb.Clientset, ns string) (*kubedbV1.MongoDB, error)  {
	// MongoDB
	templatePath := "templates/tpl_mongodb.yaml"
	// 获取模板 TODO getTemplate
	templateContent, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(templateContent))
	// 模板解析 TODO parseTemplate

	var mongodb kubedbV1.MongoDB
	var jsonContent []byte
	jsonContent, err = yaml.YAMLToJSON(templateContent)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(jsonContent))
	err = json.Unmarshal(jsonContent, &mongodb)
	if err != nil {
		return nil, err
	}

	var mongodbDatabase *kubedbV1.MongoDB
	mongodbInterface := clientKubedb.KubedbV1alpha1().MongoDBs(ns)
	mongodbDatabase, err = mongodbInterface.Create(&mongodb)
	return mongodbDatabase, err
}

func GetMongodb(clientKubedb *kubedb.Clientset, ns string, name string) (*kubedbV1.MongoDB, error) {
	mongodbInterface := clientKubedb.KubedbV1alpha1().MongoDBs(ns)
	var mongodbDatabase *kubedbV1.MongoDB
	mongodbDatabase, err := mongodbInterface.Get(name, metav1.GetOptions{})
	return mongodbDatabase, err
}

func UpdateMongodb(clientKubedb *kubedb.Clientset, ns string, mongodb *kubedbV1.MongoDB) (*kubedbV1.MongoDB, error) {
	mongodbInterface := clientKubedb.KubedbV1alpha1().MongoDBs(ns)
	mongodbDatabase, err := mongodbInterface.Update(mongodb)
	return mongodbDatabase, err
}

func DeleteMongodb(clientKubedb *kubedb.Clientset, ns string, name string) error {
	mongodbInterface := clientKubedb.KubedbV1alpha1().MongoDBs(ns)
	err := mongodbInterface.Delete(name, nil)
	return err
}

func ListMongoDBVersions(clientKubedb *kubedb.Clientset) (*catalogV1.MongoDBVersionList, error) {
	mongoDBVersionInterface := clientKubedb.CatalogV1alpha1().MongoDBVersions()
	mongoDBVersionList, err := mongoDBVersionInterface.List(metav1.ListOptions{})
	return mongoDBVersionList, err
}
