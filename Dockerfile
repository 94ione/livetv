FROM golang:alpine AS builder
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk update && apk --no-cache add build-base
WORKDIR /go/src/github.com/94ione/livetv/
COPY . . 
RUN GOPROXY="https://goproxy.io" GO111MODULE=on go build -o livetv .

FROM alpine:latest
RUN apk --no-cache add ca-certificates youtube-dl tzdata libc6-compat libgcc libstdc++
WORKDIR /root
COPY --from=builder /go/src/github.com/94ione/livetv/view ./view
COPY --from=builder /go/src/github.com/94ione/livetv/assert ./assert
COPY --from=builder /go/src/github.com/94ione/livetv/.env .
COPY --from=builder /go/src/github.com/94ione/livetv/livetv .
EXPOSE 9000
VOLUME ["/root/data"]
CMD ["./livetv"]
