# Deploy Go service which makes calls to Node service.

##  Prerequisites 

1. GKE cluster along with Istio is installed. Check the `IstioSample` directory for instructions.

## Instructions to deploy GO service on Istio

1. Run the below command to deploy the `downstreamgo-v1` pod and service. Make sure that Node service already installed. Check the directory `upstreamNode` for instructions.
```
kubectl apply -f downstreamGo.yaml
```

2. Install a test pod `sleep` using the below command. The istio download has the `sleep.yaml`. Skip this command if the `sleep` pod is already installed.
```
kubectl -f apply <istio_unzipped_dir>/istio-1.0.6/samples/sleep/sleep.yaml
```

3. Run the following command to list the containers of the pods
```
kubectl get namespace default -o yaml
kubectl get pods -n default
kubectl get pods downstreamgo-v1-d8847485-9g84w -o jsonpath='{.spec.containers[*].name}' && echo ""
kubectl get pods sleep-7ffd5cc988-6ldh2 -o jsonpath='{.spec.containers[*].name}' && echo ""
```

4. Test the downstream Node service using the below command.
```
kubectl exec -it sleep-7ffd5cc988-6ldh2 -- curl http://downstreamgo:8080/downstreamGo && echo ""
Defaulting container name to sleep.
Use 'kubectl describe pod/sleep-7ffd5cc988-6ldh2 -n default' to see all of the containers in this pod.
{"Message":"Hello from Upstream Nodejs service","version":"1","Time":"2018-12-19T11:58:28.375Z","IP":"10.16.0.16"}
```