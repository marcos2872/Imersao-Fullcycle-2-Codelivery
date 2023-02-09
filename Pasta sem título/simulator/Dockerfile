FROM golang:1.20

WORKDIR /go/src
ENV PATH="/go/bn:${PATH}"

RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y

CMD ["tail", "-f", "/dev/null"]