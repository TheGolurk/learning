
centos container, using httpd inside

```bash

sudo podman run -d -it -p 80:80 centos:latest /usr/sbin/init
sudo podman exec -it 29555a1693dc011ff39a433d8bdf4bfc2f926fba329d410ef7f71e866fe038f9 /bin/bash

```
