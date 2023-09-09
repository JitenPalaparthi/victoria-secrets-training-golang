# To Run kafka


# To run using podman-compose 


- Ensure podman is installed on your machine.
- In a terminal, install podman compose with the following command:
- pip3 install https://github.com/containers/podman-compose/archive/devel.tar.gz
- cd into the directory your docker-compose file is located in
- Run podman-compose up -d

- Podman compose github.com link

- https://github.com/containers/podman-compose


```
docker-compose up -d 

or 

podman-compose up -d
```

- To run kafka without podman compose or docker compose

- Kafka
```
podman run -d --name=kafka_b -p 9092:9092 -p 9094:9094  --env-file ./kafka.env docker.io/bitnami/kafka:3.4
```

- Kafka UI

```
podman run -d --name=kafka_ui -p 48080:8080 -e DYNAMIC_CONFIG_ENABLED=true provectuslabs/kafka-ui
```


- To delete all containers

```
podman rm -f $(podman ps -aq)
```