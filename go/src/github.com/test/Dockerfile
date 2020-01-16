# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev
WORKDIR /app
# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV SRC_DIR=/SRC
# Add the source code:
COPY ./assignment/* $SRC_DIR/
# Build it:
RUN cd $SRC_DIR; go build -o movie_rating; cp movie_rating /app/
#RUN cd $SRC_DIR; cp movie_rating /app/
ENTRYPOINT ["./movie_rating"]