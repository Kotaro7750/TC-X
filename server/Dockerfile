FROM golang:latest AS build
ADD . /go/src/server/
WORKDIR /go/src/server
ENV GO111MODULE=on
ENV CGO_ENABLED=0
RUN go install server

FROM alpine
COPY --from=build /go/bin/server /app/server
COPY ./AozoraMinchoRegular.ttf /app/AozoraMinchoRegular.ttf
WORKDIR /app
EXPOSE 8888
CMD ["./server"]