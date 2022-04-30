FROM alpine:latest

COPY ./bin/github-webhooks-exporter_linux_amd64 /app/github-webhooks-exporter
RUN addgroup -S exporter \
    && adduser -S exporter -G exporter \
    && chown exporter:exporter /app/github-webhooks-exporter

EXPOSE 9212/tcp 9213/tcp
USER exporter

ENTRYPOINT ["app/github-webhooks-exporter"]
