package dify

import (
	"crypto/tls"
	"github.com/pkg/errors"
	"net/http"
	"strings"
	"time"
)

type ClientConfig struct {
	Key         string
	Host        string
	ConsoleHost string
	Timeout     int
	SkipTLS     bool
}

type Client struct {
	ApiKey       string
	Host         string
	ConsoleHost  string
	ConsoleToken string
	Timeout      time.Duration
	SkipTLS      bool

	client *http.Client
}

func NewClient(config ClientConfig) (*Client, error) {
	key := strings.TrimSpace(config.Key)
	if key == "" {
		return nil, errors.New("dify ApiKey is required")
	}

	host := strings.TrimSpace(config.Host)
	if host == "" {
		return nil, errors.New("dify Host is required")
	}

	consoleURL := strings.TrimSpace(config.ConsoleHost)
	if consoleURL == "" {
		consoleURL = strings.ReplaceAll(host, "/v1", "/console/api")
	}

	if config.Timeout < 0 {
		return nil, errors.New("Timeout should be a positive number")
	}
	var timeout time.Duration
	if config.Timeout == 0 {
		timeout = DefaultTimeoutSeconds * time.Second
	}

	skipTLS := false
	if config.SkipTLS {
		skipTLS = true
	}

	client := &http.Client{}
	if skipTLS {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	client.Timeout = timeout

	return &Client{
		ApiKey:      key,
		Host:        host,
		ConsoleHost: consoleURL,
		Timeout:     timeout,
		SkipTLS:     skipTLS,
		client:      client,
	}, nil
}
