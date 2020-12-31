# Basics Docker Command Rules
`docker <command> <subcommand>`

# Commands
## Containers
* For backwards compatibility, most of these command can be ran without the command `container` (e.g. `docker container run` can be ran as `docker run`)

### Run:
`docker run -p <listening_port>:<forward_port> -d --name webhost nginx:1.11 nginx -T`
`docker run -it mysql bash`
* -p/--publish: Publish containers port  
* -d/--detached: Detach and run conainter in background  
* --name: Set custom container name  
* -i: Interactive to keep stdin open
* -t: Open as a TTY prompt
* --rm: Remove container upon exit
* --network: Add virtual network (or use `docker network connect`)
* --network-alias: Add alias (in addition to one used by the container name)
* -v/--volume: Mount vol (e.g. named vol `NAME:DEST_PATH`, bind mount [starts with forward slash]`/SOMEPATH:DEST_PATH`)
* CMD: Start up command to run on the container (e.g. nginx -T / bash)  

### Exec:
`docker exec -it mysql bash`
* Run command in a already running container
* -it and bash is to run a interactive bash prompt in container

### Start: 
`docker start <container id/name>`
* Resume stopped container

### Stop:
`docker stop <container id/name>...`

### Rm:
`docker rm <container id/name>...`
* Cannot rm container that has not been stop, use -f

### ls:
`docker container ls -a` or `docker ps`

### Logs:
`docker logs <container id/name>`

### Top:
`docker top <container id/name>`
* Check the processes running in the container (e.g. `ps` command)

### Inspect:
`docker inspect <container id/name>`
* Check container config (e.g. network, mounts status)

### Stats:
`docker stats [container_id/name]`
* List conainers resource usage


## Network
* Manage container virtual networks

### ls:
`docker network ls`

### Inspect:
`docker network inspect <network>`
* Shows network config
* Containers in this virual network
* Has 3 default: 
  * Bridge (default virtual network)
  * Host (connect directly to host interface without NAT)
  * None (not attach to any network)

### Create:
`docker network create <network_name>`
* Create new virtual network
* Has DNS enabled. Containers within this virtual network can communication with each other just using container name. Don't even need ip addr as ip will change when redeployed. (e.g. `docker exec -it nginx ping nginx2` given that nginx and nginx2 are containers in the same virtual network)
* For the bridge (the default network) to have DNS enabled, need to use --link

### Connect:
`docker network connect <network> <container>`
* `--alias <alias_name>`: To add additional alias
* Connect container to additional virtual network

### Disconnect:
`docker network disconnect <network> <container>`

## Image
* Image are identifiable via tags in the format of `<org>/<image_name>:version`
* There are images without `<org`, e.g. official images
* By default, when version is not given it will use `latest`

### ls:
`docker image ls`

### History:
`docker history <image>`
* Shows the docker layers

### Inspect:
`docker inspect <image>`

### Tag:
`docker tag <image> <new_tag>`

### Push:
`docker image push <image>`
* Images pushed should be tagged with my dockerhub_name as org

### Build (Also part of Builder):
`docker build -t <tag> -f <file_name> <path>`
* Build a image from a dockerfile
* Without -f/--file, defaults to search in path for file named `Dockerfile`

## Volume
* 2 conventional types: Volume (Managed by docker) and bind mount (in OS filesystem)
* Volume can be specified in Dockerfile but not bind mount.

### ls
`docker volume ls`

### Inspect
`docker volume inspect`

### Create
`docker volume create <volume_name>`

## Prune
`docker container/volume/network/image prune` or `docker system prune`
* To remove dangling/unused container/volume/network/image \
* -a to remove all

`docker system df`
* To check disk usage


# MacOS Virtual Machine
* Docker containers run in a linux VM on Mac (instead of directly on MacOS)
* Run the following command to access the VM: `docker run -it --rm --privileged --pid=host justincormack/nsenter1`
[Docker for Mac Commands for Getting Into The Local Docker VM](https://www.bretfisher.com/docker-for-mac-commands-for-getting-into-local-docker-vm/)
