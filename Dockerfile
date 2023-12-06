FROM golang:1.21.4-alpine AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY *.go .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /labs

FROM --platform=amd64 scratch
COPY --from=build /labs .
COPY views/ /views/
EXPOSE 8080
ENTRYPOINT [ "/labs" ]