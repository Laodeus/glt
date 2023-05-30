FROM golang:1.16 as builder

WORKDIR /app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -v -o ./dist/glt ./main/main.go

FROM alpine:3.13 
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/dist/glt ./dist/glt

COPY --from=builder /app/.env ./.env

# Add docker-compose-wait tool
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

ENTRYPOINT /root/dist/glt