###snap install helm --classic

###install nginx-ingress

```
helm repo add nginx-stable https://helm.nginx.com/stable
"nginx-stable" has been added to your repositories
root@wxk-k8s:/opt/homework# helm repo update
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "nginx-stable" chart repository
Update Complete. ⎈Happy Helming!⎈
root@wxk-k8s:/opt/homework# helm install my-release nginx-stable/nginx-ingress
NAME: my-release
LAST DEPLOYED: Sun Mar  6 00:24:14 2022
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
The NGINX Ingress Controller has been installed.
```

###install cert-manager

```
root@wxk-k8s:/opt/homework# helm install cert-manager jetstack/cert-manager --namespace cert-manager --create-namespace --version v1.7.1
NAME: cert-manager
LAST DEPLOYED: Sun Mar  6 00:31:10 2022
NAMESPACE: cert-manager
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
cert-manager v1.7.1 has been deployed successfully!

In order to begin issuing certificates, you will need to set up a ClusterIssuer
or Issuer resource (for example, by creating a 'letsencrypt-staging' issuer).

More information on the different types of issuers and how to configure them
can be found in our documentation:

https://cert-manager.io/docs/configuration/

For information on how to configure cert-manager to automatically provision
Certificates for Ingress resources, take a look at the `ingress-shim`
documentation:

https://cert-manager.io/docs/usage/ingress/
```

失败



换方式

### 部署ingress

```
kubectl apply -f nginx-ingress-deployment.yaml
```

```
namespace/ingress-nginx created
serviceaccount/ingress-nginx created
configmap/ingress-nginx-controller created
clusterrole.rbac.authorization.k8s.io/ingress-nginx created
clusterrolebinding.rbac.authorization.k8s.io/ingress-nginx created
role.rbac.authorization.k8s.io/ingress-nginx created
rolebinding.rbac.authorization.k8s.io/ingress-nginx created
service/ingress-nginx-controller-admission created
service/ingress-nginx-controller created
deployment.apps/ingress-nginx-controller created
ingressclass.networking.k8s.io/nginx created
validatingwebhookconfiguration.admissionregistration.k8s.io/ingress-nginx-admission created
serviceaccount/ingress-nginx-admission created
clusterrole.rbac.authorization.k8s.io/ingress-nginx-admission created
clusterrolebinding.rbac.authorization.k8s.io/ingress-nginx-admission created
role.rbac.authorization.k8s.io/ingress-nginx-admission created
rolebinding.rbac.authorization.k8s.io/ingress-nginx-admission created
job.batch/ingress-nginx-admission-create created
job.batch/ingress-nginx-admission-patch created
```

### 证书
```
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=httpserver.com/O=httpserver"
```

### tls.crt

```
cat tls.crt |base64
```

