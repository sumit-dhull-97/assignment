## Build
FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /article

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /article /article

EXPOSE 3001

ENTRYPOINT ["/article"]