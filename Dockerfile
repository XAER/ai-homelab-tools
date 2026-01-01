# Build
FROM golang:1.22-alpine AS build
WORKDIR /src
RUN apk add --no-cache git ca-certificates
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /out/ai-homelab-tools ./cmd/server

# Run
FROM gcr.io/distroless/static:nonroot
ENV ADDR=:7070
USER nonroot:nonroot
COPY --from=build /out/ai-homelab-tools /ai-homelab-tools
EXPOSE 7070
ENTRYPOINT ["/ai-homelab-tools"]

