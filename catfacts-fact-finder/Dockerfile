FROM golang:1.15-alpine
COPY . /app
RUN cd /app && go build -o catfact-finder
WORKDIR /app

ENV LOCAL_PORT=8080
ENV LOCAL_PROTOCOL="http"
ENV FINDFACT_OFFLINE=false

EXPOSE $LOCAL_PORT

CMD ["./catfact-finder"]
