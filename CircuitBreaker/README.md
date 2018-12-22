# This sample demonstrates how to configure circuit breaker on a service

1. Deploy the upstream node service using the below command.
```
kubectl apply -f upstreamnode.yaml
```
It deploys the version 1 of the upstream node service. I does NOT configure any timeouts or any other circuit breaking settings.

2. Install the fortio pod from the istio installation zip file. Set the value of `FORTIO_POD` shell variable to the name of fortio pod.

```
kubectl -f apply samples/httpbin/sample-client/fortio-deploy.yaml
FORTIO_POD=$(kubectl get pod | grep fortio | awk '{ print $1 }')
 ```

3. Call the upstream node service as below.
```
kubectl exec -it $FORTIO_POD  -c fortio /usr/local/bin/fortio -- load -c 2 -qps 0 -n 20 -loglevel Warning http://upstreamnode:8080/upstreamNode
14:23:51 I logger.go:97> Log level is now 3 Warning (was 2 Info)
Fortio 1.0.1 running at 0 queries per second, 1->1 procs, for 20 calls: http://upstreamnode:8080/upstreamNode
Starting at max qps with 2 thread(s) [gomax 1] for exactly 20 calls (10 per thread + 0)
Ended after 23.147776ms : 20 calls. qps=864.01
Aggregated Function Time : count 20 avg 0.0023015114 +/- 0.0008063 min 0.001613425 max 0.00447507 sum 0.046030228
# range, mid point, percentile, count
>= 0.00161342 <= 0.002 , 0.00180671 , 55.00, 11
> 0.002 <= 0.003 , 0.0025 , 85.00, 6
> 0.003 <= 0.004 , 0.0035 , 90.00, 1
> 0.004 <= 0.00447507 , 0.00423754 , 100.00, 2
# target 50% 0.00196134
# target 75% 0.00266667
# target 90% 0.004
# target 99% 0.00442756
# target 99.9% 0.00447032
Sockets used: 2 (for perfect keepalive, would be 2)
Code 200 : 20 (100.0 %)
Response Header Sizes : count 20 avg 238 +/- 0 min 238 max 238 sum 4760
Response Body/Total Sizes : count 20 avg 352 +/- 0 min 352 max 352 sum 7040
All done 20 calls (plus 0 warmup) 2.302 ms avg, 864.0 qps
```
We can observe that all the calls are successful.

4. Now apply the circuit breaking settings 
```
kubectl apply -f upstreamnode_circuitbreaker.yaml.yaml
```

5. Call the upstream node service as below.
```
kubectl exec -it $FORTIO_POD  -c fortio /usr/local/bin/fortio -- load -c 2 -qps 0 -n 20 -loglevel Warning http://upstreamnode:8080/upstreamNode
14:22:43 I logger.go:97> Log level is now 3 Warning (was 2 Info)
Fortio 1.0.1 running at 0 queries per second, 1->1 procs, for 20 calls: http://upstreamnode:8080/upstreamNode
Starting at max qps with 2 thread(s) [gomax 1] for exactly 20 calls (10 per thread + 0)
14:22:43 W http_client.go:604> Parsed non ok code 503 (HTTP/1.1 503)
14:22:43 W http_client.go:604> Parsed non ok code 503 (HTTP/1.1 503)
14:22:43 W http_client.go:604> Parsed non ok code 503 (HTTP/1.1 503)
14:22:43 W http_client.go:604> Parsed non ok code 503 (HTTP/1.1 503)
14:22:43 W http_client.go:604> Parsed non ok code 503 (HTTP/1.1 503)
14:22:43 W http_client.go:604> Parsed non ok code 503 (HTTP/1.1 503)
14:22:43 W http_client.go:604> Parsed non ok code 503 (HTTP/1.1 503)
14:22:43 W http_client.go:604> Parsed non ok code 503 (HTTP/1.1 503)
14:22:43 W http_client.go:604> Parsed non ok code 503 (HTTP/1.1 503)
14:22:43 W http_client.go:604> Parsed non ok code 503 (HTTP/1.1 503)
14:22:43 W http_client.go:604> Parsed non ok code 503 (HTTP/1.1 503)
14:22:43 W http_client.go:604> Parsed non ok code 503 (HTTP/1.1 503)
14:22:43 W http_client.go:604> Parsed non ok code 503 (HTTP/1.1 503)
14:22:43 W http_client.go:604> Parsed non ok code 503 (HTTP/1.1 503)
14:22:43 W http_client.go:604> Parsed non ok code 503 (HTTP/1.1 503)
14:22:43 W http_client.go:604> Parsed non ok code 503 (HTTP/1.1 503)
Ended after 39.421057ms : 20 calls. qps=507.34
Aggregated Function Time : count 20 avg 0.0037331368 +/- 0.005255 min 0.00050929 max 0.016738594 sum 0.074662735
# range, mid point, percentile, count
>= 0.00050929 <= 0.001 , 0.000754645 , 40.00, 8
> 0.001 <= 0.002 , 0.0015 , 70.00, 6
> 0.002 <= 0.003 , 0.0025 , 80.00, 2
> 0.009 <= 0.01 , 0.0095 , 85.00, 1
> 0.012 <= 0.014 , 0.013 , 90.00, 1
> 0.016 <= 0.0167386 , 0.0163693 , 100.00, 2
# target 50% 0.00133333
# target 75% 0.0025
# target 90% 0.014
# target 99% 0.0166647
# target 99.9% 0.0167312
Sockets used: 17 (for perfect keepalive, would be 2)
Code 200 : 4 (20.0 %)
Code 503 : 16 (80.0 %)
Response Header Sizes : count 20 avg 47.75 +/- 95.5 min 0 max 239 sum 955
Response Body/Total Sizes : count 20 avg 244.15 +/- 54.3 min 217 max 353 sum 4883
All done 20 calls (plus 0 warmup) 3.733 ms avg, 507.3 qps
```
Observe that some calls failed due to circuit breaking settings.

