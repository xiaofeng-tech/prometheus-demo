# Introduce
This repo is a demo about prometheus. It includes:
-   Set up prometheus with docker contatiner, manually config targets, create recording and alert rule using yaml file.
-   Set up prometheus with prometheus-operator, add targets with crd in promethus-operator, also create recording and alert rule using same way.

## Project structrue
```
prometheus-demo             // here is prometheus demo project, it's including prometheus setting up and create a metric, recording and alert rules.
|- alerts-and-dashboard     // Mainly contains rule files
    |- prometheus           // Contains prometheus configuraiton and rule configs
    |- prometheus-operator  // Using crd of prometheus-operator to set target and create rules/
        |- my-example       // The example from my, contains go example app deployment, target configuration and rule configuration file
        |- operator-doc     // The example from prometheus-oprator offical document
|- cmd                      // Go example code, contains main fucniton, entry point.
|- monitoring               // Contains setting up prometheus with kube-prometheus and docker
    |- docker               // set up prometheus with docker
    |- prometheus-opearator // Set up kube-prometheus
|- src                      // go example application code
```

# Set up prometheus
## Requirments
-   Go vesion 1.9
-   Docker server and cli on your local machine
-   A kubernetes cluster supporting CRD and RBAC

## Set up prometheus with docker and run go app on local machine
First, set up prometheus. In `/monitoring/docker/`, There is a shell to run prometheus with docker using 9090 port. When running prometheus, will mount the directory `./alerts-and-dashboard/prometheus/` to `/etc/prometheus/` in prometheus container. Prometheues reads the configs from `/etc/prometheus/`.

Second, run go app on the local. Jump into `./cmd/`, then run the command `go run main.go`.

## Set up prometheus with prometheues-operator, also deploy go-app to kubernetes
First, create crd from prometheus-operator using command `kubectl create -f ./monitoring/prometheus-operator/setup/`.

Sceond, deploy prometheus, alertmanager, grafana and other componets to kubernetes.

Finally, deploy go app to kuberetes `kubectl apply -f ./alerts-and-dashboard/prometheus-operator/my-example/deployment.yaml` and deploy Service for go app `kubectl apply -f ./alerts-and-dashboard/prometheus-operator/my-example/service.yaml`, then deploy serviceMonitor, it can add a target for go app in prometheus `kubectl apply -f ./alerts-and-dashboard/prometheus-operator/my-example/serviceMonitor.yaml`, in the end, add alert rule `kubectl apply -f ./alerts-and-dashboard/prometheus-operator/my-example/prometheusRule.yaml`





