# Docker for web developers


### Docker Volumes

* Special type of directory in a container typically referred to as a "data volume" 
* Can be shared and reused among containers
* Data volumes are persisted even after the container is deleted
* Updated to an image won't affect a data volume


#### Creating a data volume
```bash
$ docker run -p 8080:3000 -v /var/www node 
						   |     |
						   v     v
				Create a volume  Container Volume

# Locating a volume
$ docker inspect mycontainer
> "Mounts": [{
>  "Name": "a...",
>  "source": "/mnt/..", # Host location
>  "destination": "/var/www", # Volume location in container
>  ....
> }]

# Customizing the Host Location for a Data Volume
$ docker run -p 8080:3000 -v $(pwd):/var/www node
								|
								v
							Host Location


```

Errors trying to run this container
