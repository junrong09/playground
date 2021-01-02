# swarm
* Orchestration tool for managing of containers deployed access multiple host machines.

# Commands
## Swarm
* Manage swarm network

### Init:
`docker swarm init --advertise-addr <ip_addr>`
* --Advertise-addr: Set for members of swarm to access API and overlay network

### Join Token:
`docker swarm join-token (worker/manager)`
* Get the join token/command that can be run on non-member VM to join the swarm


## Node
* Manage node(s) in swarm network
* Once manager and VMs are in swarm, can just command the swarm in the manager node
### Update:
`docker node update --role manager <hostname>`

### ps:
`docker node ps <hostname>`
* Use `service ps` to look at all the processes by the service for a complete picture

### ls:
`docker node ls`


## Service
* Manage swarm services (i.e. containers replicas)
### Create:
`docker service create --name <name> --network <network> --replica <int> --secret <key> <image>`
* Similar to docker run OPTIONS
* Use overlay network (e.g. docker network create --driver overlay my_network_name) to create virtual subnet that connects all nodes in the swarm (even if they are in a different vm)

### Update:
`docker service update --replicas <int> --mount <mount_string> --secret-add/rm <secret> <service>`
* A service is updated by removing its containers and re-running with new ones.
* Use `docker service update --force <service>` to rebalance replicas (e.g. when new vm is added and they have no tasks).

### ps:
`docker service ps <service>`

### Remove:
`docker service rm <service>`

## Secret
* A way to encrypt data that is shared by services in a swarm (e.g. env variable)
* Support string/binary. Can be use by docker-compose, but it is not encrypted.
* Format similar to key-value pair.
* Can be access via `/run/secrets/<key>`, be default
* Example for usage: `docker service create --name psql --secret psql_user --secret psql_pass -e POSTGRES_USER_FILE=/run/secrets/psql_user -e POSTGRES_PASSWORD_FILE=/run/secrets/psql_pass postgres` and the secret can also be access by/within the container via the filepath.

### create:
`docker secret create <key> <file_with_value>` or `echo "value" | docker secret create <key> -`

### ls:
`docker secret ls`


# Docker Machine
* Tool to manage and install docker engine on remote/local VM. (e.g. mass install docker on mutiple cloud servers to prepare for the running of containers)
* Use commonly for automate dev/test environment.
* Alternatively, use [get.docker.com](https://get.docker.com) script to install.
