# flaskapp

### A sample Flask API

Build the app
```bash
docker build -t flaskapp:v1 .
```
Run the app
```bash
docker run -d -p 5000:5000 flaskapp:v1
```
Test
```bash
curl http://127.0.0.1:5000/hello
```
