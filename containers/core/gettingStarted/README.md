# Getting started with docker


### Build an image and host in a registry
```bash
$ docker image build -t golurk/gsd:first-ctr .
$ docker image ls
$ docker image push golurk/gsd:first-ctr

```

### Runing a containerized app

```bash
# Runing detach
$ docker container run -d --name web -p 8080:8080 golurk/gsd:first-ctr
$ docker container ls
$ docker container stop web
$ docker container ls -a
$ docker container rm web
$ docker container ls -a

# Runing interactive
$ docker container run -it --name test alpine sh

# if you want to leave of interactive container use CTRL + P + Q
```

