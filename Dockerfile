FROM golang:1.20.4-alpine AS build_base

# Set the Current Working Directory inside the container

WORKDIR /vkcontest

COPY . .

RUN go mod download

WORKDIR ./cmd/SendMessage

# Build the Go app
RUN go build

# Start fresh from a smaller image
FROM alpine:3.17
RUN apk add ca-certificates
RUN apk add --no-cache bash

COPY --from=build_base /vkcontest/cmd/SendMessage/SendMessage /app/app
COPY --from=build_base /vkcontest/config/config.yaml /app/config/config.yaml
COPY --from=build_base /vkcontest/config/TimeDuration.yaml /app/config/TimeDuration.yaml

WORKDIR /app

# Run the binary program produced by `go install`
CMD ["./app"]