## task 2

### use supervisor

use `supervisor` 


### push to docker hub



1. login

   ```
   $ sudo docker login
   Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
   Username: catyuanbao
   Password: 
   WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
   Configure a credential helper to remove this warning. See
   https://docs.docker.com/engine/reference/commandline/login/#credentials-store
   
   Login Succeeded
   ```

   

2. tag image

   ```
   sudo docker tag bd1b9fc27f4a catyuanbao/httpserver:v.1.3
   ```

   

3. push image

   ```
   $ sudo docker push catyuanbao/httpserver:v.1.3
   The push refers to repository [docker.io/catyuanbao/httpserver]
   fd728e7847f8: Layer already exists 
   f1e124efc129: Layer already exists 
   cca347fb5dfe: Layer already exists 
   1ace9b92f976: Layer already exists 
   bc4d014ce082: Layer already exists 
   ac6a943b6965: Layer already exists 
   0f5a06a84d99: Layer already exists 
   6c318ab56d7e: Layer already exists 
   606d67d8e1b8: Layer already exists 
   v.1.3: digest: sha256:7106005ed05e867334014268c600d8cb38c3d4cf8d3fb086c960bbbdc3e35a00 size: 2204
   
   ```



### show container IP

```
$ sudo docker exec -it 63c1ccb10abe bash            
[root@63c1ccb10abe /]# ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
103: eth0@if104: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:05 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.5/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever

```

```
$ sudo docker inspect -f {{.State.Pid}} 63c1ccb10abe
18634

# back @ backbox in ~/geektime_tasks/task2 on git:main x [21:21:05] 
$ sudo nsenter -n -t 18634                          
backbox# ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
103: eth0@if104: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:05 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.5/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```


