# syntax=docker/dockerfile:1
FROM golang:1.21-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0


COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app .


FROM gcr.io/distroless/static
COPY --from=build /app /app
EXPOSE 8080
USER 65532:65532
ENTRYPOINT ["/app"]
