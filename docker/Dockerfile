FROM postgres:latest

ENV POSTGRES_USER
ENV POSTGRES_PASSWORD
ENV POSTGRES_DB

COPY init.sql /docker-entrypoint-initdb.d/

EXPOSE