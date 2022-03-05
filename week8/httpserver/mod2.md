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
