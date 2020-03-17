# URL Shortener API

### API DOC

[Is here](api/README.md)

### Run with docker-compose

```
make run
or 
make stop
```

### Run with Dockerfile
```
docker build -t url_shortener_api_web .
docker run -it -p 5000:5000 --name url_shortener_api_web -t url_shortener_api_web
docker stop url_shortener_api_web
docker rm url_shortener_api_web
```   