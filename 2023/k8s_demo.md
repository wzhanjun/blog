### k8s-demo

#### 构建镜像
docker build . -t aaasayok/hellok8s:v1
docker push aaasayok/hellok8s:v1

#### 创建pod
kubectl apply -f nginx.yaml
```
# nginx.yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx-pod
spec:
  containers:
    - name: nginx-container
      image: nginx

```

#### 查看pod
kubectl get pod
kubectl get pod -o wide

#### pod 端口转发
kubectl port-forward nginx-pod 8080:80

#### 进入pod容器
kubectl exec -it nginx-pod /bin/bash

#### 查看pod日志
kubectl logs --follow nginx-pod

#### 执行pod命令
kubectl exec nginx-pod -- ls

#### 删除pod
kubectl delete pod nginx-pod
kubectl delete -f nginx.yaml

#### 查看pod信息
kubectl describe pod hellok8s-deployment-5997b7bf96-wfhl5


#### deployment 回滚
kubectl rollout undo deployment hellok8s-deployment
kubectl rollout history deployment hellok8s-deployment
kubectl rollout undo deployment/hellok8s-deployment --to-revision=2

#### 查看endpoints
kubectl get endpoints

#### 查看 service
kubectl get service

#### 查看configmap
kubectl get configmap --all-namespaces

#### helm
helm install my-hello-helm hellok8s/hello-helm --version 0.1.0
helm list -A   # see argocd in namespace argocd
helm uninstall argocd -n argocd

helm upgrade --install hello-helm --values values.yaml .

helm ls
helm rollback hello-helm 1



https://github.com/guangzhengli/k8s-tutorials