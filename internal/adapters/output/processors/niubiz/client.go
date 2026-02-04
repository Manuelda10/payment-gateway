package niubiz

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"payment-gateway/internal/ports/output"
	"strings"
	"time"
)

type Client struct {
	http *http.Client
	cfg  Config
	log  output.Logger
}

func NewClient(log output.Logger, cfg Config) *Client {
	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = 10 * time.Second
	}
	return &Client{
		http: &http.Client{Timeout: timeout},
		cfg:  cfg,
		log:  log,
	}
}

func (c *Client) GetSecurityToken(ctx context.Context) (string, error) {
	if strings.TrimSpace(c.cfg.User) == "" || strings.TrimSpace(c.cfg.Password) == "" {
		return "", fmt.Errorf("niubiz credentials missing")
	}

	url := strings.TrimRight(c.cfg.BaseURL, "/") + "/api.security/v1/security"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	raw := c.cfg.User + ":" + c.cfg.Password
	b64 := base64.StdEncoding.EncodeToString([]byte(raw))
	req.Header.Set("Authorization", "Basic "+b64)

	start := time.Now()
	resp, err := c.http.Do(req)
	if err != nil {
		c.log.Error(ctx, "niubiz: request failed", err,
			output.F("url", url),
		)
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	body := strings.TrimSpace(string(bodyBytes))

	c.log.Info(ctx, "niubiz: response received",
		output.F("status", resp.StatusCode),
		output.F("latency_ms", time.Since(start).Milliseconds()),
	)

	// Según tu info: 201 Created y body con token
	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("niubiz unexpected status=%d body=%s", resp.StatusCode, body)
	}

	if body == "" {
		return "", fmt.Errorf("niubiz empty token body")
	}

	return body, nil
}
