FROM alpine:3.14
RUN apk add go

EXPOSE 8000
COPY ["cmd", "/app/"]

ENTRYPOINT go run /app/main.go