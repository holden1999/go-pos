# We specify the base image we need for our
# go application
FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

# Add this go mod download command to pull in any dependencies
RUN go mod tidy

# we run go build to compile the binary
# executable of our Go program
RUN go build -o binary

#Port to the outside world
#EXPOSE 3030

ENTRYPOINT ["/app/binary"]