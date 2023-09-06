- Performing database operations

- Up and run MYSQL instance using docker

- Run MYSQL docker instance 

# docker
```
docker run -d --name=mysql1 -p 33306:3306 -e MYSQL_ROOT_PASSWORD=admin123 -e MYSQL_DATABASE=contactsdb -e MYSQL_USER=admin -e MYSQL_PASSWORD=admin123 mysql
```

- To run adminer, that is ueed to run admin UI for rdbms databases

```
docker run -d --name=adminui -p 48088:8080 adminer
```

docker run -d --name=mysql1 -p 33306:3306 -e MYSQL_ROOT_PASSWORD=admin123 -e MYSQL_DATABASE=contactsdb -e MYSQL_USER=admin -e MYSQL_PASSWORD=admin123 docker.io/library/mysql


# Podman =
- Steps to follow to use podman. Give docker hub user and password

```
podman login docker.io
```

- Pull mysql and adminer containers
```
podman pull --tls-verify=false docker.io/library/mysql
podman pull --tls-verify=false docker.io/library/adminer
```

```
podman run -d --name=mysql1 -p 33306:3306 -e MYSQL_ROOT_PASSWORD=admin123 -e MYSQL_DATABASE=contactsdb -e MYSQL_USER=admin -e MYSQL_PASSWORD=admin123 docker.io/library/mysql
```

```
podman run -d --name=adminui -p 48088:8080 adminer
```
