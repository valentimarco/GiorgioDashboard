# Compile css from tailwind
FROM oven/bun:1-debian as tailwind
WORKDIR /usr/src/app

COPY tailwind.config.js package.json bun.lockb ./
COPY templates ./templates
RUN mkdir css
COPY css/input.css ./css

RUN bun install --frozen-lockfile
RUN bun run build


# Base
FROM golang:1.22.2-bullseye AS base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY templates static src public /app/

COPY --from=tailwind /usr/src/app/css/output.css ./css

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/a-h/templ/cmd/templ@v0.2.663

RUN go mod download

ARG UID
ARG GID

RUN groupadd -g $GID appgroup && \
    useradd -u $UID -g $GID -m appuser
USER appuser

CMD ["air", "-c", ".air.toml"]