FROM golang:1.13

WORKDIR /go/src/github.com/Ketarin/myapp

ENV GO111MODULE "on"

RUN mkdir -p /etc/myapp/casbin/config

CMD ["/bin/bash", "-c", "make start"]

#uncomment the following line for debug purpose
# ENTRYPOINT ["/bin/bash", "-c", "while true;do sleep 1;done;"]
