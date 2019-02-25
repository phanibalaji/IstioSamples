# Application architecture 

![Service Interaction](https://github.com/phanibalaji/IstioSamples/blob/master/StandaloneSimpleEnvoy/Diagram.png)

# Install ISTIO on GKE cluster

## Prerequisites 

1. Update `gcloud` command line tool
```
gcloud components update
```

2. If you have not installed `kubectl` as part of gcloud, install it using the below command
```
gcloud components install kubectl
```

3. Configure the project and zone for `gcloud`
```
gcloud config set project apigeeopdksupport-155404
gcloud config set compute/zone asia-northeast1-a
```

## Create GKE Cluster and configure

1. Create the GKE Cluster using the below command
```
gcloud container clusters create phani-cluster-1 --machine-type=n1-standard-1 --num-nodes=4 
```

2. Verify that cluster is created successfully 
```
kubectl get nodes
```

3. Run the below commands to get credentials and to provide admin permissions to the user
```
gcloud container clusters get-credentials phani-cluster-1 

kubectl create clusterrolebinding cluster-admin-binding --clusterrole=cluster-admin --user=$(gcloud config get-value core/account)
```

## Install istio and verify 

1. Download istio from the location https://github.com/istio/istio/releases. Download and unzip the files.
```
https://github.com/istio/istio/releases/download/1.0.6/istio-1.0.6-osx.tar.gz
```

2. Install istio using the below command
```
cd istio-1.0.6
kubectl apply -f install/kubernetes/istio-demo-auth.yaml
```

3. Verify that Istio pods are installed
```
kubectl get service -n istio-system
kubectl get pods -n istio-system
```

4. Enable automatic injection of pods to `default` namespace
```
kubectl label namespace default istio-injection=enabled
kubectl get namespace default -o yaml
```

## Viewing logs of istio components 

1. Obtain the ControlZ port of the istio component using the below command. We can connect to any istio component like `pilot`, `policy` or `telemetry` etc.
```
kubectl logs istio-pilot-878499dd5-689pt -c discovery -n istio-system | grep ControlZ

2019-02-25T06:16:57.922416Z	info	ControlZ available at 10.40.0.9:9876
```

2. Use the port forward to connect to ControlZ using the below command and then open the browser and point to `http://localhost:9876`
```
kubectl port-forward pods/istio-pilot-878499dd5-689pt  -n istio-system 9876:9876
``` 

