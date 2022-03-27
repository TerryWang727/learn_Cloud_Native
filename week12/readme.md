```
把我们的 httpserver 服务以 Istio Ingress Gateway 的形式发布出来。以下是你需要考虑的几点：

如何实现安全保证；
七层路由规则；
考虑 open tracing 的接入。
```

make push

安装部署集群\
```
#部署
make deployment

#清理部署
make clean 
```


单独部署envoy进行测试
```
#将httpserver的deployment以及service部署好之后，部署envoy
kubectl create configmap envoy-config --from-file=envoy.yaml
kubectl create -f envoy-deploy.yaml
```

安装istio

```
- 添加istioctl到PATH路径
vi ~/.bash_profile
source ~/.bash_profile

- 添加istio的自动补全脚本
cp tools/istioctl.bash ~/.istioctl.bash
source ~/.istioctl.bash

- 查看istio的版本号
istioctl version --remote=false

- 安装install -- 确保网络插件 cni 正常安装
istioctl install --set profile=demo -y

- 查看安装情况
kubectl get svc -n istio-system

- 卸载istio
istioctl manifest generate --set profile=demo | kubectl delete -f -
```

安全保证
```
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
    - match:
        - port: 443
      route:
        - destination:
            host: dhtobb-httpserver-service
            port:
              number: 80

---
apiVersion: networking.istio.io/v1beta1
kind: Gateway # 对应 listener, 使istio增加对80端口的监听器， 然后根据VirtualService的规则做转发
    - hosts:
        - http.wxk.io
      port:
        name: https-wxk-port
        number: 443
        protocol: HTTPS
      tls:
        mode: SIMPLE #单向认证
        credentialName: wxk-credential 
```

七层路由规则
访问/not-delay时将被重新定向到/delay
```
      - match:
        - uri:
            exact: "/not-delay"
      rewrite:
        uri: "/delay"
      route:
          - destination:
              host: wxk-httpserver-service
              port:
                number: 80
```

open tracing 的接入
```
1.golang中向下发送header时需要转换小写，否则首字母大写导致tracing失效

2.tracing需要依赖sidecar
```
```

kubectl apply -f jaeger.yaml

-查看面板
istioctl dashboard jaeger --address=IPADDR --browser=false
```
Access the httpserver via ingress
```
curl --resolve httpsserver.cncamp.io:443:$INGRESS_IP https://httpsserver.cncamp.io/healthz -v -k
```


mark  tetrate
```
https://academy.tetrate.io/courses/istio-fundamentals-zh
```
