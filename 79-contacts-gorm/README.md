- Performing database operations

- Up and run MYSQL instance using docker

- Run MYSQL docker instance 

# docker
```
docker run -d --name=mysql2 -p 33406:3306 -e MYSQL_ROOT_PASSWORD=admin123 -e MYSQL_DATABASE=contactsdb -e MYSQL_USER=admin -e MYSQL_PASSWORD=admin123 mysql
```

- To run adminer, that is ueed to run admin UI for rdbms databases

```
docker run -d --name=adminui -p 48088:8080 adminer
```
```
docker run -d --name=mysql2 -p 33406:3306 -e MYSQL_ROOT_PASSWORD=admin123 -e MYSQL_DATABASE=contactsdb -e MYSQL_USER=admin -e MYSQL_PASSWORD=admin123 docker.io/library/mysql
```


# Podman

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
podman run -d --name=mysq2 -p 33406:3306 -e MYSQL_ROOT_PASSWORD=admin123 -e MYSQL_DATABASE=contactsdb -e MYSQL_USER=admin -e MYSQL_PASSWORD=admin123 docker.io/library/mysql
```

```
podman run -d --name=adminui -p 48088:8080 adminer
```


# Table related notes

- When using GORM then tables are automatically created provided by proper tags.While inserting/creating a first record , AutoMigrate method creates the table in the database.
- When using database/sql standard pacakge from golang(not gorm) then you have to create the table in the database prior to inserting a record
- To use postgres 1. Use postgres container. Use environment variables for dabase, username and password. Up and run the container using podman or docker
- use Postgres driver in GORM to open a connection to postgres
