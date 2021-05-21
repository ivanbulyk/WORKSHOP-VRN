package api

//Client interacts with 3rd party joke API
type Client interface {
	//GeJoke returns one joke
	GetJoke() (*JokeResponse, error)
}
