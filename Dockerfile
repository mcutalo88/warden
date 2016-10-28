FROM golang:1.7.3

RUN mkdir -p $GOPATH/src/warden
WORKDIR $GOPATH/src/warden

ADD . $GOPATH/src/warden

RUN go get
RUN go build

ENTRYPOINT ["/go/bin/warden", "-t", "Bot MjI5MDM5Mjk0NTAzODQ1ODk4.Cu7R_g.0EwXEq6MrGEyXjSAbQma4GqRa1I"]

EXPOSE 80 443
