FROM golang:1.14-alpine as builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /opt/gauthz

COPY . .

# do this in a separate layer to cache deps from build to build
# RUN go mod download
RUN go build -o gauthz-server -ldflags '-libgcc=none -s -w' cmd/server/main.go 

# FROM gcr.io/distroless/base
FROM alpine:3.12.0
RUN mkdir -p /opt/gauthz /opt/gauthz/configs
COPY --from=builder /opt/gauthz/configs/application.yaml /opt/gauthz/configs
COPY --from=builder /opt/gauthz/gauthz-server /bin
ENTRYPOINT ["./bin/gauthz-server"]