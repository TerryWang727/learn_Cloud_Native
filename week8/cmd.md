1、push 镜像
```
docker build --build-arg version=v0.0.1 . -t wxk-k8s/httpserver:0.0.1

docker push wxk-k8s/httpserver:0.0.1
```
2、创建 k8s 资源
```
kubectl apply -f deployment.yaml

kubectl logs -f -l app=cloudnative-httpserver
```
3、滚动更新
```
kubectl set image deployment/cloudnative-httpserver httpserver=wxk-k8s/httpserver:0.0.1

kubectl get events -w
```
