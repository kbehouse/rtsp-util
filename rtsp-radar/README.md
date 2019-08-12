# RTSP Radar

See: https://github.com/Ullaakut/cameradar

## Examples

1. Download [credentials.json](https://github.com/Ullaakut/cameradar/blob/master/dictionaries/credentials.json) to this directory

2. run docker ullaakut/cameradar
```
# 192.168.0.0 ~ 192.168.0.255
docker run -v ${PWD}:/tmp --net=host -t ullaakut/cameradar -t 192.168.0.0/24 --custom-credentials="/tmp/credentials.json"

# 192.168.0.50 ~ 192.168.0.60
docker run -v ${PWD}:/tmp --net=host -t ullaakut/cameradar -t 192.168.0.50-60 --custom-credentials="/tmp/credentials.json"

# Specified IP 192.168.0.54
docker run -v ${PWD}:/tmp --net=host -t ullaakut/cameradar -t 192.168.0.54 --custom-credentials="/tmp/credentials.json"
```



