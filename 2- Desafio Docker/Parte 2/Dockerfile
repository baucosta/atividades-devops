FROM golang:1.15.2-alpine as builder

WORKDIR /go/src/desafio-docker

COPY . .

# RUN go install -v ./...
RUN go get -d -v

RUN GOOS=linux GOARCH=amd64 go build


EXPOSE 8001

# CMD ["atividades-devops"]

FROM scratch

COPY --from=builder /go/src/desafio-docker/atividades-devops /go/src/desafio-docker/atividades-devops

CMD ["/go/src/desafio-docker/atividades-devops"]