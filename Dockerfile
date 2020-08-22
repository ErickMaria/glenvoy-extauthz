FROM golang:1.14-alpine as builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /opt/gauthz
COPY . .

RUN go mod download
RUN go build -o gauthz-server -ldflags '-libgcc=none -s -w' cmd/server/main.go 

FROM alpine:3.12.0
RUN mkdir -p /opt/gauthz /opt/gauthz/configs
COPY --from=builder /opt/gauthz/configs/application.yaml /opt/gauthz/configs
COPY --from=builder /opt/gauthz/gauthz-server /bin
ENTRYPOINT ["./bin/gauthz-server"]