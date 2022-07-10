FROM postgres:alpine

COPY ./ca/server.key /var/lib/postgresql/server.key
COPY ./ca/server.crt /var/lib/postgresql/server.crt

RUN chown postgres:postgres /var/lib/postgresql/server.key /var/lib/postgresql/server.crt
RUN chmod 600 /var/lib/postgresql/server.key /var/lib/postgresql/server.crt