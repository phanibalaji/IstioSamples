* `docker pull https://hub.docker.com/r/envoyproxy/envoy/`
   This pulls envoy image which is based on Ubuntu 16.04.5 LTS. This has also envoy installed at `/usr/local/bin/envoy`

* Start the container on the above image. Login into the container
   ```
   docker exec -it 9bf4341e8903 /bin/bash
   ```

* While container is running install GO 
   ```
   Install GO
   https://golang.org/doc/install
   ```

* Copy the contents of the `downstreamGo.go, downstream.yaml, startService.sh` to the running container under 
   directory say `gotest`. In the `downstream.yaml` file, the upstreamNode IP is hardcoded to `10.140.0.4`. 
   We need to change this to actual IP of the upstreamNode service. Compile the go program using the command.
   ```
   go build downstreamGo.go
   ```

* Test by running `startService.sh`

* Commit the Docker container if everything is wokring fine
   https://docs.docker.com/engine/reference/commandline/commit/#parent-command
   ```
   docker commit <container-id> <name_of_the_image>
   ```

* You can build another image based on the image created in the step 5. Using the below contents for `Dockerfile`
  ```
  FROM <name_of_the_image>
  WORKDIR /gotest
  EXPOSE 80
  ENTRYPOINT /gotest/startService.sh
  ```
  
* Build the Docker image
 ```
  docker build -t GoEnvoyFinal .
 ```

* Run the container with the below command
   ```
   docker run -p 49161:80 -d GoEnvoyFinal
   ```

* Test the node from the host machine
   ```
   curl -i http://$(hostname -i):49161/downstreamGo && echo ""
   ```
* If you hit any WARNINGs (or) errors while starting the container like below, follow the instructions in the stackoverflow to enable the IP4 forwarding
   
   ```
   WARNING: IPv4 forwarding is disabled. Networking will not work.
   https://stackoverflow.com/questions/41453263/docker-networking-disabled-warning-ipv4-forwarding-is-disabled-networking-wil
   ```

