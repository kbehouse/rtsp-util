FROM gocv-build-env

WORKDIR $GOPATH/src/github.com/kbehouse/rtsp-util/gocv-show
COPY . .
RUN go build -o gocv-show; cp gocv-show /gocv-show

ENTRYPOINT ["/gocv-show"]
