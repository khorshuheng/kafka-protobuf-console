FROM golang:1.12-stretch
WORKDIR /usr/src
ADD $PWD/go.mod /usr/src/
RUN go mod download
ADD $PWD /usr/src/kafka-protobuf-console
WORKDIR /usr/src/kafka-protobuf-console
RUN make install
CMD ["tail", "-f", "/dev/null"]
