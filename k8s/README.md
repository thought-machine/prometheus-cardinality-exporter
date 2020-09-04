# K8s Files
These yamls are provided to give an idea of how this exporter can work within kubernetes clusters
## ```deployment.yaml```
Defines the deployment of the exporter. It specifies the replicas and resources required to run it as well as the http port to expose and the arguments to use.
## ```service.yaml```
Defines the kubernetes service, which is an abstraction away from the actual pods running the code. The service allows the pod(s) to be accessed without a fixed IP address.
## ```service-account.yaml```
Defines the service account for the service.
## ```clusterrole.yaml```
States the resources that service discovery is allowed to query (namespaces, endpoints) and with what method (list).
## ```clusterrolesbinding.yaml```
Binds the cluster role to the service account.