```
LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURQVENDQWlXZ0F3SUJBZ0lVRTUvOE5neTlOclMzbW40QXlMWUQrY3k0MEZvd0RRWUpLb1pJaHZjTkFRRUwKQlFBd0xqRVhNQlVHQTFVRUF3d09hSFIwY0hObGNuWmxjaTVqYjIweEV6QVJCZ05WQkFvTUNtaDBkSEJ6WlhKMgpaWEl3SGhjTk1qSXdNekExTVRjMU1USXpXaGNOTWpNd016QTFNVGMxTVRJeldqQXVNUmN3RlFZRFZRUUREQTVvCmRIUndjMlZ5ZG1WeUxtTnZiVEVUTUJFR0ExVUVDZ3dLYUhSMGNITmxjblpsY2pDQ0FTSXdEUVlKS29aSWh2Y04KQVFFQkJRQURnZ0VQQURDQ0FRb0NnZ0VCQU1TaUxOYkpEbllHS3lJYThhQ2lkelppbC81ZTV2ZWlUbWpvdkk1QQppTEExUkhzNDQ1VWJhS1VNQU1jc0YyTjYyVkZoMTZWM2dsVnFGMi8rRGpBZEVyVmtBTlV2WDgyTzE4SlRMc1JoCisyOUdQN0pwZnR2V0dJajNMdWRocEx2ZXAydS9kQVVpOVQwZW9JTnRNVkVPTVpMN2o4ODBOb0pQaWJqL255SXkKTGRVMmdJZlJpbGpwWGRFOWo5ckFBTnBtMmtkOE1LbTBOaEEzYXN6d2llYWdTZkFFNlV2UlRYSTQ0TFVVR2JXegpBaHdiWVJSZi9IejJ6Yk43UzlVSDJtR2RXeUEyRERZYXBCK0IrQ2tLM0JZQlJqRm5KYmFrT3lnbzlHczJHSjhCCkIyT1BCeFFkUVhqNkQ1djNQeHFuQU1BT0c5YTAyZEZueGdRN2doT1BpeTlsREpFQ0F3RUFBYU5UTUZFd0hRWUQKVlIwT0JCWUVGTFphSmhNc0FHY0o2WGFCeGM0R3RpU1VpTk1LTUI4R0ExVWRJd1FZTUJhQUZMWmFKaE1zQUdjSgo2WGFCeGM0R3RpU1VpTk1LTUE4R0ExVWRFd0VCL3dRRk1BTUJBZjh3RFFZSktvWklodmNOQVFFTEJRQURnZ0VCCkFCWkpleUFNOHgyTDgxckhycnRSL3FBcldUSEhrcFFHOThTd2hUanhKTllpd3JJanJTOGM1MHFleVhsamE4MWsKcnI2amNiSXA0dFd3bUxPQ2t1MStZZE95d01TYW4rWk9tWmxqWWpjMWdxZU9lR0VSS01LSi94bDJpZ0JmNnQxNgprTTd3RkhaTXVaL1dVL2ZVZEZVYXpkYWFheUU5V2U3ZnVSTXZxWjV1QWthbzFlN3ZDaWh3Rk92am11bzYybTRBCnlrWkNlc3lIQmhvc3NCR05Xb1N1elQvb1UyZnRXY1RiUmU3OHFPL25ONVZtV1poWXJ2aTNQL2ZsRnBNcjZsRk4KRzVaQUlMa29Uc2xtMzRwaEpkbDNQVzFkSTF1dXN3ZHJvUWdJSFJEVCttbndNRnR3aVhVNTFZaS9aY3grNTNibApBVUVWaWgvdVZ2Qmh1ZG1FLzFjUWxNdz0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
```

### tls.key

```
cat tls.key |base64
```

