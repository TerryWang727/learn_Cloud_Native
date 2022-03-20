1、使用istioctl analyzer 分析集群istio可能存在的问题
```
istioctl analyze -n dhtobb-hs
```
2、检查secret是否被正确加载
```
ki get pod

istioctl proxy-config secret istio-ingressgateway-78f69bd5db-wnspm.istio-system
```
3、检查ingress网关日志
```
istioctl proxy-config log {ingress-pod}.istio-system

istioctl dashboard envoy {ingress-pod}.istio-system
```

```
kubectl get namespace -L istio-injection #查看标签情况
kubectl label ns dhtobb-hs istio-injection- #移除命名空间的istio-injection标签
istioctl kube-inject -f specs/deployment.yaml -o  specs/deployment.istio.injected.yaml
```

```
# 自动注入的优先级
# 1.Pod 注解（优先级最高，如 Pod 中含注解 sidecar.istio.io/inject: "true/false"，则会被优先处理） 
# 2.neverInjectSelector  (ConfigMap istio-sidecar-injector)
# 3.alwaysInjectSelector (ConfigMap istio-sidecar-injector)
# 4.命名空间策略。
# 1 > 2 > 3 > 4
```
