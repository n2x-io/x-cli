FROM alpine
RUN apk update && apk add --update --no-cache ca-certificates
COPY n2xctl /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/n2xctl"]
