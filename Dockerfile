FROM golang:1.22-alpine3.18 AS base

# Add Maintainer info
LABEL maintainer="Rule4"

WORKDIR /fintel

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN apk add --no-cache curl
RUN apk add build-base git

RUN git clone https://github.com/roerohan/wait-for-it
RUN cd wait-for-it && go build -o ./bin/wait-for-it

COPY . .

FROM base AS appbuild
WORKDIR /fintel
RUN GOOS=linux go build -o application main.go

FROM base AS pubsubbuild
WORKDIR /fintel
RUN GOOS=linux go build -o pubsub cmd/pubsub/main.go

FROM alpine:3.18 AS app
COPY --from=appbuild /fintel/application /application
COPY --from=appbuild /fintel/wait-for-it/bin/wait-for-it /usr/local/bin/
ARG dbHost=db
ENV db_host=$dbHost
ARG dbPort=5432
ENV db_port=$dbPort
EXPOSE 8080
RUN ls -aril
CMD wait-for-it -w $db_host:$db_port -t 60 -- application

FROM alpine:3.18 AS pubsub
COPY --from=pubsubbuild /fintel/pubsub /pubsub
RUN ls -aril
CMD ["./pubsub"]
