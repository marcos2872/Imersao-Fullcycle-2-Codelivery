FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bn:${PATH}"

RUN apt-get updates && \
    apt-get install build-essential librdkafka-dev -y

CMD ["tail", "-f", "/dev/null"]