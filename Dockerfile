FROM golang as build-env
WORKDIR /app

ADD . /app
ENV GOOS=linux 
ENV GOARCH=amd64 
ENV CGO_ENABLED=0

RUN cd /app && go mod download
RUN cd /app && go build -o datapoa

FROM scratch
COPY --from=build-env /app/datapoa /app/datapoa
ADD json /app/json
ADD dados /app/dados
WORKDIR /app
ENTRYPOINT [ "./datapoa"]
