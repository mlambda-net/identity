
kubectl create secret docker-registry regdocker --docker-server=docker.pkg.github.com --docker-username=RoyGI --docker-password=$DOCKERPASS


kubectl create namespace user
kubectl config set-context --current --namespace=user

helm install ./helm/userserver --generate-name

helm install data ./helm/userdata

kubectl expose deployment userdata-1603069798 --type=LoadBalancer --port=5432

echo $DOCKERPASS | docker login https://docker.pkg.github.com -u USERNAME --password-stdin


kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.41.2/deploy/static/provider/do/deploy.yaml

kubectl delete -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.41.2/deploy/static/provider/do/deploy.yaml


cat ~/.kube/config  | base64 | tr -d \\n | pbcopy
cat ~/config | base64 --decode > ~/config.yaml


