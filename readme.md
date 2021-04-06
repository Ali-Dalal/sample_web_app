#Go Sample web app

This repo contains a go sample web app to demonstrate deployment of go application in k8s environment

# Requirements to run the application
- git cli
- kubectl cli
- helm chart
- minikube (for testing)

# Build the application
Docker image for this application is already available on docker registry
https://hub.docker.com/r/allloush92/travix_test

You can build the image locally using the following command

```shell
docker build -t {TAG} . //replace tag with desired value
```

# Running the application
by default, helm chart will pull the image directly from docker registry, you can modify `web_api/templates/deployment.yml` to use a local image if built locally

### install the application

to install the app, run the following command
```shell
helm install web web_api/
```

the previous command will install k8s templates that include:
- namespace
- deployment
- service

Configuration variables are written in `web_api/values.yaml` and can be updated on demand (number of replicas, the name of namespace, ..etc)

after installing the helm chart, if you query the number of available pods, `kubectl` will print the following

```shell
kubectl get pods -n travix-dev
```
![get pods](./screenshots/get_pods.png?raw=true "get_pods")

### test the application

if the installation process passes successfully, you can test the service using minikube by running the following cmd
```shell
minikube service web-travix-test-service -n travix-dev 
```
the previous command will produce the following and open the app in the browser
![minikube run service](./screenshots/minikube_run_service.png?raw=true "minikube run service")