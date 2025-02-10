#! /bin/bash

echo "!Iniciado minikube"

status_output=$(minikube status)

if echo "$status_output" | grep -q "host: Running"; then
    echo "Minikube está corriendo."
else
    minikube start
fi


echo "!Corriendo kluster de minikuebe"
echo "kubectl apply -f k8/namespace.yml"
kubectl apply -f k8/namespace.yml
echo "kubectl apply -f k8/"
kubectl apply -f k8/

echo "!FIn =)"