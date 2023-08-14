FROM ubuntu:22.04
RUN apt update && apt install tzdata netcat curl wget net-tools ca-certificates -y

WORKDIR /app
COPY ./release/api-server /app/api-server
CMD ["sh", "-c", "/app/api-server -c /app/api-config.yaml"]
