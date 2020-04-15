# k8s-hello-go

- get google api project

```
gcloud config get-value project
```

- build container to gcr repo

```
gcloud builds submit --tag gcr.io/testing-saja-262603/helloworld-gke .
```

- Creating kubernetes zone

```
gcloud container clusters create helloworld-gke \
   --num-nodes 1 \
   --enable-basic-auth \
   --issue-client-certificate \
   --zone asia-southeast1

   #get nodes
   kubectl get nodes
```

- Deploy cluster
```
kubectl apply -f k8s/deployment.yaml
# Track the status of the Deployment:
kubectl get deployments
# The pods status
kubectl get pods
```



- Deploy load balancer service
```
kubectl apply -f k8s/service.yaml
# get service external ip
kubectl get services
# The pods status
kubectl get pods
```

for delete 
```
# Delete the container
gcloud container clusters delete k8s-hello-go-gke --zone asia-southeast1
# Delete the images
gcloud container images delete gcr.io/testing-saja-262603/k8s-hello-go-gke
```
