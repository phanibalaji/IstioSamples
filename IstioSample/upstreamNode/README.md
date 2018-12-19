The steps to install the upstream service running on Nodejs

1. Create a K8S cluster along with istio-beta. Select strict-MTLS for the security.
   `https://cloud.google.com/istio/docs/istio-on-gke/installing`

2. The docker image is already built and uploaded to phanibalaji/istio-sample-images at 
   `https://cloud.docker.com/repository/registry-1.docker.io/phanibalaji/istio-sample-images`

3. Run the below command to deploy the upstreamNode POD, associated service and detsinationRule.
   ```
   kubectl apply -f upstreamNode.yaml
   ```

4. Install a test POD `sleep` using the below command
   ``` 
   kubectl -f apply samples/sleep/sleep.yaml
   ```
5. Test the upstream Node service using the below command
   ```
   kubectl exec -it sleep-79cc87b6b9-4h6zr -- curl http://upstreamnode:8080/upstreamNode  

   Defaulting container name to sleep.
   Use 'kubectl describe pod/sleep-79cc87b6b9-4h6zr -n default' to see all of the containers in this pod.
   {"Message":"Hello from Upstream Nodejs service","version":"1","Time":"2018-12-19T09:47:12.749Z","IP":"10.16.1.9"}
   ```



