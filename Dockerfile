FROM golang:1.21-alpine AS build

ARG BUILD_VERSION=snapshot
COPY . /thoom/grand

WORKDIR /thoom/grand
RUN go get -d ./... 
RUN CGO_ENABLED=0 go build -ldflags "-X main.buildVersion=$BUILD_VERSION" -o grand

FROM scratch
LABEL author="Zach Peacock <zach@thoom.net>"
LABEL org.opencontainers.image.source="https://github.com/thoom/grand"

COPY --from=build /thoom/grand/grand /bin/

ENTRYPOINT ["/bin/grand"]