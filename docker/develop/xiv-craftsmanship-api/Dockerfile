################################################################################
# build stage
################################################################################
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

################################################################################
# final stage
################################################################################
FROM alpine:latest

ENV STAGE=
ENV PORT=
ENV POSTGRE_SQL_HOST=
ENV POSTGRE_SQL_USERNAME=
ENV POSTGRE_SQL_PASSWORD=
ENV POSTGRE_SQL_DB=

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["./main"]

HEALTHCHECK --interval=30s --timeout=10s --retries=3 \
  CMD wget --quiet --spider http://localhost:8080/health || exit 1
