FROM golang:1.16

# Set the current working Directory inside the container
WORKDIR /go/src/server

# Copy go mod and sum files
COPY go.mod go.sum /go/src/server/

# Download all dependencies. they will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the WORKDIR inisde the container
COPY . .

# Build the Go app, since it is run from WORKDIR, the executable will be created in WORKDIR
RUN go build calServer/server.go

# Expose port 50051 to the outisde world
EXPOSE 50051

#start the "server" exeutable
CMD ["./server"]






# FROM golang:alpine
# COPY calculator /go/src/calculator
# WORKDIR /go/src/calculator
# RUN go run calClient/*.go