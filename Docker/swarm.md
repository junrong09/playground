# swarm
* Orchestration tool for managing of containers deployed access multiple host machines.

# Commands
## Swarm
* Manage swarm network

### Init:
`docker swarm init`


## Node
* Manage node(s) in swarm network

## Service
* Manage swarm services (i.e. containers replicas)
### Create:
`docker service create <image> --replica <int>`

### Update:
`docker service update <service> --replicas <int>`

### ps:
`docker service ps <service>`

### Remove:
`docker service rm <service>`



# Docker Machine
* Tool to manage and install docker engine on remote/local VM. (e.g. mass install docker on mutiple cloud servers to prepare for the running of containers)
* Use commonly for automate dev/test environment.
* Alternatively, use [get.docker.com](https://get.docker.com) script to install.
