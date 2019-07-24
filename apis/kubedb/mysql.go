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

func CreateMysql(clientKubedb *kubedb.Clientset, ns string) (*kubedbV1.MySQL, error) {
	// MySQL
	//templatePath := "templates/tpl_mysql.yaml"
	templatePath := "templates/tpl_cluster_mysql.yaml"
	// 获取模板 TODO getTemplate
	templateContent, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(templateContent))
	// 模板解析 TODO parseTemplate

	var mysql kubedbV1.MySQL
	var jsonContent []byte
	jsonContent, err = yaml.YAMLToJSON(templateContent)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(jsonContent))
	err = json.Unmarshal(jsonContent, &mysql)
	if err != nil {
		return nil, err
	}

	var mysqlDatabase *kubedbV1.MySQL
	mysqlInterface := clientKubedb.KubedbV1alpha1().MySQLs(ns)
	mysqlDatabase, err = mysqlInterface.Create(&mysql)
	return mysqlDatabase, err
}

func GetMysql(clientKubedb *kubedb.Clientset, ns string, name string) (*kubedbV1.MySQL, error) {
	mysqlInterface := clientKubedb.KubedbV1alpha1().MySQLs(ns)
	var mysqlDatabase *kubedbV1.MySQL
	mysqlDatabase, err := mysqlInterface.Get(name, metav1.GetOptions{})
	return mysqlDatabase, err
}

func UpdateMysql(clientKubedb *kubedb.Clientset, ns string, mysql *kubedbV1.MySQL) (*kubedbV1.MySQL, error) {
	mysqlInterface := clientKubedb.KubedbV1alpha1().MySQLs(ns)
	mysqlDatabase, err := mysqlInterface.Update(mysql)
	return mysqlDatabase, err
}

func DeleteMysql(clientKubedb *kubedb.Clientset, ns string, name string) error {
	mysqlInterface := clientKubedb.KubedbV1alpha1().MySQLs(ns)
	err := mysqlInterface.Delete(name, nil)
	return err
}

func ListMySQLVersions(clientKubedb *kubedb.Clientset) (*catalogV1.MySQLVersionList, error) {
	mySQLVersionInterface := clientKubedb.CatalogV1alpha1().MySQLVersions()
	mySQLVersionList, err := mySQLVersionInterface.List(metav1.ListOptions{})
	return mySQLVersionList, err
}
