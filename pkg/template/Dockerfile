FROM golang AS builder
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn
WORKDIR /go/src/app
COPY . .
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o main ./cmd/app

FROM alpine:latest
COPY --from=builder /go/src/app/main /main
RUN chmod +x /main

ENV TZ Asia/Shanghai
EXPOSE 8080
CMD ["/main"]
