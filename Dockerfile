FROM golang:1.13-stretch AS build

WORKDIR /usr/src/app

#COPY go.mod .
#COPY go.sum .
#RUN go mod download

COPY . .
RUN make build


FROM ubuntu:18.04 AS release

MAINTAINER Nozim Yunusov

ENV BINARY url_shortener_api
ENV DBUSER url_shorten
ENV DBPASS url_shorten
ENV DBNAME url_shorten
#
# Установка postgresql
#
ENV PGVER 10
RUN apt -y update && apt install -y postgresql-$PGVER

# Run the rest of the commands as the ``postgres`` user created by the ``postgres-$PGVER`` package when it was ``apt-get installed``
USER postgres

# Create a PostgreSQL role named ``avito`` with ``avito`` as the password and
# then create a database `avito` owned by the ``avito`` role.
RUN /etc/init.d/postgresql start &&\
    psql --command "CREATE USER $DBUSER WITH SUPERUSER PASSWORD '$DBPASS';" &&\
    createdb -O $DBUSER $DBNAME &&\
    /etc/init.d/postgresql stop

# Adjust PostgreSQL configuration so that remote connections to the
# database are possible.
RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/$PGVER/main/pg_hba.conf

# And add ``listen_addresses`` to ``/etc/postgresql/$PGVER/main/postgresql.conf``
RUN echo "listen_addresses='*'" >> /etc/postgresql/$PGVER/main/postgresql.conf
RUN echo "include_dir='conf.d'" >> /etc/postgresql/$PGVER/main/postgresql.conf
#ADD ./postgresql.conf /etc/postgresql/$PGVER/main/conf.d/basic.conf

# Expose the PostgreSQL port
EXPOSE 5432

# Add VOLUMEs to allow backup of config, logs and databases
VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

# Back to the root user
USER root

# Объявлем порт сервера
EXPOSE 5000

COPY ./sql/init.sql ./sql/init.sql
COPY ./configs/config.json ./configs/config.json

# Собранный ранее сервер
COPY --from=build /usr/src/app/$BINARY .

#
# Запускаем PostgreSQL и сервер
#
CMD service postgresql start && ./$BINARY --port=5000

