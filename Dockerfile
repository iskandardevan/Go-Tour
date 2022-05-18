#stageI (Build Binary)
FROM golang:1.17.2-alpine3.14 AS builder
WORKDIR /app
COPY ./ ./ ./
RUN ls
RUN go mod download
RUN go build -o main 

#stageII
FROM alpine:3.14
WORKDIR /root/
COPY --from=builder ./app/config.json .
COPY --from=builder ./app/main .
EXPOSE 8080
CMD [ "./main" ]
