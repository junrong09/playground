# Configuration YAML
* For k8s to config its cluster
* File contains 1 or more manifests
* Each manifest describes an API object (deployment, job, secret)
* Each manifest has 4 parts.
* For Documentation: [API Reference](https://kubernetes.io/docs/reference/#api-reference)

## Manifest Parts
### apiVersion
* To retrieve version list use `kubectl api-versions`.

### kind
* apiVersion + kind tells that exact resource the k8s should run. `kubectl api-resources` to get the list available on ur k8s.

### metadata
* A user-defined names such as labels.
* Can be used by commands as selector using the `-l` flag. e.g. `kubectl get pods -l <label_key>=<label_value>` or apply only specific resource `kubectl apply -f <file>.yml -l <label_key>=<label_value>`

### spec
* Configuration for that resource. 
* To get the list of attributes available to set in yaml `kubectl explain <kind> --recursive`. 
* For the detail of a specific attribute `kubectl explain <kind>.<spec>.<attribute>.<...>`


## Commands
### Apply
`kubectl apply -f <file/directory/url>`
* Use yaml to start/control/manage cluster (similar to `docker stack deploy`)
* Use `--dry-run=client/server` to test validity

### Diff
`kubectl diff -f <file>.yml`
* Check the changes before reapplying
