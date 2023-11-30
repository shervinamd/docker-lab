# Docker sample projects



This repository contains multiple projects:

1. **Flask App (Single-stage build)**
2. **Golang App (Multi-stage build)**
3. **Wordpress + Postgres + NginX**
4. **Gitea Repository**
5. **Grafana**
6. **Heimdall Dashboard**
7. **Docker Registry (Distribution)**



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

