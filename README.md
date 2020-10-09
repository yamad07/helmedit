# helm-edit
helm-edit is cli tools for overwrite variable in `values.yaml` file.

## Usage
```
$ go get -u github.com/yamad07/helmedit/cmd/helmedit
```

## Example
```fixtures/values.yaml
replicaCount: "10"
service:
  annotations:
    meta: civilization.solution

```

if you want to overwrite `replicaCount`,
```
$ helmedit -f fixtures/values.yaml -v replicaCount=5
```

and then

```fixtures/values.yaml
replicaCount: "5"
service:
  annotations:
    meta: civilization.solution

```

