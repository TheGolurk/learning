# centos 7 notes


##  troubleshooting

max_user_namespaces err
```bash
echo “user.max_user_namespaces=10000” > /etc/sysctl.d/42-rootless.conf
sudo echo “user.max_user_namespaces=10000” > /etc/sysctl.d/42-rootless.conf
sudo su
sudo shutdown -r
sudo shutdown -r now
ls
podman
vim /etc/sysctl.d/42-rootless.conf
cat /proc/sys/user/max_user_namespaces
sudo sysctl user.max_user_namespaces=15000
cat /proc/sys/user/max_user_namespaces
sudo usermod --add-subuids 200000-201000 --add-subgids 200000-201000 golurk
grep golurk /etc/subuid /etc/subgid
```
```


