FROM denismakogon/gocv-alpine:4.0.1-buildstage as build-stage



WORKDIR $GOPATH/src/github.com/kbehouse/rtsp-util/gocv-alpine
COPY . .

RUN go get -u -d gocv.io/x/gocv
RUN go build -o $GOPATH/bin/gocv-alpine



FROM denismakogon/gocv-alpine:4.0.1-runtime

COPY --from=build-stage /go/bin/gocv-alpine /gocv-alpine
ENTRYPOINT ["/gocv-alpine"]
