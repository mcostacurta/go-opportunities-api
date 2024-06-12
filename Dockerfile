FROM golang:latest

# Set the current working directory inside the container
WORKDIR /app

ENV GOFLAGS="-buildvcs=false"

RUN go install github.com/air-verse/air@latest

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the workspace
COPY . .

# Build the Go app
RUN go build -o gooportunities .

# Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
# CMD ["air"]
CMD ["air", "-c", ".air.toml"]