FROM golang:1.19-alpine3.16 AS build

ARG version

RUN apk update && \
    apk add gcc git

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -ldflags="-X 'github.com/yeldiRium/spotify-rules-based-playlists-backend/version.Version=${version}'" -o ../bin/spotify-rules-based-playlists-backend

# ---------------------------------------------------------

FROM alpine:3.16.1

COPY --from=build /go/bin/spotify-rules-based-playlists-backend /home/root/spotify-rules-based-playlists-backend

ENTRYPOINT [ "/home/root/spotify-rules-based-playlists-backend" ]
