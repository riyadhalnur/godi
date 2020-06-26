FROM golang:alpine AS builder

# set go specific env vars
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64

RUN mkdir /build
COPY . /build/
WORKDIR /build

# download dependencies
RUN go mod download

# run tests
RUN go test ./...

# build single linked binary
RUN go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o api /build/cmd/api

# start over using scratch image. no need for anything else anymore
FROM scratch
COPY --from=builder /build/api /api/

WORKDIR /api

CMD ["./api"]
