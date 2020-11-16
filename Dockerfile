FROM golang:alpine as builder

ARG SSH_PRIVATE_KEY
RUN apk add --no-cache brotli-dev git mercurial gcc g++ openssh-client \
        && mkdir -p ~/.ssh && umask 0077 \
        && echo "${SSH_PRIVATE_KEY}" > ~/.ssh/id_rsa \
        && ssh-keyscan bitbucket.org >> ~/.ssh/known_hosts \
        && ssh-keyscan github.com >> ~/.ssh/known_hosts \
        && git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/" \
        && git config --global url."git@github.com:".insteadOf "https://github.com/"

ENV GO111MODULE=on
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /app

COPY .gitconfig /root/
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -a -o dummy-server $(pwd)/server/cmd/rest-server

FROM node:alpine AS node_builder
COPY --from=builder /app/client ./
RUN npm install
RUN npm run build

FROM alpine:latest
WORKDIR /root
RUN apk add --no-cache brotli-dev ca-certificates
COPY --from=builder /app/rest-server /usr/local/bin
COPY --from=builder /app/server/.env .

COPY --from=node_builder /dist ./web
EXPOSE 8080

ENTRYPOINT ["rest-server"]