FROM alpine:latest

RUN apk add --no-cache bash postgresql-client curl
RUN curl -sSf https://atlasgo.sh | sh
COPY deploy_migrations.sh /deploy_migrations.sh
RUN chmod +x /deploy_migrations.sh && \
    sed -i 's/\r$//' /deploy_migrations.sh

CMD ["/deploy_migrations.sh"]