# Golang app (Multi-stage build)

Sampel 'Hello World!' API

Build the app
```bash
docker build -t goapp:v1 .
```

Run the app
```bash
docker run -d -p 8080:8080 goapp:v1

Test
```bash
curl 127.0.0.1:8080
```
