ARG ARTIFACT_REPOSITORY_DOMAIN

FROM ${ARTIFACT_REPOSITORY_DOMAIN}/golang:1.27-alpine AS builder
RUN apk add --no-cache git gcc musl-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -tags "containers_image_openpgp exclude_graphdriver_btrfs" -o /app/build-doctor .

FROM ${ARTIFACT_REPOSITORY_DOMAIN}/quay.io/podman/stable:v5.8.9

COPY --from=builder /app/build-doctor /usr/local/bin/build-doctor

RUN echo "podman:100000:65536" > /etc/subuid && \
    echo "podman:100000:65536" > /etc/subgid

ENV STORAGE_DRIVER=overlay

WORKDIR /workspace

ENTRYPOINT ["/bin/sh", "-c", "/usr/local/bin/build-doctor < /dev/null"]
