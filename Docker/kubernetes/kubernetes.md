# Kubernetes
* A alternate orchestrator

## Terminology:
### Kubernates 
* The orchestration system. Commonly called K8s.

### Kubectl
* CLI tool to talk to k8s API so as to config and manage k8s
* Stands for "cube control".

### Node
* Single server in the k8s cluster
* Sometimes, they are also call "worker"

### Kubelet
* Kubernetes agent running on a node for workers
* Use `kube-proxy` for network

### Control Plane
* The containers that manages k8s cluster (e.g. manager)
* Make global decisions about the cluster
* Includes
  * etcd: key-value store for k8s to store cluster data
  * API server: Expose k8s API
  * scheduler: Assign pods to node
  * controller manager: Ensure that the k8s cluster is in correct state
      * Node Controller: Responsible when nodes go down
      * Replication Controller: Correct number of pods
      * Endpoints Controller: Populates endpoints object
      * Service Account / Token Controller: Create account and API access token for namespaces
  * coreDNS
  * more...

## K8s Container Abstractions
### Pod
* One or more containers running on 1 node
* Basic unit of deployment

### Controller
* For creating/updating pods and other objects

### Service
* Network endpoint assigned to pod for connection

### Namespace
* Filtered group of objects in cluster (e.g. filter out system containers when using cli)

## Commands
### Create Pod
`kubectl run` Only for pod creation (similar to `docker run`) \
`kubectl create` Create resources via CLI or YAML (similar to `docker service create`) \
`kubectl apply` Create/update anything via YAML (similar to `docker stack deploy`) 

# Source
* [Components](https://kubernetes.io/docs/concepts/overview/components/)
