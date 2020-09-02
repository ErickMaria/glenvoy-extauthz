FROM golang:1.14-alpine as builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /gauthz
COPY . .

RUN go mod download

RUN echo "building auth server"
RUN go build -o gauthz-server -ldflags '-libgcc=none -s -w' cmd/server/main.go

RUN echo "building auth migration"
RUN go build -o gauthz-migrate -ldflags '-libgcc=none -s -w' cmd/migration/main.go

FROM alpine:3.12.0

RUN mkdir -p /configs/

COPY --from=builder /gauthz/configs/*.yaml /configs/
COPY --from=builder /gauthz/gauthz-server /bin/
COPY --from=builder /gauthz/gauthz-migrate /bin/

# ENTRYPOINT ["./bin/gauthz-server"]
