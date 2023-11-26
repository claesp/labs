docker build --rm -t labs:alpha .
docker run -d -p 8080:8081 --name labs labs:alpha
