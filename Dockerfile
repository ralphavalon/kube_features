FROM golang:1.11.5-alpine3.8

ADD kube_features /go/bin
ADD swagger-ui/ /go/bin/swagger-ui

EXPOSE 8081

ENTRYPOINT ["/go/bin/kube_features"]
