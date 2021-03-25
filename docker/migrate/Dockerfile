FROM flyway/flyway:latest
COPY ./deps/db /db
WORKDIR /db

CMD ["clean", "migrate"]

