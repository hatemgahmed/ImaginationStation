# Start from the latest golang base image
FROM golang:latest



RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg 

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH   


RUN mkdir -p $GOPATH/src/Impact/server
ADD . $GOPATH/src/Impact/server

# should be able to build now
WORKDIR $GOPATH/src/Impact/server






# Copy go mod and sum files
RUN export GOPATH=/go
RUN export PATH=$PATH:$GOPATH/bin
RUN go get -u github.com/kardianos/govendor
COPY vendor/vendor.json ./vendor
RUN govendor  sync

COPY ./.env ./.env
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
#RUN go mod download

# ARG LOG_DIR=$GOPATH/Impact/server/logs

# # Create Log Directory
# RUN mkdir -p ${LOG_DIR}

# # Environment Variables
# ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log 





# Copy the source from the current directory to the Working Directory inside the container
COPY . .


# Dependencies
#RUN go get github.com/gorilla/mux
#RUN go get github.com/lib/pq
# RUN go get github.com/go-pg/migrations/v7
#RUN go get -u github.com/kardianos/govendor
#RUN export GOPATH=WORKDIR
#RUN export PATH=$PATH:$GOPATH/bin
#RUN echo $GOPATH
#RUN govendor init .
#RUN govendor  sync .

RUN ls
RUN echo $HOST
# Build the Go app
RUN go build -o ./bin 

# Expose port 3000 to the outside world just for documentation
EXPOSE 3000




# VOLUME [${LOG_DIR}]



# Command to run the executable
CMD ["./bin/server"]

# Terminal commands
# sudo docker build -t go-docker-impact  .
# sudo docker run --rm -p 8080:3000 -it go-docker-impact

# sudo docker run --rm -it -v $PWD:/go/src/my-project -w /go/src/my-project trifs/govendor init


# RUN apt-get update && \
#       apt-get -y install sudo

# RUN useradd -m docker && echo "docker:docker" | chpasswd && adduser docker sudo

# USER docker
# CMD /bin/bash