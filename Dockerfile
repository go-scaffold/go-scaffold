# Build image
########################
FROM golang:1.13.0-alpine3.10 as builder

WORKDIR /var/tmp/app

RUN apk add git

# copy artifacts into the container
ADD ./main.go ./main.go
ADD ./go.mod ./go.mod
ADD ./go.sum ./go.sum
ADD ./pkg ./pkg

# Build the app
RUN go build -o .build/app .

# Final image
########################
FROM alpine:3.10

WORKDIR /opt/app

COPY --from=builder /var/tmp/app/.build/app .

CMD [ "./app" ]
