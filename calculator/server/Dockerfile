FROM golang:1.16

# Set the current working Directory inside the container
WORKDIR /go/src/server

# Copy the source from the current directory to the WORKDIR inisde the container
COPY . .

# call Makefile to download dependencies and build the executable
RUN make all
# Expose port 50051 to the outisde world
EXPOSE 50051

#start the "server" exeutable
CMD ["./server"]