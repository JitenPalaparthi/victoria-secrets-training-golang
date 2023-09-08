# Integration Testing

- Application must be up and running
- Database must be up and running
- All of them are clean that means application and database instances to be created only for the purpose of integration.
- Up the instances and down the instances are very important.

Step-1: Docker instance of the application

Step-2: Build docker instance and push it to docker hub

```
docker build  . -t jpalaparthi/app:v0.0.1
```

docker push -t jpalaparthi/app:v0.0.1
