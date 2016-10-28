FROM golang:1.7.3-onbuild
RUN go build /go/src/app
CMD ['/go/bin/warden', '-t', 'MjI5MDM5Mjk0NTAzODQ1ODk4.Cu7R_g.0EwXEq6MrGEyXjSAbQma4GqRa1I']
EXPOSE 80 443
