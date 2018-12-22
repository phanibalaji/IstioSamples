# This sample demonstrates how to set timeouts on the services

1. Deploy the upstream node service using the below command.
```
kubectl apply -f upstreamnode.yaml
```
It deploys the version 1 of the upstream node service induces a fixed delay of 4 seconds on all the calls to this service.

2. Call the upstream node service as below and observe the delay
```
phanim-macbookpro:RequestTimeouts phanim$ kubectl exec -it sleep-79cc87b6b9-cxh6q -- time curl http://upstreamnode:8080/upstreamnode && echo ""

Defaulting container name to sleep.
Use 'kubectl describe pod/sleep-79cc87b6b9-cxh6q -n default' to see all of the containers in this pod.

{"Message":"Hello from Upstream Nodejs service","version":"1","Time":"2018-12-22T13:38:09.535Z","IP":"10.16.2.14"}
real	0m 4.01s
user	0m 0.00s
sys	0m 0.00s
```
We can observe that the call always takes slightly more than 4 seconds to respond.

3. Deploy the downstream GO service which call the upstream nodejs service. It sets the request timeout to 2 seconds on the requests to this service. 
```
kubectl apply -f downstreamgo.yaml
```

4. Call the downstream GO service as below
```
phanim-macbookpro:RequestTimeouts phanim$ kubectl exec -it sleep-79cc87b6b9-cxh6q -- time curl http://downstreamgo:8080/downstreamGo && echo ""

Defaulting container name to sleep.
Use 'kubectl describe pod/sleep-79cc87b6b9-cxh6q -n default' to see all of the containers in this pod.

upstream request timeout
real	0m 2.01s
user	0m 0.00s
sys	0m 0.00s
```
Since the request time is set to 2 seconds but the upstream takes 4 seconds to respond. So, all the calls to this service fail with the error `upstream request timeout`. The response takes slightly more than 2 seconds.

5. Increase the request timeout to a value more than 4 seconds and apply the confguration to get a successful response.

```
kubectl apply -f downstreamgo1.yaml

phanim-macbookpro:RequestTimeouts phanim$ kubectl exec -it sleep-79cc87b6b9-cxh6q -- time curl http://downstreamgo:8080/downstreamGo && echo ""

Defaulting container name to sleep.
Use 'kubectl describe pod/sleep-79cc87b6b9-cxh6q -n default' to see all of the containers in this pod.

{"Message":"Hello from Upstream Nodejs service","version":"1","Time":"2018-12-22T13:47:31.220Z","IP":"10.16.2.14"}
real	0m 4.02s
user	0m 0.00s
sys	0m 0.00s
```

