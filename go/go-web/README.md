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


## Connection Poolin

* Connection Max Lifetime: Sets the maximum amount of time a connection may be used
* Max Idle Connection: Sets the maximum number of connections in the idle connection pool
* Max Open Connections: Sets the maximum number of open connections to the database

You probably need to use Context

## Context
Allows you to set a deadline, cancel a signal, or set other request-scoped values across API
boundaries and between processes
