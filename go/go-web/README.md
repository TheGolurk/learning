# Creating go web services

## database
I will run a mysql container

```
$ podman pull mysql/mysql-server // docker image
$ sudo podman run --name go-web-mysql -p 3306:3306 -v /home/golurk/data/mysql:/var/lib/mysql:Z --user 1000:1000 -e MYSQL_ROOT_PASSWORD=golurk -d mysql:latest
$ sudo podman exec -it go-web-mysql bash
$ mysql -u root -p
$ golurk
```
