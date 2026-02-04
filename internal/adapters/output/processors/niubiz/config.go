package niubiz

import "time"

type Config struct {
	BaseURL  string
	User     string
	Password string
	Timeout  time.Duration
}
