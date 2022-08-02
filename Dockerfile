# Builder
FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY main.go ./
COPY api ./api

RUN go build


# Runner
FROM alpine:latest

ENV PORT 3000

WORKDIR /app

COPY --from=build /app/nobl9-recruitment-task ./nobl9-recruitment-task

EXPOSE $PORT

CMD ["sh", "-c", "./nobl9-recruitment-task --port $PORT"]
