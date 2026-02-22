package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
{
    "id": "u_12345",
    "email_address": "alice@example.com",
    "first_name": "Alice",
    "last_name": "Johnson",
    "age": 29,
    "is_active": true,
    "created_at": "2024-05-01T12:30:00Z",
    "last_login_at": "2024-06-01T10:00:00Z",
    "metadata": {
        "plan": "pro",
        "region": "eu"
    }
}
*/

func main() {

	raw := []byte(`{
	    "id": "u_12345",
	    "email_address": "alice@example.com",
	    "first_name": "Alice",
	    "last_name": "Johnson",
	    "age": 29,
	    "is_active": true,
	    "created_at": "2024-05-01T12:30:00Z",
	    "last_login_at": "2024-06-01T10:00:00Z",
	    "metadata": {
	        "plan": "pro",
	        "region": "eu"
	    }
	}`)
	var u userResponse

	if err := json.Unmarshal(raw, &u); err != nil {
		log.Fatal("Error unmarshaling JSON:", err)
	}

	fc := FakeHTTPClient{
		users: map[string]userResponse{
			u.ID: u,
		},
	}
	srvc := Service{
		users: &fc,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Println(srvc.ProcessUser(ctx, "u_12345"))

}

// Faking HTTP layer to receive response without calling real API.
type FakeHTTPClient struct {
	users map[string]userResponse
}

func (fake *FakeHTTPClient) GetUser(ctx context.Context, id string) (userResponse, error) {

	u, ok := fake.users[id]
	if !ok {
		return userResponse{}, errors.New("user not found")
	}
	return u, nil

}

type User struct {
	ID        string
	Email     string
	FullName  string
	Active    bool
	Plan      string
	CreatedAt time.Time
}

type userResponse struct {
	ID          string            `json:"id"`
	Email       string            `json:"email_address"`
	FirstName   string            `json:"first_name"`
	LastName    string            `json:"last_name"`
	Age         int               `json:"age"`
	IsActive    bool              `json:"is_active"`
	CreatedAt   string            `json:"created_at"`
	LastLoginAt string            `json:"last_login_at"`
	Metadata    map[string]string `json:"metadata"`
}

// Business layer
type Service struct {
	users UserClient // interface exposes methods from the HTTP layer declared in UserClient
}

// Business layer
func (s *Service) ProcessUser(ctx context.Context, id string) (User, error) {
	data, err := s.users.GetUser(ctx, id)
	if err != nil {
		return User{}, err
	}

	createdAt, err := time.Parse("2006-01-02T15:04:05Z07:00", data.CreatedAt)
	if err != nil {
		return User{}, err
	}

	return User{
		ID:        data.ID,
		Email:     data.Email,
		FullName:  data.FirstName + data.LastName,
		Active:    data.IsActive,
		CreatedAt: createdAt,
	}, nil
}

// Interface exposed to Service
type UserClient interface {
	GetUser(ctx context.Context, id string) (userResponse, error)
}

// HTTP layer
type HTTPClient struct {
	baseUrl string
	client  *http.Client
}

// HTTP layer
func (c *HTTPClient) GetUser(ctx context.Context, id string) (userResponse, error) {

	// Creating request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseUrl, nil)
	if err != nil {
		return userResponse{}, err
	}

	// Sending request with prepared body
	resp, err := c.client.Do(req)
	if err != nil {
		return userResponse{}, err
	}
	defer resp.Body.Close()

	// Verifying response status
	if resp.StatusCode != http.StatusOK {
		return userResponse{}, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	// Parsing response into u User
	var u userResponse
	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		return userResponse{}, err
	}

	return u, nil

}
