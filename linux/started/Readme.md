
centos container, using httpd inside

```bash

sudo podman run -d -it -p 80:80 centos:latest /usr/sbin/init
sudo podman exec -it 29555a1693dc011ff39a433d8bdf4bfc2f926fba329d410ef7f71e866fe038f9 /bin/bash

```


DNS CUSTOM
/etc/hosts
ej: 192.168.1.5 mysite.com


net status
sudo iftop -i eth0

dh -ht ext4

journalctl --since "10 minutes ago"

cat /var/log/syslog | grep sshd

kill {PID}

nice -19 yes > /dev/null &
