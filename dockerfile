# Start from the latest golang base image
FROM golang:latest AS builder

# Add Maintainer Info
LABEL maintainer="Steven <steven.infinity29@gmail.com>"

# Create an /app directory within our image to hold our application source files
RUN mkdir /app
# Copy everything in the root directory into our /app directory
ADD . /app
# Set the Current Working Directory inside the container
WORKDIR /app

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app
RUN go build -o main .

# ---------------------------------------------------
# the lightweight scratch image we'll
# run our application within
FROM alpine:latest AS production

# copy output from builder stage to production stage
COPY --from=builder /app .
# Run the binary program produced by `go install`
CMD ["./main"]