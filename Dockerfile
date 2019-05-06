FROM golang:latest as builder

ARG VCS_REF
ARG VCS_URL
ARG BUILD_DATE
ARG VERSION

LABEL org.label-schema.schema-version="1.0" \
    org.label-schema.version=$VERSION \
    org.label-schema.build-date=$BUILD_DATE \
    org.label-schema.vcs-ref=$VCS_REF \
    org.label-schema.vcs-url=$VCS_URL \
    org.label-schema.vcs-type="Git" \
    org.label-schema.license="MIT" \
    org.label-schema.docker.dockerfile="/Dockerfile"

ENV GO111MODULE=on

WORKDIR /go-modules
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v .

EXPOSE 8000
ENTRYPOINT ["./gin-blog"]