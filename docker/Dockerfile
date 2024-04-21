FROM golang:1.22.2-bullseye

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/a-h/templ/cmd/templ@v0.2.663

COPY go.mod go.sum ./
RUN go mod download

ARG UID
ARG GID

RUN groupadd -g $GID appgroup && \
    useradd -u $UID -g $GID -m appuser
USER appuser

CMD ["air", "-c", ".air.toml"]