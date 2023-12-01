# Docker sample projects



This repository contains multiple projects:

1. **[Flask App](1-flask-app) (Single-stage build)**
2. **[Golang App](2-golang-app) (Multi-stage build)**
3. **[Wordpress + MySQL](3-wordpress-mysql)**
4. **[Gitea Repository](4-gitea)**
5. **[Grafana](5-grafana)**
6. **[Heimdall Dashboard](6-heimdall)**
7. **[Docker Registry](7-registry) (Distribution)**



### Tip:

> It is best practice to create docker network for your homelab (specially in the production environment)

```bash
docker network create \
  --driver=bridge \
  --subnet=192.168.25.0/24 \
  --ip-range=192.168.25.0/24 \
  --gateway=192.168.25.254 \
  containers
```

