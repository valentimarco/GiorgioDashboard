# Compile css from tailwind
FROM oven/bun:1-debian as tailwind
WORKDIR /usr/src/app

COPY tailwind.config.ts package.json bun.lockb ./
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

COPY templates ./templates
COPY src ./src
COPY public ./public
COPY static ./static

COPY --from=tailwind /usr/src/app/css/output.css ./css/output.css
COPY --from=tailwind /usr/src/app/node_modules/htmx.org/dist ./static/htmxlib
COPY --from=tailwind /usr/src/app/node_modules/jquery/dist ./static/jquery

RUN go install github.com/a-h/templ/cmd/templ@latest

RUN go mod download

RUN templ generate

ARG UID
ARG GID

RUN CGO_ENABLED=0 GOOS=linux go build -o ./main ./src/main.go

RUN groupadd -g $GID appgroup && \
    useradd -u $UID -g $GID -m appuser
USER appuser

CMD ["./main"]