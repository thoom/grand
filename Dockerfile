FROM golang:alpine as goalpine

ARG BUILD_VERSION=snapshot
COPY . /go/src/github.com/thoom/goron

WORKDIR /go/src/github.com/thoom/goron
RUN CGO_ENABLED=0 go build -ldflags "-X main.buildVersion=$BUILD_VERSION-docker" -o goron

FROM scratch
LABEL author="Zach Peacock <zdp@thoomtech.com>"

COPY --from=goalpine /go/src/github.com/thoom/goron/goron /bin/

ENTRYPOINT ["/bin/goron"]