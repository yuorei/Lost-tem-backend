FROM golang:latest

ARG USERNAME=docker
ARG GROUPNAME=docker
ARG UID=1000
ARG GID=1000

RUN apt-get update &&  apt-get install -y git
RUN go install github.com/cosmtrek/air@latest
WORKDIR /app

COPY ./back/go.mod ./back/go.sum ./
RUN go mod download

RUN groupadd -g $GID $GROUPNAME && \
    useradd -m -s /bin/bash -u $UID -g $GID $USERNAME
# USER $USERNAME

CMD ["air", "-c", ".air.toml"]
