FROM golang:1.20.1-alpine

WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code
COPY *.go ./

# Build
RUN go build -o /taraxa-indexer

EXPOSE 8080
ENV HTTP_PORT=8080

# Run
CMD [ "/taraxa-indexer" ]