# The base go-image
FROM golang:1.14-alpine
 
# Create a directory for the app
RUN mkdir /app
 
# Set working directory
WORKDIR /app
 
# Run command as described:
# 	go get -u github.com/adshao/go-binance/v2; \ --> 
# go build will build an executable file named server in the current directory
RUN apk update;\
	apk add git; \
	go get -u github.com/gorilla/mux; \
	go get github.com/pdepip/go-binance/binance; \
	go get github.com/canthefason/go-watcher; \
	go install github.com/canthefason/go-watcher/cmd/watcher;
 
# Run the server executable
CMD [ "watcher" ]