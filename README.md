# dummyauth

go mod init myproject

go mod tidy

go build -o app

go build -o app &&\
docker build -f Dockerfile -t myapp . &&\
docker tag myapp:latest localhost:32000/myapp:latest &&\
docker push localhost:32000/myapp:latest &&\
microk8s kubectl rollout restart deploy myapp -n mynamespace


microk8s kubectl apply -f myapp-deployment.yaml --namespace mynamespace
microk8s kubectl apply -f myapp-service.yaml --namespace mynamespace
microk8s kubectl apply -f myapp-ingress.yaml --namespace mynamespace
##
sudo vim /etc/hosts
127.0.0.1   myapp.localhost
ping myapp.localhost

http://myapp.localhost/drop

http://myapp.localhost/init

http://myapp.localhost/aplicaciones



curl -k http://myapp.localhost/aplicaciones

curl -k -X POST \
  -H "Content-Type: application/json" \
  -d '{"nombre":"app 1"}' \
  http://myapp.localhost/aplicaciones

curl -k -X POST \
  -H "Content-Type: application/json" \
  -d '{"nombre":"app 2"}' \
  http://myapp.localhost/aplicaciones

curl -k http://myapp.localhost/aplicaciones | jq

curl -k -X PUT \
  -H "Content-Type: application/json" \
  -d '{"id":"1"}' \
  http://myapp.localhost/aplicaciones

curl -k -X PUT \
  -H "Content-Type: application/json" \
  -d '{"id":"2"}' \
  http://myapp.localhost/aplicaciones

curl -k -X PUT \
  -H "Content-Type: application/json" \
  -d '{"id":"3"}' \
  http://myapp.localhost/aplicaciones

curl -k -X DELETE \
  -H "Content-Type: application/json" \
  -d '{"id":"3"}' \
  http://myapp.localhost/aplicaciones

curl -k -X DELETE \
  -H "Content-Type: application/json" \
  -d '{"id":"2"}' \
  http://myapp.localhost/aplicaciones

curl -k http://myapp.localhost/aplicaciones | jq