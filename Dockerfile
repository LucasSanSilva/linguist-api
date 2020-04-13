# Based on this image: https://hub.docker.com/_/golang/
FROM golang:1.13
RUN apt-get -y update && apt-get install -y fswatch
WORKDIR /linguist-api
ENTRYPOINT (go build -o linguist-api-docker && ./linguist-api-docker) & \
(fswatch -m poll_monitor -o /linguist-api/**/*.go /linguist-api/*.go | while read num; \
do \
  echo "Killing server..."; \
  pkill go; \
  pkill linguist-api; \
  (go build -o linguist-api-docker && \
  (./linguist-api-docker)) & \
  echo "Starting server..."; \
done)
EXPOSE 8888
