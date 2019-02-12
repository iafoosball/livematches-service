microk8s.kubectl delete services livematches

microk8s.kubectl create -f manifest.yaml

microk8s.docker push localhost:32000/livematches:latest

microk8s.docker build -t localhost:32000/livematches:latest .

microk8s.docker tag livematches localhost:32000/livematches

microk8s.kubectl port-forward livematches 8013:8013

microk8s.kubectl get all --all-namespaces

microk8s.kubectl apply -f manifest-all.yaml

