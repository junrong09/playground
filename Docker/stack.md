# Swarm Stack
* The docker-compose for swarm (e.g. automate swarm config and services runtime)
* Uses docker-compose format with an addition of `deploy` tag. 
* Note: Docker-compose ignores `deploy` tag, while stack ignores `build` tag.  
* Build should not be done in swarm, it should be done in CI tool.

# Commands
## Stack
### Deploy:
`docker stack deploy -c <compose_file> <stack>`
* Note: To update the swarm stack, use the same command with the same `<stack>` name.

### ls:
`docker stack ls`

### ps:
`docker stack ps <stack>`
* The entry are actually swarm tasks (not containers, even though they get spin-up as container in a node) 

### service:
`docker stack service <stack>`
* Like `docker service ls`, but for stack

### rm:
`docker stack rm <stack>`
