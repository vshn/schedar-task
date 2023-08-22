FROM docker.io/library/alpine:3.15 as runtime

ENTRYPOINT ["schedar-task"]

RUN \
  apk add --update --no-cache \
    bash \
    ca-certificates \
    curl

RUN \
  mkdir /.cache && chmod -R g=u /.cache

COPY schedar-task /usr/local/bin/

RUN chmod a+x /usr/local/bin/schedar-task

USER 65532:0
