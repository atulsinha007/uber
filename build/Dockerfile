#STAGE1
#CREATING THE BINARY FROM CODE
FROM golang:1.15.3-alpine3.12 as builder

ENV PROJECT_PATH=${GOPATH}/src/github.com/atulsinha007/uber

RUN mkdir -p ${PROJECT_PATH}
WORKDIR ${PROJECT_PATH}
COPY . .

RUN wget https://github.com/golang-migrate/migrate/releases/download/v4.7.1/migrate.linux-amd64.tar.gz
RUN tar xvzf migrate.linux-amd64.tar.gz

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o ./bin/uber ./cmd


# STAGE 2
# COPYING THE BINARY FROM STAGE 1 AND RUNNING IT
FROM ubuntu

RUN apt update && apt install -y ca-certificates curl && rm -rf /var/cache/apt/*

ENV GOPATH=/go
ENV PROJECT_PATH=${GOPATH}/src/github.com/atulsinha007/uber

COPY --from=builder ${PROJECT_PATH}/bin/ ./bin/
COPY --from=builder ${PROJECT_PATH}/config ${PROJECT_PATH}/config/
COPY --from=builder ${PROJECT_PATH}/migrations ${PROJECT_PATH}/migrations
COPY --from=builder ${PROJECT_PATH}/migrate.linux-amd64 .

CMD ["/bin/bash", "-c", "./bin/uber"]
