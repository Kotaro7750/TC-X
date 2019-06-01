ARG timezome="Asia/Tokyo"
ARG client="http://localhost:8080"
ARG port="8888"

FROM golang:latest AS build
ADD . /go/src/server/
WORKDIR /go/src/server
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV TZ=${timezome}
ENV CLIENT_URL=${client}
ENV LISTEN_PORT=${port}
RUN go install server

FROM alpine
COPY --from=build /go/bin/server /app/server
COPY ./AozoraMinchoRegular.ttf /app/AozoraMinchoRegular.ttf
WORKDIR /app
ENV TZ=${timezome}
ENV CLIENT_URL=${client}
ENV LISTEN_PORT=${port}
CMD ["./server"]