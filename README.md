# cxcloud_assignment


GO example assignment(Movie Ratings)

	Access OMDb API's and return Rotten Tomatoes rating.  	


Getting Started   

	To start using main package, install Golang.  for docker images, install docker. 
	
Scripts committed.

	assignment/main.go
	assignment/main_test.go

Keyfile..

	assignment/.movieAPIKey

Docker file.

	Dockerfile 

Compile source code    

	$ go build -o movie_rating main.go


To run the Golang pkg.     

	$ ./movie_rating
	Movie Rating Source: Rotten Tomatoes
	Movie Rating: 82%


Building docker image(movie-rating-cli)..	

	$ docker build -t movie-ratting-cli .


Run the docker image..

	$ docker run movie-ratting-cli Frozen
	Movie Rating Source: Rotten Tomatoes
	Movie Rating: 90%



Test

	$ go test -v

Code Coverage

  	$ go test -v -cover 
  
  
