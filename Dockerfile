# Builder image
FROM golang:1.22-alpine3.20 AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

ARG ACCESS_TOKEN

RUN apk add git

WORKDIR /{{SERVICE_NAME}}

COPY . .

RUN git config --global url."https://${ACCESS_TOKEN}@github.com/".insteadOf "https://github.com/"

RUN go mod download

RUN go build -o {{SERVICE_NAME}}

# ----------------------------------------------------------------------------
# Production image
FROM alpine:3.19

RUN apk --update --no-cache add ca-certificates tzdata

COPY --from=builder /{{SERVICE_NAME}}/{{SERVICE_NAME}} /go/bin/{{SERVICE_NAME}}

ENTRYPOINT ["/go/bin/{{SERVICE_NAME}}"]
