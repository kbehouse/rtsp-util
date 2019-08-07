# GOCV Show RTSP

## RUN

```
go run main.go  [RTSP_URL]

EX:
go run main.go  rtsp://admin:123456@192.168.0.55:554/0/video1
```


## Dockerfile

### Build Docker

refer: https://github.com/hybridgroup/gocv

Generate **gocv-build-env**

```
docker build --build-arg OPENCV_VERSION=4.1.0 --build-arg GOVERSION=1.12.6 -t gocv-build-env -f Dockerfile-build-env .
```

Generate **gocv-show** docker 

```
docker build -t gocv-show .
```

Run Docker 

```
docker run -ti --rm -e DISPLAY=$DISPLAY -v /tmp/.X11-unix:/tmp/.X11-unix -v $HOME/.Xauthority:/root/.Xauthority  --net=host  gocv-show rtsp://admin:123456@192.168.0.55:554/0/video1
```