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

func CreatePostgres(clientKubedb *kubedb.Clientset, ns string) (*kubedbV1.Postgres, error) {
	// Postgres
	templatePath := "templates/tpl_postgres.yaml"
	// 获取模板 TODO getTemplate
	templateContent, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(templateContent))
	// 模板解析 TODO parseTemplate

	var postgres kubedbV1.Postgres
	var jsonContent []byte
	jsonContent, err = yaml.YAMLToJSON(templateContent)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(jsonContent))
	err = json.Unmarshal(jsonContent, &postgres)
	if err != nil {
		return nil, err
	}

	var postgresDatabase *kubedbV1.Postgres
	postgresInterface := clientKubedb.KubedbV1alpha1().Postgreses(ns)
	postgresDatabase, err = postgresInterface.Create(&postgres)
	return postgresDatabase, err
}

func GetPostgres(clientKubedb *kubedb.Clientset, ns string, name string) (*kubedbV1.Postgres, error) {
	postgresInterface := clientKubedb.KubedbV1alpha1().Postgreses(ns)
	var postgresDatabase *kubedbV1.Postgres
	postgresDatabase, err := postgresInterface.Get(name, metav1.GetOptions{})
	return postgresDatabase, err
}

func UpdatePostgres(clientKubedb *kubedb.Clientset, ns string, postgres *kubedbV1.Postgres) (*kubedbV1.Postgres, error) {
	postgresInterface := clientKubedb.KubedbV1alpha1().Postgreses(ns)
	postgresDatabase, err := postgresInterface.Update(postgres)
	return postgresDatabase, err
}

func DeletePostgres(clientKubedb *kubedb.Clientset, ns string, name string) error {
	postgresInterface := clientKubedb.KubedbV1alpha1().Postgreses(ns)
	err := postgresInterface.Delete(name, nil)
	return err
}

func ListPostgresVersions(clientKubedb *kubedb.Clientset) (*catalogV1.PostgresVersionList, error) {
	postgresVersionInterface := clientKubedb.CatalogV1alpha1().PostgresVersions()
	postgresVersionList, err := postgresVersionInterface.List(metav1.ListOptions{})
	return postgresVersionList, err
}
