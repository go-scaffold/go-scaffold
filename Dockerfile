# Build image
########################
FROM golang:1.20.7-alpine3.18 as builder

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
FROM alpine:3.18.4

WORKDIR /opt/app

COPY --from=builder /var/tmp/app/.build/app .

CMD [ "./app" ]
