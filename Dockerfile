FROM golang:latest AS builder
USER root
ARG FILE_NAME
COPY . /app
WORKDIR /app
RUN go mod vendor & go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app .
ARG FILE_NAME
ENV FILE=$FILE_NAME
ENTRYPOINT ./app -file=/app/data/${FILE}
EXPOSE 8080