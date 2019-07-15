# kubernetes

[https://github.com/kubernetes/client-go](https://github.com/kubernetes/client-go)

[https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/client-go](https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/client-go)

[kubedb 示例模板](https://github.com/kubedb/cli/tree/0.12.0/docs/examples)

[kubedb 示例指导](https://kubedb.com/docs/0.12.0/guides/)

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
govendor add +e
```


获取 Kubeconfig
```
kubectl config view
cat .kube/config
```

查询可用版本
```
kubectl get mysqlversions
```

TODO
- 私有镜像仓库
- 模板解析
