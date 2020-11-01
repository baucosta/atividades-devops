FROM golang:1.15.3-alpine as builder

WORKDIR /go/src/desafio-go

COPY . .

RUN go get -d -v

RUN GOOS=linux GOARCH=amd64 go build

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /atividades-devops

EXPOSE 8000

FROM scratch

COPY --from=builder /atividades-devops /atividades-devops

CMD ["/atividades-devops"]