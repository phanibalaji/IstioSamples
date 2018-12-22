# This sample illustrates the HTTP header based traffic routing mechanism in Istio

If the request payload contains HTTP header "version: v1" then the request goes to the version 1 of the nodejs service.
If the request payload contains HTTP header "version: v2" then the request goes to the version 2 of the nodejs service.
If this header is absent then the request goes to the version 1 by default.

1. Deploy the upstream node service using the below command.
```
kubectl apply -f upstreamNode.yaml
```
It deploys the two versions of the Nodejs service version 1 & version 2. 

2. Call the upstream node service as below
```
phanim-macbookpro:TrafficControl phanim$ kubectl exec -it sleep-79cc87b6b9-cxh6q -- curl http://upstreamnode:8080/upstreamNode && echo ""

phanim-macbookpro:TrafficRoutingBasedOnHeader phanim$ kubectl exec -it sleep-79cc87b6b9-cxh6q -- curl -H "version:v1" http://upstreamnode:8080/upstreamnode && echo ""
Defaulting container name to sleep.
Use 'kubectl describe pod/sleep-79cc87b6b9-cxh6q -n default' to see all of the containers in this pod.

{"Message":"Hello from Upstream Nodejs service","version":"1","Time":"2018-12-22T12:18:33.222Z","IP":"10.16.2.14"}

phanim-macbookpro:TrafficRoutingBasedOnHeader phanim$ kubectl exec -it sleep-79cc87b6b9-cxh6q -- curl -H "version:v2" http://upstreamnode:8080/upstreamnode && echo ""
Defaulting container name to sleep.
Use 'kubectl describe pod/sleep-79cc87b6b9-cxh6q -n default' to see all of the containers in this pod.

{"Message":"Hello from Upstream Nodejs service","version":"2","Time":"2018-12-22T12:18:44.828Z","IP":"10.16.1.12"}

phanim-macbookpro:TrafficRoutingBasedOnHeader phanim$ kubectl exec -it sleep-79cc87b6b9-cxh6q -- curl http://upstreamnode:8080/upstreamnode && echo ""
Defaulting container name to sleep.
Use 'kubectl describe pod/sleep-79cc87b6b9-cxh6q -n default' to see all of the containers in this pod.

{"Message":"Hello from Upstream Nodejs service","version":"1","Time":"2018-12-22T12:18:59.391Z","IP":"10.16.2.14"}
```
Observe how the traffic is routed to version1 or version 2 depending on the value of http header "version". 

