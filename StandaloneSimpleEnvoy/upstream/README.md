* `docker pull https://hub.docker.com/r/envoyproxy/envoy/`
   This pulls envoy image which is based on Ubuntu 16.04.5 LTS. This has also envoy installed at `/usr/local/bin/envoy`

* Start the container on the above image. Login into the container
   `docker exec -it 9bf4341e8903 /bin/bash`

* While container is running install all the required 
   ```
   Install Node.js v8.x
   https://github.com/nodesource/distributions/blob/master/README.md
   ```

* Copy the contents of the `upstreamNode.js, package.json, upstream.yaml, startService.sh` to the running container under 
   directory say `nodetest`

* Test by running `startService.sh`

* Commit the Docker container if everything is working fine
   ```
   https://docs.docker.com/engine/reference/commandline/commit/#parent-command
   docker commit <container-id> <name_of_the_image>
   ```

* You can build another image based on the image created in the previous step. Using the below contents for the Dockerfile

  ```
  FROM <name_of_the_image_step_5>
  WORKDIR /nodetest
  EXPOSE 80
  ENTRYPOINT /nodetest/startService.sh
  ```
        
 * Build the image with the below command

    ```
    docker build -t NodeEnvoyFinal .
    ```

* Run the container with the below command
  ```
  docker run -p 49161:80 -d NodeEnvoyFinal
  ```

* Test the node from the host machine
   ```
   curl -i http://$(hostname -i):49161/upstreamNode && echo ""
   ```

* If you hit any WARNINGs (or) errors while starting the container like below, follow the instructions in the stackoverflow to enable the IP4 forwarding
   ```
   WARNING: IPv4 forwarding is disabled. Networking will not work.
   https://stackoverflow.com/questions/41453263/docker-networking-disabled-warning-ipv4-forwarding-is-disabled-networking-wil
   ```







