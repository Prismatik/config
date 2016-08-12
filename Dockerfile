FROM ubuntu:14.04
MAINTAINER david@prismatik.com.au

RUN apt-get update && apt-get install --no-install-recommends -y \
    ca-certificates \
    curl \
    mercurial \
    git-core
RUN curl -s https://storage.googleapis.com/golang/go1.6.3.linux-amd64.tar.gz | tar -v -C /usr/local -xz

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PATH /usr/local/go/bin:/go/bin:/usr/local/bin:$PATH

RUN mkdir -p /go/bin
RUN curl https://glide.sh/get | sh

ENV DOCKER=true

ADD . /go/src/github.com/prismatik/config

WORKDIR /go/src/github.com/prismatik/config
RUN glide install

RUN go install github.com/prismatik/config

RUN /go/bin/config

ENTRYPOINT /bin/bash
