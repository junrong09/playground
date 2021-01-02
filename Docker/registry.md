# Registry
* Host own private registry apart from using Docker Hub
* One option is to use `registry` a open-source registry

## Steps to host a registry:
1) Run as container `docker container run -d -p 5000:5000 --name registry -v $(pwd):registry registry` or run as swarm `docker service create -p 5000:5000 --name registry registry`

## Steps to store a image in private registry:
1) Tag that image `docker tag <container_name> <ip:port>/<container_name>`. Note: Using a ip and port instead of a docker hub ID.
1) Push the tagged image using its tag `docker push <ip:port>/<container_name>`

## Steps to use a image in private registry:
1) To use the image via pull `docker pull <ip:port>/<container_name>`
* Ensure that all nodes in the swarm have access to the image hosted in private registry.
