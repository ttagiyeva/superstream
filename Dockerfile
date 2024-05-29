FROM golang:1.21 AS build

ENV CGO_ENABLED=0

WORKDIR /workspace

ADD go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o build/server ./cmd/server

FROM gcr.io/distroless/base:nonroot

WORKDIR /

COPY --from=build /workspace/build/server /

EXPOSE 8080  
ENTRYPOINT ["/server"]
