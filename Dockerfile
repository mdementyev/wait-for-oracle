ARG GO_VERSION=1.18

FROM golang:${GO_VERSION} as build

ENV GO111MODULE=on

WORKDIR /app/

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s -extldflags '-static'" -a .

FROM scratch as publish

COPY --from=build /app/wait-for-oracle /app/wait-for-oracle

ENTRYPOINT [ "/app/wait-for-oracle" ]