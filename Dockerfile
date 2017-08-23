FROM golang:1.5.2

COPY . /go/src/github.com/edlvj/gn-api
WORKDIR /go/src/github.com/edlvj/gn-api

# Install Revel CLI
RUN go get github.com/revel/cmd/revel
RUN go get gopkg.in/mgo.v2

# Install project dependencies
RUN go get github.com/tools/godep
RUN godep go install ./app

# Run app in production mode
EXPOSE 9000
ENTRYPOINT revel run github.com/edlvj/gn-api prod 9000