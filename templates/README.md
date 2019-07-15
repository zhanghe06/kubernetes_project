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


elasticsearch
```
kubectl get secrets -n demo elasticsearch-quickstart-auth -o yaml
kubectl get secrets -n demo elasticsearch-quickstart-auth -o jsonpath='{.data.\ADMIN_USERNAME}' | base64 -D
kubectl get secrets -n demo elasticsearch-quickstart-auth -o jsonpath='{.data.\ADMIN_PASSWORD}' | base64 -D

kubectl port-forward -n demo elasticsearch-quickstart-0 9200
curl --user "admin:sbhbga47" "localhost:9200/_cluster/health?pretty"
```


mysql
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


postgres
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


redis
```
➜  ~ kubectl exec -it redis-quickstart-0 -n demo sh
/data # redis-cli
127.0.0.1:6379> CONFIG GET dir
1) "dir"
2) "/data"
127.0.0.1:6379>
```
