package gofeas

type APIClient interface {
	CommonAPIClient
}

var _ APIClient = &Client{}
