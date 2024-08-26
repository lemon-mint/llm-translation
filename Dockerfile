FROM golang:latest AS build

RUN apt update && \
  apt install curl unzip -y && \
  curl -fsSL https://bun.sh/install | bash

WORKDIR /work


COPY go.mod /work/go.mod
COPY go.sum /work/go.sum
RUN go mod download

ADD . /work
RUN cd web && \
  $HOME/.bun/bin/bun install && \
  $HOME/.bun/bin/bun run build && \
  cd .. && \
  CGO_ENABLED=0 go build -o main.exe -ldflags "-s -w" .

FROM gcr.io/distroless/static-debian12:latest

WORKDIR /
COPY --from=build /work/main.exe /main.exe

ENV PORT=14402
EXPOSE 14402

ENTRYPOINT [ "/main.exe" ]
