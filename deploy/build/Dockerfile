ARG REGISTRY=docker.io
# FROM ${REGISTRY}/golang:1.24-alpine AS builder

# ARG APP_RELATIVE_PATH

# COPY .. /data/app
# WORKDIR /data/app

# RUN rm -rf /data/app/bin/
# RUN go mod tidy && go build -ldflags="-s -w" -o ./bin/server ${APP_RELATIVE_PATH}
# RUN mv config /data/app/bin/


FROM ${REGISTRY}/alpine:3.16

ARG APP_ENV
ENV APP_ENV=${APP_ENV}

WORKDIR /data/app

COPY config /data/app/config
COPY cmd /data/app/cmd

EXPOSE 8291

ENTRYPOINT ./cmd/server/main -conf ./config/prod.yml
