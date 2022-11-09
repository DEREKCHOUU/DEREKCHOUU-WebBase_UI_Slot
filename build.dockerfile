FROM golang:1.19.1-alpine
WORKDIR G:\docker
ADD . G:\docker
RUN cd G:\docker && go build
EXPOSE 4041
ENTRYPOINT G:\docker