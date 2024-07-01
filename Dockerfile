FROM golang:1.20-alpine AS build

ARG COMMIT

ENV GOPROXY "https://goproxy.io,direct"

RUN apk update && \
    apk add --no-cache tzdata ca-certificates make gettext bash

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

RUN make build COMMIT=${COMMIT}

FROM alpine

RUN apk update && \
    apk add --no-cache tzdata ca-certificates gettext openssl bash

RUN addgroup -S app && adduser -S -g app app

WORKDIR /app

COPY --from=build /app/config.yaml .
COPY --from=build /app/server .
COPY --from=build /app/i18n/en.json /app/i18n/
COPY --from=build /app/i18n/vi.json /app/i18n/

RUN chmod +x /app/server

RUN chown -R app /app

USER app

EXPOSE 8000

ENTRYPOINT ["./server"]
