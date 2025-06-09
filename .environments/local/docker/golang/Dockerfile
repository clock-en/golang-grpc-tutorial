FROM golang:1.24-bullseye

RUN apt-get update && \
    apt-get install -y git bash wget unzip && \
    wget -q https://github.com/protocolbuffers/protobuf/releases/download/v31.1/protoc-31.1-linux-x86_64.zip -O /tmp/protoc.zip && \
    unzip -o /tmp/protoc.zip -d /usr/local bin/protoc && \
    unzip -o /tmp/protoc.zip -d /usr/local include/* && \
    rm /tmp/protoc.zip && \
    chmod +x /usr/local/bin/protoc && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN go install github.com/air-verse/air@v1.61.7 && \
    go install github.com/golang/protobuf/protoc-gen-go@v1.5.4

WORKDIR /app

# Install dependencies
COPY go.mod go.sum .air.toml ./
RUN go mod download

# Copy source code
COPY . .

CMD ["air"]
