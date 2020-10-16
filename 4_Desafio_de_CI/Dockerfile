FROM golang:1.15.2-alpine as builder

WORKDIR /go/src/desafio-ci

COPY ./src/ .

RUN go get -d -v

RUN GOOS=linux GOARCH=amd64 go build


EXPOSE 8001

FROM scratch

COPY --from=builder /go/src/desafio-ci/atividades-devops /go/src/desafio-ci/atividades-devops

CMD ["/go/src/desafio-ci/atividades-devops"]