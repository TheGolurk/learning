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


```bash
$ docker service create --name web -p 8080:8080 \\n--replicas 3 golurk/gsd:first-ctr
$ docker service ls
$ docker container ls
$ top
$ docker service ps web
$ docker service scale web=10
$ docker container ls
$ docker container rm 326115391ed6 8c0b3a6c7f07 20a956f13b0a -f
$ docker container ls
$ docker service rm web
$ docker service ls
$ docker container ls
```


```bash
$ docker image build -t golurk/gsd:swarm-stack .
$ docker image push golurk/gsd:swarm-stack
$ docker stack deploy -c docker-compose.yml counter
$ docker stack ls
$ docker stack services counter
$ docker stack ps counter
$ vim docker-compose.yml
$ docker stack deploy -c docker-compose.yml counter
$ docker stack ps counter
$ docker stack rm counter
 ```
