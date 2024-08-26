FROM golang:1.21-alpine AS build
# Install necessary build tools
RUN apk update && apk upgrade && \
    apk add --no-cache git

WORKDIR /go/app-project
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ./out/appapi .

# lightweight alpine image to run the server 
FROM alpine:3.20.2

# Adds CA Certificates to the image
RUN apk --no-cache add ca-certificates

COPY --from=build /go/app-project/out/appapi /app/appapi

# Switches working directory to /app
WORKDIR "/app"

EXPOSE 80

CMD ["./appapi"]