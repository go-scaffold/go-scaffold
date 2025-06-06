# Build image
########################
FROM golang:1.24.4-alpine3.21 AS builder

WORKDIR /var/tmp/app

RUN apk add git

# copy artifacts into the container
ADD ./cmd ./cmd
ADD ./go.mod ./go.mod
ADD ./go.sum ./go.sum
ADD ./pkg ./pkg

# Build the app
RUN go build -o .build/app ./cmd/cli

# Final image
########################
FROM alpine:3.22.0

WORKDIR /opt/app

COPY --from=builder /var/tmp/app/.build/app .

CMD [ "./app" ]
