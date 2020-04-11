FROM alpine
ARG GOARCH=amd64

RUN apk add -u ca-certificates
ADD ./bin/linux/${GOARCH}/cronohub /app/

WORKDIR /app/
ENTRYPOINT [ "/app/cronohub" ]
