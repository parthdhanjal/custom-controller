# Build the manager binary
FROM golang:1.18 as builder

WORKDIR /workspace

# Copy the go source
COPY . /workspace

# Build
RUN go build ./...

ENTRYPOINT ["/workspace/custom-controller"]