# go-min-docker

Minimal go app, with a small docker image using Docker multi-stage builds. Deployed on kubernetes, and reads a secret config file.

Deploy on kubernetes using a gitlab CI script like.

## How to use

First, clone the repo and make sure you have go installed.
Then install dep as Golang dependecies manager and run `dep ensure`.

You're then good to go `go run main.go`.

To deploy on kubernetes, configure correctly your kubectl, change both the path for your cluster and docker registry in `deploy.bash` and run it `./deploy.bash` (you might need to `chmod +x deploy.bash` first).

## GRPC

### Compile proto

Inside the proto folder :

```bash
protoc --go_out=plugins=grpc:. *.proto
```

### Test GRPC

Use npm `grpcc`.

```bash
grpcc -p grpc/helloworld.proto -a localhost:50051 -i
```

Then, in the REPL

```javascript
client.sayHello({name: "dixneuf19"}, printReply)
```
