FROM golang:1.22.1-alpine3.19 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /api ./cmd/api/main.go

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /api /api

USER nonroot:nonroot

ENTRYPOINT ["/api"]
