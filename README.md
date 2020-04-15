# k8s-hello-go

docker build --tag k8shellogo:0.1 .
docker run --publish 8080:8080 --detach --name bb k8shellogo:0.1