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
