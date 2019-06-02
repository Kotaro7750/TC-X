FROM golang:latest AS build
ADD . /go/src/server/
WORKDIR /go/src/server
ENV GO111MODULE=on
ENV CGO_ENABLED=0
RUN go install server

FROM alpine
RUN apk --update add tzdata && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
COPY --from=build /go/bin/server /app/server
COPY ./AozoraMinchoRegular.ttf /app/AozoraMinchoRegular.ttf
ENV TZ="Asia/Tokyo"
WORKDIR /app
EXPOSE 8888
CMD ["./server"]