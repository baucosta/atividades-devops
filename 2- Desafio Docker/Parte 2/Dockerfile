FROM golang:1.15.2-alpine

WORKDIR /go/src/desafio-docker

COPY . .

RUN go install -v ./...

EXPOSE 8001

CMD ["atividades-devops"]