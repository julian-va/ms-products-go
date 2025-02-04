#! /bin/bash

echo "!Iniciado minikube"

status_output=$(minikube status)

if echo "$status_output" | grep -q "host: Running"; then
    echo "Minikube está corriendo."
else
    minikube start
fi


echo "!Agregando docker a minikube"
eval $(minikube -p minikube docker-env)

echo "!Creando imagen de docker"
docker build -t ms-products .

sleep 10

echo "!Corriendo kluster de minikuebe"
kubectl apply -f k8/

echo "!FIn =)"