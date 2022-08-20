FROM golang:1.18 as builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN  go build -o /app/backend 

FROM alpine:3.6 as runner 

WORKDIR /app

COPY --from=builder /app/backend ./backend
COPY --from=builder /app/configs/* ./configs/

EXPOSE 8080

CMD [ "/app/backend" ]