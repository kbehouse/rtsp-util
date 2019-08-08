#  Generate Batch Dockers of video2rtsp

 Generate batch bockers of video2rtsp to IP  192.168.0.140 ~ 192.168.0.140 [+video_num]

 The [video_num] depend on the videos in **videos folder** 

## Create Docker Network Macvlan 

Because we want to brocast video rtsp to 192.168.0.140 ~ 192.168.0.[140+video_num]

We need to use macvlan

Open modprobe macvlan
```
sudo modprobe macvlan
```

Generate new docker network **macvlan-192-168-0**
```
docker network create -d macvlan \
  --subnet=192.168.0.0/24 \
  --gateway=192.168.0.1 \
  -o parent=enp3s0 \
  macvlan-192-168-0
```

modify **enp3s0** to your network interface (# ifconfig)


Check docker network whether the **macvlan-192-168-0** exist?
```
docker network ls
```

## Generate Dockers 


Generate dockers 
```
python batch_rtsp_docker.py
```


**NOTE: Maybe you should test in different computer**

Test 
```
go run ../gocv-show/main.go rtsp://192.168.0.147:5554/video
```


## Other Tool

Check network IP usage
```
nmap -sn 192.168.0.0/24
```

Test Only 1 IP
```
docker run -d --network macvlan-192-168-0  --ip=192.168.0.146 -v $PWD/videos/view_001.mp4:/input_video --name 192.168.0.144_view_001.mp4 kbehouse/video2rtsp
```