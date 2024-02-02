// client.go
package proxmox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client represents the Proxmox API client
type Client struct {
	BaseURL   string
	Ticket    string
	CSRFToken string
}

// NewClient initializes a new Proxmox API client
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
	}
}

// AuthRequest represents the request body for authentication
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Realm    string `json:"realm,omitempty"`
}

// AuthResponse represents the response from Proxmox after authentication
type AuthResponse struct {
	Data struct {
		Ticket    string `json:"ticket"`
		CSRFToken string `json:"CSRFPreventionToken"`
	} `json:"data"`
}

// Authenticate performs the authentication with Proxmox and stores the ticket and CSRF token
func (c *Client) Authenticate(username, password, realm string) error {
	authURL := fmt.Sprintf("%s/access/ticket", c.BaseURL)
	authBody := AuthRequest{Username: username, Password: password, Realm: realm}
	jsonBody, _ := json.Marshal(authBody)

	resp, err := http.Post(authURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("authentication request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("authentication failed with status code: %d", resp.StatusCode)
	}

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return fmt.Errorf("failed to decode authentication response: %v", err)
	}

	c.Ticket = authResp.Data.Ticket
	c.CSRFToken = authResp.Data.CSRFToken
	return nil
}
