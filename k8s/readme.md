
ubuntu 22.04

sudo snap install microk8s --classic

microk8s disable ha-cluster --force

microk8s stop

microk8s start

microk8s status

microk8s enable dns dashboard storage

microk8s enable ingress

# https://microk8s.io/docs/registry-built-in
microk8s enable registry:size=40Gi

--------------
microk8s kubectl get services --namespace kube-system | grep kubernetes-dashboard
kubernetes-dashboard        ClusterIP   10.152.183.196   <none>        443/TCP                  14m
https://10.152.183.196/
microk8s kubectl create token default
--------------

microk8s kubectl create namespace mynamespace
microk8s kubectl get namespaces | grep mynamespace

microk8s kubectl apply -f 11-postgres-pv.yaml --namespace mynamespace
microk8s kubectl apply -f 12-postgres-pvc.yaml --namespace mynamespace
microk8s kubectl apply -f 13-postgres-config.yaml --namespace mynamespace
microk8s kubectl apply -f 14-postgres-secret.yaml --namespace mynamespace
microk8s kubectl apply -f 15-postgres-deployment.yaml --namespace mynamespace
microk8s kubectl apply -f 16-postgres-service.yaml --namespace mynamespace

-------------------

microk8s kubectl apply -f 21-myapp-deployment.yaml --namespace mynamespace
microk8s kubectl apply -f 22-myapp-service.yaml --namespace mynamespace
microk8s kubectl apply -f 23-myapp-ingress.yaml --namespace mynamespace