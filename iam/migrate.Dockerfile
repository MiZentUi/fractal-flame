FROM alpine:3.22

WORKDIR /migrate

RUN wget https://github.com/pressly/goose/releases/download/v3.27.0/goose_linux_x86_64 && \
    mv goose_linux_x86_64 goose && \
    useradd -S migrategroup && \
    useradd -S migrateuser -G migrategroup && \
    chown migrateuser:migrategroup goose

USER migrateuser

COPY /migrations/*.sql ./migrations/

ENTRYPOINT [ "./goose", "up" ]
