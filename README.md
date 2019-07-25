# Kubernetes Operator 实践

通过 K8S Operator 创建基础组件服务

支持单实例和集群两种模式

目前支持下列组件

- [X] ElasticSearch
- [X] MemCached
- [X] MongoDB
- [X] MySQL
- [X] PostgresQL
- [X] Redis
- [ ] RabbitMQ


项目演示
```
# 配置代理
export http_proxy=http://127.0.0.1:1087
export https_proxy=http://127.0.0.1:1087

# 配置依赖
go get k8s.io/client-go/kubernetes

go get github.com/kubedb/apimachinery/apis/kubedb/v1alpha1
go get github.com/imdario/mergo

# 配置环境
govendor init
govendor sync
```


获取 Kubeconfig
```
kubectl config view
cat .kube/config
```

获取 CRD 可用版本
```
kubectl get elasticsearchversions
kubectl get memcachedversions
kubectl get mongodbversions
kubectl get mysqlversions
kubectl get postgresversions
kubectl get redisversions
```

TODO
- 私有镜像仓库
- 模板解析


参考资料：

[https://github.com/kubernetes/client-go](https://github.com/kubernetes/client-go)

[https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/client-go](https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/client-go)

[kubedb 示例模板](https://github.com/kubedb/cli/tree/0.12.0/docs/examples)

[kubedb 示例指导](https://kubedb.com/docs/0.12.0/guides/)
