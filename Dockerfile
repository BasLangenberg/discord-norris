FROM golang:latest
LABEL maintainer="Bas Langenberg <baslangenberg@gmail.com>"
WORKDIR /app
COPY . .
RUN go build cmd/bot/main.go -o main .
CMD ["./main"]