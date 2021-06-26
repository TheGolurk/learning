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


### Microservices and the real world

```bash
$ sudo dnf install docker-compose
$ cd multi-container
$ docker-compose up -d
$ docker image ls
$ docker-compose down

```

### troubleshouting
Using fedora 34 you cannot use swarm by an error in --live-restore so you need to update the docker
file

```bash
# /etc/sysconfig/docker
# OLD 
# Modify these options if you want to change the way the docker daemon runs
OPTIONS="--selinux-enabled \
  --log-driver=journald \
  --live-restore \
  --default-ulimit nofile=1024:1024 \
  --init-path /usr/libexec/docker/docker-init \
  --userland-proxy-path /usr/libexec/docker/docker-proxy \
"


# NEW
OPTIONS="--selinux-enabled --log-driver=journald"

if [ -z "${DOCKER_CERT_PATH}" ]; then
    DOCKER_CERT_PATH=/etc/docker
fi
```

