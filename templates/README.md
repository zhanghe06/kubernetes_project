# Templates

[https://kubedb.com/docs/0.12.0/guides](https://kubedb.com/docs/0.12.0/guides)

[https://github.com/kubedb/cli/tree/0.12.0/docs/examples](https://github.com/kubedb/cli/tree/0.12.0/docs/examples)


```
➜  ~ kubectl get po -n demo
NAME                                  READY   STATUS    RESTARTS   AGE
elasticsearch-quickstart-0            1/1     Running   1          62m
memcached-quickstart-8ddb8966-8hk96   1/1     Running   1          62m
memcached-quickstart-8ddb8966-c9jdp   1/1     Running   1          62m
memcached-quickstart-8ddb8966-n8tnl   1/1     Running   1          62m
mongodb-quickstart-0                  1/1     Running   2          62m
mysql-quickstart-0                    1/1     Running   0          62m
postgres-quickstart-0                 1/1     Running   1          62m
redis-quickstart-0                    1/1     Running   1          62m
```


## elasticsearch
```
kubectl get secrets -n demo elasticsearch-quickstart-auth -o yaml
kubectl get secrets -n demo elasticsearch-quickstart-auth -o jsonpath='{.data.\ADMIN_USERNAME}' | base64 -D
kubectl get secrets -n demo elasticsearch-quickstart-auth -o jsonpath='{.data.\ADMIN_PASSWORD}' | base64 -D

kubectl port-forward -n demo elasticsearch-quickstart-0 9200
curl --user "admin:sbhbga47" "localhost:9200/_cluster/health?pretty"
```


## mysql

单实例
```
kubectl get secrets -n demo mysql-quickstart-auth -o yaml
kubectl get secrets -n demo mysql-quickstart-auth -o jsonpath='{.data.\username}' | base64 -D
kubectl get secrets -n demo mysql-quickstart-auth -o jsonpath='{.data.\password}' | base64 -D

➜  ~ kubectl exec -it mysql-quickstart-0 -n demo sh                                                                                                                                                              ➜  ~ kubectl exec -it mysql-quickstart-0 -n demo sh
# mysql -uroot -p
Enter password:
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 10
Server version: 8.0.14 MySQL Community Server - GPL

Copyright (c) 2000, 2019, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql>
```

集群模式
```
kubectl get secrets -n demo mysql-cluster-auth -o yaml
kubectl get secrets -n demo mysql-cluster-auth -o jsonpath='{.data.\username}' | base64 -D
kubectl get secrets -n demo mysql-cluster-auth -o jsonpath='{.data.\password}' | base64 -D

➜  ~ kubectl get pods -n demo -l kubedb.com/name=mysql-cluster
NAME              READY   STATUS    RESTARTS   AGE
mysql-cluster-0   1/1     Running   0          15m
mysql-cluster-1   1/1     Running   0          11m
mysql-cluster-2   1/1     Running   0          10m

➜  ~ kubectl get service -n demo
NAME                TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
kubedb              ClusterIP   None            <none>        <none>     6d
mysql-cluster       ClusterIP   10.106.121.64   <none>        3306/TCP   24m
mysql-cluster-gvr   ClusterIP   None            <none>        3306/TCP   24m

➜  ~ kubectl get pods -n demo -l kubedb.com/name=mysql-cluster -o jsonpath='{range.items[*]}{.metadata.name} ........... {.status.podIP} ............ {.metadata.name}.mysql-cluster-gvr.{.metadata.namespace}{"\n"}{end}'
mysql-cluster-0 ........... 10.1.14.170 ............ mysql-cluster-0.mysql-cluster-gvr.demo
mysql-cluster-1 ........... 10.1.14.171 ............ mysql-cluster-1.mysql-cluster-gvr.demo
mysql-cluster-2 ........... 10.1.14.172 ............ mysql-cluster-2.mysql-cluster-gvr.demo

➜  ~ kubectl exec -it -n demo mysql-cluster-0 -- mysql -u root --password=UlEUwzZw14gseQEA --host=mysql-cluster-0.mysql-cluster-gvr.demo -e "select 1;"
mysql: [Warning] Using a password on the command line interface can be insecure.
+---+
| 1 |
+---+
| 1 |
+---+

➜  ~ kubectl exec -it -n demo mysql-cluster-0 -- mysql -u root --password=UlEUwzZw14gseQEA --host=mysql-cluster-0.mysql-cluster-gvr.demo -e "SELECT * FROM performance_schema.replication_group_members;"
mysql: [Warning] Using a password on the command line interface can be insecure.
+---------------------------+--------------------------------------+----------------------------------------+-------------+--------------+
| CHANNEL_NAME              | MEMBER_ID                            | MEMBER_HOST                            | MEMBER_PORT | MEMBER_STATE |
+---------------------------+--------------------------------------+----------------------------------------+-------------+--------------+
| group_replication_applier | 1aa8b8f7-a90f-11e9-bd5a-d26dab1e2134 | mysql-cluster-0.mysql-cluster-gvr.demo |        3306 | ONLINE       |
| group_replication_applier | 380a71ea-a90f-11e9-8505-b29d1dc29a88 | mysql-cluster-1.mysql-cluster-gvr.demo |        3306 | ONLINE       |
| group_replication_applier | 4fe5726f-a90f-11e9-9178-dec66283e82d | mysql-cluster-2.mysql-cluster-gvr.demo |        3306 | ONLINE       |
+---------------------------+--------------------------------------+----------------------------------------+-------------+--------------+

➜  ~ kubectl exec -it -n demo mysql-cluster-0 -- mysql -u root --password=UlEUwzZw14gseQEA --host=mysql-cluster-0.mysql-cluster-gvr.demo -e "show status like '%primary%';"
mysql: [Warning] Using a password on the command line interface can be insecure.
+----------------------------------+--------------------------------------+
| Variable_name                    | Value                                |
+----------------------------------+--------------------------------------+
| group_replication_primary_member | 1aa8b8f7-a90f-11e9-bd5a-d26dab1e2134 |
+----------------------------------+--------------------------------------+
```


## postgres
```
kubectl get secrets -n demo postgres-quickstart-auth -o yaml
kubectl get secrets -n demo postgres-quickstart-auth -o jsonpath='{.data.\POSTGRES_USER}' | base64 -D
kubectl get secrets -n demo postgres-quickstart-auth -o jsonpath='{.data.\POSTGRES_PASSWORD}' | base64 -D

➜  ~ kubectl exec -it postgres-quickstart-0 -n demo sh
/ # psql -U postgres
psql (10.2)
Type "help" for help.

postgres=# \l
                                 List of databases
   Name    |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges
-----------+----------+----------+------------+------------+-----------------------
 postgres  | postgres | UTF8     | en_US.utf8 | en_US.utf8 |
 template0 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
 template1 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
(3 rows)

postgres=#
```


## redis
```
➜  ~ kubectl exec -it redis-quickstart-0 -n demo sh
/data # redis-cli
127.0.0.1:6379> CONFIG GET dir
1) "dir"
2) "/data"
127.0.0.1:6379>
```
