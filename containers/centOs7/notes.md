# centos 7 notes


##  troubleshooting

```
max_user_namespaces err
```
  69  echo “user.max_user_namespaces=10000” > /etc/sysctl.d/42-rootless.conf
   70  sudo echo “user.max_user_namespaces=10000” > /etc/sysctl.d/42-rootless.conf
   71  sudo su
   72  sudo shutdown -r
   73  sudo shutdown -r now
   74  ls
   75  podman
   76  vim /etc/sysctl.d/42-rootless.conf
   77  cat /proc/sys/user/max_user_namespaces
   78  sudo sysctl user.max_user_namespaces=15000
   79  cat /proc/sys/user/max_user_namespaces
   80  sudo usermod --add-subuids 200000-201000 --add-subgids 200000-201000 golurk
   81  grep golurk /etc/subuid /etc/subgid
```
```


