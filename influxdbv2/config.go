package influxdbv2

type Config struct{}

type Client struct{}

func (c *Config) Client() (interface{}, error) {
	var client Client
	return &client, nil
}

type Ready struct {
	Status  string `json:"status"`
	Started string `json:"started"`
	Up      string `json:"up"`
}
