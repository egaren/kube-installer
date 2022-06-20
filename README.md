#Kubernetes deployment tool
> Small elegant tool which deploys kubernetes cluster. Easy to use and useful for newbies or someone who wants to deploy
> Kubernetes cluster but does not want to understand deployment process in depth.
> 
> Currently cli tool is ready to deploy single node K3s cluster on same node where utility deployed.

## Features
- Automatically detects Architecture

## Deploy single node cluster
### Usage
```shell
sudo ./kube-installer deploy -f k3s -n 192.168.0.44 -v v1.23.7+k3s1
```

### Results
```shell
ubuntu@server5-vm:~$ sudo systemctl status k3s.service 
  k3s.service - Lightweight Kubernetes
     Loaded: loaded (/etc/systemd/system/k3s.service; enabled; vendor preset: enabled)
     Active: active (running) since Sun 2022-06-19 13:42:04 UTC; 13s ago
       Docs: https://k3s.io
```
```shell
ubuntu@server5-vm:~$ sudo kubectl get nodes
NAME         STATUS   ROLES                  AGE     VERSION
server5-vm   Ready    control-plane,master   9m23s   v1.23.7+k3s1

ubuntu@server5-vm:~$ sudo kubectl get pods -A
NAMESPACE     NAME                              READY   STATUS      RESTARTS   AGE
kube-system   coredns-d76bd69b-97frh            1/1     Running     0          8m52s
kube-system   helm-install-traefik-crd-j5zpf    0/1     Completed   0          8m52s
kube-system   metrics-server-7cd5fcb6b7-qspxx   1/1     Running     0          8m52s
kube-system   helm-install-traefik-595tl        0/1     Completed   1          8m52s
kube-system   svclb-traefik-m44nn               2/2     Running     0          8m2s
kube-system   traefik-df4ff85d6-hnvvf           1/1     Running     0          8m2s
```
## Upgrade single node cluster
### Usage
```shell
sudo ./kube-installer upgrade -f k3s -v v1.24.1+k3s1
```
### Results
```shell
ubuntu@server5-vm:~$ sudo systemctl status k3s
 k3s.service - Lightweight Kubernetes
     Loaded: loaded (/etc/systemd/system/k3s.service; enabled; vendor preset: enabled)
     Active: active (running) since Mon 2022-06-20 11:12:12 UTC; 22s ago
       Docs: https://k3s.io
```
```shell
ubuntu@server5-vm:~$ sudo kubectl get nodes
NAME         STATUS   ROLES                  AGE   VERSION
server5-vm   Ready    control-plane,master   21h   v1.24.1+k3s1

ubuntu@server5-vm:~$ sudo kubectl get pods -A
NAMESPACE     NAME                              READY   STATUS      RESTARTS   AGE
kube-system   helm-install-traefik-crd-j5zpf    0/1     Completed   0          21h
kube-system   helm-install-traefik-595tl        0/1     Completed   1          21h
kube-system   svclb-traefik-m44nn               2/2     Running     0          21h
kube-system   coredns-d76bd69b-97frh            1/1     Running     0          21h
kube-system   traefik-df4ff85d6-hnvvf           1/1     Running     0          21h
kube-system   metrics-server-7cd5fcb6b7-qspxx   1/1     Running     0          21h
```