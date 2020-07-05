FROM golang as build-env
WORKDIR /app
#ADD ./vendor/* /go/src/github.com/
ADD . /app
ENV GOOS=linux 
ENV GOARCH=amd64 
ENV CGO_ENABLED=0
#RUN go get github.com/gorilla/mux
#RUN go get github.com/kwahome/go-haversine/pkg/haversine
RUN cd /app && go mod download
RUN cd /app && go build -o datapoa

# FROM golang:alpine
FROM scratch
COPY --from=build-env /app/datapoa /app/datapoa
WORKDIR /app
ENTRYPOINT [ "./datapoa"]