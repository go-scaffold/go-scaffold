# Build image
########################
FROM golang:1.12.5-alpine3.9 as builder

WORKDIR /var/tmp/app

# copy artifacts into the container
ADD . .

# Build the app
RUN go build -o .build/app .

# Final image
########################
FROM alpine:3.9

WORKDIR /opt/app

COPY --from=builder /var/tmp/app/.build/app .

CMD [ "./app" ]
