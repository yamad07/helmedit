# helm-edit
helm-edit is cli tools for overwrite variable in `values.yaml` file.

## Usage
```
$ go get -u github.com/yamad07/helmedit/cmd/helmedit
$ helmedit --help

NAME:
   helmedit - edit helm

USAGE:
   helmedit [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --values value, -v value  set key-values like image.tag.repository=nginx,tag=latest
   --file value, -f value    set file (default: "values.yaml")
   --help, -h                show help (default: false)
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

