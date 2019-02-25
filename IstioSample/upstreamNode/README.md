# Deploy Node service

##  Prerequisites 

1. GKE cluster along with Istio is installed. Check the `IstioSample` directory for instructions.

## Instructions to deploy Node service on Istio

1. Run the below command to deploy the `upstreamnode-v1` pod, associated service and detsinationRule.
```
kubectl apply -f upstreamNode.yaml
```

2. Install a test pod `sleep` using the below command. The istio download has the sleep.yaml
``` 
kubectl apply -f <istio_unzipped_dir>/istio-1.0.6/samples/sleep/sleep.yaml
```

3. Run the following commands to list the containers of the pods
```
kubectl get namespace default -o yaml
kubectl get pods -n default
kubectl get pods upstreamnode-v1-65bf9d4c5d-92bcq -o jsonpath='{.spec.containers[*].name}' && echo ""
kubectl get pods sleep-7ffd5cc988-6ldh2 -o jsonpath='{.spec.containers[*].name}' && echo ""
```

4. Test the upstream Node service using the below command
```
kubectl exec -it sleep-7ffd5cc988-6ldh2 -- curl http://upstreamnode:8080/upstreamNode && echo ""

Defaulting container name to sleep.
Use 'kubectl describe pod/sleep-7ffd5cc988-6ldh2 -n default' to see all of the containers in this pod.
{"Message":"Hello from Upstream Nodejs service","version":"1","Time":"2018-12-19T09:47:12.749Z","IP":"10.16.1.9"}
```

5. View the individual container logs using the below commands
```
kubectl logs upstreamnode-v1-65bf9d4c5d-87bb6 -c upstreamnode
kubectl logs upstreamnode-v1-65bf9d4c5d-87bb6 -c istio-proxy
```

6. View the cluster/listener info from the upstreamnode pod using the below command
```
istioctl proxy-config cluster upstreamnode-v1-65bf9d4c5d-87bb6
istioctl proxy-config listener upstreamnode-v1-65bf9d4c5d-87bb6

```
