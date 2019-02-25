# The steps to deploy the upstream service running in Nodejs runtime on GKE. 

1. Create a K8S cluster along with istio-beta. Select strict-MTLS for the security.
`https://cloud.google.com/istio/docs/istio-on-gke/installing`

2. The docker image is already built and uploaded to phanibalaji/istio-sample-images at 
`https://cloud.docker.com/repository/registry-1.docker.io/phanibalaji/istio-sample-images`
Otherwise, if you want to build the image by yourself follow the steps in the section __Steps to build upstream node image__ and comeback after completing all the steps.

3. Run the below command to deploy the `upstreamnode-v1` POD, associated service and detsinationRule.
```
kubectl apply -f upstreamNode.yaml
```

4. Install a test POD `sleep` using the below command. The istio download has the sleep.yaml
``` 
kubectl apply -f <istio_unzipped_dir>/istio-1.0.6/samples/sleep/sleep.yaml
```

5. Run the following command to list the containers of the pods
```
kubectl get namespace default -o yaml
kubectl get pods -n default
kubectl get pods upstreamnode-v1-65bf9d4c5d-92bcq -o jsonpath='{.spec.containers[*].name}' && echo ""
kubectl get pods sleep-7ffd5cc988-6ldh2 -o jsonpath='{.spec.containers[*].name}' && echo ""
```

6. Test the upstream Node service using the below command
```
kubectl exec -it sleep-7ffd5cc988-6ldh2 -- curl http://upstreamnode:8080/upstreamNode  

Defaulting container name to sleep.
Use 'kubectl describe pod/sleep-7ffd5cc988-6ldh2 -n default' to see all of the containers in this pod.
{"Message":"Hello from Upstream Nodejs service","version":"1","Time":"2018-12-19T09:47:12.749Z","IP":"10.16.1.9"}
```


# Steps to build upstream node image
1. Run the below command to build the image. The Docker file is already provided
```
docker build -t <docker_hub_username>/<your_repo-name> .
```

2. Run the container with the below command
```
docker run -p 49160:8080 -d <docker_hub_username>/<your_repo-name>
curl http://localhost:49160/upstreamNode
```
   
3. If test is successful then tag the image and upload to the appropriate your docker hub repository.
Mention the image location in the `upstreamNode.yaml`
```
image: phanibalaji/istio-sample-images:upstream-node-v1
```
