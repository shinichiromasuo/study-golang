FROM debian:stretch

RUN apt-get update \
 && apt-get install -y supervisor tzdata curl ca-certificates \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/* \
 && : "set timezone" \
 && echo "Asia/Tokyo" > /etc/timezone \
 && rm /etc/localtime \
 && ln -s /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
 && dpkg-reconfigure -f noninteractive tzdata \
 && : "" \
 && touch /etc/revision \
 && : "" \
 && mkdir -p /var/log/supervisor

COPY supervisord.conf /etc/supervisor/supervisord.conf

COPY golang-api-fcgi /usr/local/bin/golang-api-fcgi

CMD ["/usr/bin/supervisord", "-n", "-c", "/etc/supervisor/supervisord.conf"]
