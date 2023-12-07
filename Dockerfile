FROM --platform=amd64 golang:1.21 AS build-go
RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list.d/debian.sources
RUN apt update && apt install -y git unzip tzdata curl
ARG APP_VERSION

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn/,direct
ENV CGO_ENABLED=0
ENV GOSO=linux
ENV GOARCH=amd64
ENV GOMAXPROCS=2
WORKDIR /go/src/agan83/netdoop
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go mod download
COPY ./.git .
COPY ./cmd ./cmd
COPY ./docs ./docs
COPY ./doopfx ./doopfx
COPY ./models ./models
COPY ./server ./server
COPY ./store ./store
COPY ./utils ./utils
COPY ./app.go ./app.go
COPY ./update-version.sh ./update-version.sh
RUN ./update-version.sh ${APP_VERSION}
RUN go build -o /go/bin/netdoop ./app.go

FROM node:20 AS build-node
COPY --from=build-go /go/bin/netdoop /usr/local/bin/netdoop
RUN npm config set registry https://registry.npmmirror.com 
RUN npm install -g pnpm 
ENV UMI_ENV=prod
WORKDIR /source
COPY ./app/package.json /source/package.json
COPY ./app/pnpm-lock.yaml /source/pnpm-lock.yaml
RUN pnpm install
COPY ./app/* /source/
COPY ./app/src /source/src
RUN pnpm build

RUN chmod -R og-w /source/dist
RUN chmod -R a+r /source/dist

FROM alpine:3 AS build-nginx
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && \
    apk add nginx && \
    apk add nginx-mod-stream && \
    rm -rf /var/cache/apk/*

FROM alpine:3
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && \
      apk add git tzdata

# ENV TZ=Asia/Shanghai
# RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /opt/netdoop
RUN touch /opt/netdoop/netdoop.yaml
COPY --from=build-go /go/bin/netdoop /usr/local/bin/netdoop
COPY --from=build-node /source/dist/ /usr/share/netdoop/html/

EXPOSE 9176
ENTRYPOINT [ "netdoop", "run" ]
