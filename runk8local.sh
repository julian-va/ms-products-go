#! /bin/bash

echo "!Iniciado minikube"

status_output=$(minikube status)

if echo "$status_output" | grep -q "host: Running"; then
    echo "Minikube está corriendo."
else
    minikube start
fi


echo "!Agregando docker a minikube"
echo "eval (minikube -p minikube docker-env)"
eval $(minikube -p minikube docker-env)

echo "!Creando imagen de docker"
echo "docker build -t ms-products ."
docker build -t ms-products .

echo "ya casi termina"
sleep 20

echo "!Corriendo kluster de minikuebe"
echo "kubectl apply -f k8/namespace.yml"
kubectl apply -f k8/namespace.yml
echo "kubectl apply -f k8/"
kubectl apply -f k8/

echo "!FIn =)"