# Containers

## Linux in a container world

Host administration:
* Privileges
* Kernel access

Container adminitration:
* Application design

The future of Linux
* Microservices
* DevOps
* Development support


## Docker resources
Image hosting:
* Docker hub
* Docker registry

Storage
* Docker volumes
* Third-party-solutions


```bash
$ docker build -t "webserver" -f sample.Dockerfile .
$ docker images
$ docker run -d -p 80:80 webserver /usr/sbin/apache2ctl -D FOREGROUND
$ docker ps
$ curl localhost
$ docker login
$ docker search nextcloud
$ docker tag webserver golurk/webserver
$ docker push golurk/webserver
```