```
LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2UUlCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktjd2dnU2pBZ0VBQW9JQkFRREVvaXpXeVE1MkJpc2kKR3ZHZ29uYzJZcGYrWHViM29rNW82THlPUUlpd05VUjdPT09WRJpbERBREhMQmRqZXRsUllkZWxkNEpWYWhkdgovZzR3SFJLMVpBRFZMMS9OanRmQ1V5N0VZZnR2UmoreWFYN2IxaGlJOXk3bllhUzczcWRydjNRRkl2VTlIcUNECmJURlJEakdTKzQvUE5EYUNUNG00LzU4aU1pM1ZOb0NIMFlwWTZWM1JQWS9hd0FEYVp0cEhmRENwdERZUU4yck0KOElubW9FbndCT2xMMFUxeU9PQzFGQm0xc3dJY0cyRVVYL3g4OXMyemUwdlZCOXBoblZzZ05ndzJHcVFmZ2ZncApDdHdXQVVZeFp5VzJwRHNvS1BSck5oaWZBUWRqandjVUhVRjQrZytiOXo4YXB3REFEaHZXdE5uUlo4WUVPNElUCmo0c3ZaUXlSQWdNQkFBRUNnZ0VBSHFGb3ZMaXJNUVY2K0tBYjNic1JPUDZZMEw0MTJvT0ZGV0NUNXREMHVnWkMKT2JPOTkvOUNDTnZhenl2MkpNQUJWR1VhYzZNdlFXRDZiNjZ5eFJueVRTRmI0OW15WWRJTG1ERGZDeVo4UlhRNgovZTdqSFJtM3pyWTZyejZOcTlwZVdER0Z3RHNXMnNPSG01R0o2aWQ4MzNtaWlGT0kzam5IcVhMZ2RqSTg0NGFJCml3aGxjRHJuQlQxSFBRUktHVjZwSWVreDZUcEZ4bDlnZ2JTQnFFNDRnTkhlRUhSNTRpSGFwemlyMGdkNGlDN20KdnpsYU1FYlZRSytrNFZFdjl3MHo2MEJCVmhDYTNzV2oweHpRZTBmRjVmUTN6aENndUphRFIxVGozVTBMSy8yKwphNkdjL0ZBUWZwOVFGTHlZK1Y1V2QzZ1BUS0ZzZnMvVFlsVnUvTlllUVFLQmdRRDlRbC9HLzBDdjl4bktuaHBPCjcxbzFkTmdmTGxLRjBsRnVJUEI5QWdkMEROSW5wL3NXMjhqVmFWZVVKVzNQWFZueVc0ZHVYMVcyRGlndCtrWmQKbXlTSHk4SzlNVWdwS1J4YUpMckZ0Y0hiTmM1MkZ0WG1FYWEwQmY1cm05dTZPSkgrTTAvU2Fod1NJNFZMWk1RYwppaFdTQVUveXlRem9qMEphVDJLVzlINmxhUUtCZ1FER3d1ei8xN3RuajQ1MDNIZkhNT3E2aGpPSU5MQWdOdkxVCk1vSFBJcno1aFBRbER0SjM2NU8xRzZoaWZXZnFEd210anlSbVFSeHozS3Q0VGVjWUhWU3FzZzN2NWRyYmVHQ3YKK2pkZmEzTkZPYTUxUFl3MmFpZjdTa0d6aHZqYTVXZWxxeDAxdklRL2phb0hlQlVNNzFscEtoektPbW5WRFBOSQpYR3FCNEl5QTZRS0JnUUM5RTNDdUFLTXF2REJsTk9RdlNrTHU1ZXowV29FQnA2Tmp0UFBXWUUxd05xSXZPUm9jClZybjVta3ZvTG9sczV5VGY2N1dRa1Z4TUx3V1FUZWw2dVBqczZSTGFiSkNUS2o0Z2pvdGlUL0FLMklEcmFPRUkKZ3VVU1FlZFFMMnc3KzBBRHFBdFA3Q2hJNG53QmJabDhOUnNsaGhWS0RRKzJFRDVaYmlzTXlEeEZBUUtCZ0M5egpqZDB1aGNFZGxxYnprMnpza3IxaEdLQmw0NzV0SDkvbEJ2U0ttSThCWE1BVUg4OGRZTEFXSUVjVEpXSE5vVVBjCkxwWnk2UFlJTXErUCthSGFMc0pwcThZZ0cvWFZjVS9SN3JKTEZzUHFGMnBKL1ZWb1ZvODVsU0hsRVRoQkdGT0cKM0h4ZHV6em85elM5U0ZsRU14WldSWFZLS01ZQ1IzcDVCYnhuL1dNeEFvR0FZRmhteTF4NitJbjVTWHRyb0FsZgpmdEY0SUcxdlhsNnErd3hUN2lRcThzZ3ZuOXArSHR1cWtuWUtDMy9UN25PU2xoSVpmSWIvUGZoSVUzTWJwbG9ZCnBuTUtIZkJKR1psMFBqNjVJL09teDJ5cytJMmh0eXk2bzVDTWFZRU5ETUQzZ1NILzBsdDdoQVpxbnZHcW5qSjkKRWpTNXFBbzBBVEF1Q1Z6YkdPV2tnOW89Ci0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0K
```

### 部署
```
make deploy
```

### 设置localhost
```
sudo sed -i '$a 127.0.0.1 httpserver.com' /etc/hosts
```

###查看
```
kubectl get svc -n ingress-nginx
```

```
NAME                                 TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
ingress-nginx-controller             NodePort    10.97.196.117   <none>        80:30374/TCP,443:31849/TCP   7m26s
ingress-nginx-controller-admission   ClusterIP   10.109.3.84     <none>        443/TCP                      7m26s
```
这里EXTERNAL-IP为空  没解决

### 查看配置 由于端口映射 443-->30374
```
curl https://httpserver.com:30374/config
```

