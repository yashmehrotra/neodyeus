FROM golang:latest

WORKDIR $GOPATH/src/github.com/yashmehrotra/neodyeus
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep && dep ensure -v
RUN go build -o Neodyeus main.go
ENV NEODYEUS_PATH /tmp
ENTRYPOINT ["./Neodyeus"]
