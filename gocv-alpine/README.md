# gocv-alpine

Refer from: https://github.com/denismakogon/gocv-alpine


## Test in standalone
```
go run main.go -save=${PWD}/cam/ -rtsp="rtsp://admin:123456@192.168.0.55:554/0/video1"
```


## Docker

Build 
```
docker build -t gocv-alpine .
```

Run 
```
docker run -ti --net=host -v ${PWD}/docker_save:/out gocv-alpine -save=/out -rtsp=rtsp://admin:123456@192.168.0.55:554/0/video1
```

gocv-alpine is **smaller** than Ubuntu image

gocv-alpine(278MB) < gocv-ubuntu(1.2GB)

If you want to more smaller docker image size, look denismakogon repo

https://github.com/denismakogon/gocv-alpine/blob/master/build-stage/Dockerfile