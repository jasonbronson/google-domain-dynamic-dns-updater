FROM alpine:latest

RUN apk update && \
    apk add --no-cache curl bind-tools dcron && \
    rm -rf /var/cache/apk/*

COPY dnsupdate.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/dnsupdate.sh

# Add the cron job
RUN echo "*/1 * * * * /usr/local/bin/dnsupdate.sh" >> /etc/crontabs/root && \
    chmod 0644 /etc/crontabs/root

CMD ["/usr/sbin/crond", "-f", "-l", "2", "-L", "/dev/stdout"]

