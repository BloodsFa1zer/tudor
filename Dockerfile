FROM --platform=linux/amd64 golang:1.21.5-alpine3.19 as builder
WORKDIR /workspaces/app

RUN apk update && \
  apk add --no-cache make

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make --no-print-directory build

FROM --platform=linux/amd64 alpine:3.19 as release
WORKDIR /app
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY --from=builder /workspaces/app/bin/run /app

CMD [ "./run" ]