FROM golang:1.15-alpine
COPY . /app
RUN cd /app && go build .
WORKDIR /app
CMD ["go", "run", "."]