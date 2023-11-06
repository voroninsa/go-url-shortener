# Start from golang base image
FROM golang:alpine

# LABEL maintainer=""

RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY .env .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

#Setup hot-reload for dev stage
# RUN go get github.com/githubnemo/CompileDaemon
# RUN go get -v golang.org/x/tools/gopls

# ENTRYPOINT CompileDaemon --build="go build -a -installsuffix cgo -o main ." --command=./main

# Build the Go app
RUN go build -o /build ./cmd/go-url-shortener

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD [ "/build", "-s", "SQL" ]
