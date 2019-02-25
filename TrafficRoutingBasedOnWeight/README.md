# Weighted traffic control mechanism in Istio

1. Deploy the upstream Node service using the below command. It deploys the two versions of the Nodejs service version 1 & version 2.  Creates a virtual service which route 25% traffic to V1 and 75% to V2.

```
kubectl apply -f upstreamNode.yaml
```

2. Deploy the downstream GO service which call the upstream Node service.

```
kubectl apply -f downstreamGo.yaml
```

3. Call the downstream GO service as below. Closely 25% traffic will go to V1 and 75% will go to V2 if we continuously make calls to the service.

```
phanim-macbookpro:TrafficControl phanim$ kubectl exec -it sleep-79cc87b6b9-cxh6q -- curl http://downstreamgo:8080/downstreamGo && echo ""

Defaulting container name to sleep.
Use 'kubectl describe pod/sleep-79cc87b6b9-cxh6q -n default' to see all of the containers in this pod.
{"Message":"Hello from Upstream Nodejs service","version":"1","Time":"2018-12-20T11:55:33.632Z","IP":"10.16.2.14"}

phanim-macbookpro:TrafficControl phanim$ kubectl exec -it sleep-79cc87b6b9-cxh6q -- curl http://downstreamgo:8080/downstreamGo && echo ""

Defaulting container name to sleep.
Use 'kubectl describe pod/sleep-79cc87b6b9-cxh6q -n default' to see all of the containers in this pod.
{"Message":"Hello from Upstream Nodejs service","version":"2","Time":"2018-12-20T11:55:36.600Z","IP":"10.16.1.12"}
```

4. We can directly call upstream node service and observe the weight based traffic routing. 

```
kubectl exec -it sleep-79cc87b6b9-cxh6q -- curl http://upstreamnode:8080/upstreamnode && echo ""
```

5. The information about weights gets propogated across the ISTIO proxies running in various pods. So, the weighted behavior 
is observed consistently irrespective of how the service is called; whether the `upstreamnode` service is called directly or called through `downstreamgo` service.

