# App run
```
make
docker build . --tag="rustamgainutdinov/microservice-architecture-usercrud:master"
helm install mysql -f ./conf/helm/mysql/values.yaml oci://registry-1.docker.io/bitnamicharts/mysql
k apply -f ./conf/k8s/*
```

## Mysql usage
```
MYSQL_ROOT_PASSWORD=$(kubectl get secret --namespace usercrud mysql -o jsonpath="{.data.mysql-root-password}" | base64 -d)
kubectl run mysql-client --rm --tty -i --restart='Never' --image  docker.io/bitnami/mysql:8.0.36-debian-12-r8 --namespace usercrud --env MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD --command -- bash
mysql -h mysql.usercrud.svc.cluster.local -uroot -p"$MYSQL_ROOT_PASSWORD"
```

## Local dev
```
make
eval $(minikube docker-env)
docker build -t usercrud:dev .



kind create cluster
kind load docker-image usercrud:dev
minikube image load usercrud:dev
```


## tmp
```
helm upgrade --install mysql bitnami/mysql --set auth.rootPassword=root --set primary.startupProbe.enabled=false --set primary.readinessProbe.enabled=false --set primary.livenessProbe.enabled=false --set secondary.replicaCount=0 --set primary.hostAliases=[mysql-release] --set enableMySQLX=true --set primary.extraPorts=3306
helm upgrade mysql oci://registry-1.docker.io/bitnamicharts/mysql -f ./conf/helm/mysql/values.yaml
```
https://iximiuz.com/en/posts/kubernetes-kind-load-docker-image/