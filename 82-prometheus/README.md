# Monitoring

- How to run prometheus server

```
docker pull prom/prometheus
docker run --name prometheus -d -p 127.0.0.1:9090:9090 prom/prometheus
```

- To access dashboard

```
google-chrome http://localhost:9090
firefox http://localhost:9090
```

- To run without any issues to run prometheus on host network/connect to host network applicaiton.. run prometheus with host network

```
docker run -d -v /home/jiten/workspace/personal/training/victoria-secrets-golang/82-prometheus/prometheus.yml:/etc/prometheus/prometheus.yml --network=host prom/prometheus
```

-- Presentation link

https://docs.google.com/presentation/d/1zVaXK_s4owQGyK1Pj42aRTPwrynbJRkbOhrw-9JnvYo/edit?usp=sharing
