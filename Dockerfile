# BUILD STAGE
FROM golang:1.24 AS build-stage
WORKDIR /app
# Copy Files
COPY go.mod go.sum ./
COPY rest ./rest
COPY scraper ./scraper
COPY *.go ./
# Install Dependencies
RUN go mod download
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /naverdictionary

# TEST STAGE
FROM build-stage AS test-stage
RUN go test -v ./...

# DEPLOY STAGE
FROM gcr.io/distroless/base-debian11 AS deploy-stage
WORKDIR /
COPY --from=build-stage /naverdictionary /naverdictionary
EXPOSE 8080
USER nonroot:nonroot
CMD ["/naverdictionary"]