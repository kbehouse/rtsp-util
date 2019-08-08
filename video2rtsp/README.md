# Video to RTSP

## Docker 

refer: https://hub.docker.com/r/jrottenberg/ffmpeg

### Run

Run Docker 
```
docker run --net=host --rm -v $PWD/view_001.mp4:/input_video  kbehouse/video2rtsp
```

Modify **$PWD/view_001.mp4** to your video 


### Build 

Build Docker
```
docker build -t video2rtsp .
```

Run Docker 
```
docker run --net=host --rm -v $PWD/view_001.mp4:/input_video  video2rtsp
```

Modify **$PWD/view_001.mp4** to your video 

Test (Open Other Terminal)
```
go run ../gocv-show/main.go rtsp://localhost:5554/video
```

### Try Different ffserver.conf

```
docker run --net=host --rm -v $PWD/view_001.mp4:/input_video -v $PWD/ffserver.conf:/etc/ffserver.conf  video2rtsp
```


## Standalone Test (No Docker)

Install
```
# ffmpeg 3.4
sudo apt-get install ffmpeg
```

Stream Webcam (/dev/video0)
```
ffserver -d -f ./ffserver.conf

# show webcam
ffmpeg -r 25 -s 352x288 -f video4linux2 -i /dev/video0 http://localhost:8090/feed.ffm

# show input video
ffmpeg -re -i view_001.mp4 http://localhost:8090/feed.ffm
```

Test
```
go run ../gocv-show/main.go rtsp://localhost:5554/video
```

