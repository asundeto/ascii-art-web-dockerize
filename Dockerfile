# First part

FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY ./ ./

RUN go build -o main .

# Second part

FROM alpine

WORKDIR /app

COPY --from=builder /go/src/app/ /app/

CMD ["./main"]

EXPOSE 8081

