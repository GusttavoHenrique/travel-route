FROM golang:latest AS builder
USER root
ARG FILE_NAME
WORKDIR /app
COPY . .
RUN go mod vendor & go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app .
ARG FILE_NAME
ENV FILE=$FILE_NAME
EXPOSE 8080
ENTRYPOINT ./app -file=/tmp/${FILE}
