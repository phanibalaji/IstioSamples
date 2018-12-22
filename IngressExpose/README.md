# This sample demonstrates how to expose a service through an external load balancer 

1. Deploy the upstream node service using the below command.
```
kubectl apply -f upstreamnode.yaml
```
It deploys the version 1 of the upstream node service along with a `upstreamnode` virtualservice that uses `bookinfo-gateway` gateway. Note that if you have already created a gateway on a port, you can not create another gateway on the same port.

Since I already created `bookinfo-gateway`, I reused the same. Othwerise, you can create a new one. 

2. Set the shell variables as below
```
export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
xport INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].port}')
export GATEWAY_URL=$INGRESS_HOST:$INGRESS_PORT
```
3. Make a call to upstreamNode service as below from anywhere. We need not connect to a POD to make this call as illustrated in other samples as this service is exposed through the load balancer on GKE.
```
curl http://$INGRESS_HOST:$INGRESS_PORT/upstreamNode
{"Message":"Hello from Upstream Nodejs service","version":"1","Time":"2018-12-22T15:02:10.003Z","IP":"10.16.2.14"}
```