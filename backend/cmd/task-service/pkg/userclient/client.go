package userclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"task-management-task-service/configs"
)

// UserClient is a client for interacting with the user service.
type UserClient struct{
	cfg *configs.Config
}

// NewUserClient creates a new UserClient instance.
func NewUserClient(cfg *configs.Config) *UserClient {
	return &UserClient{cfg: cfg}
}

func (c *UserClient) GetUsersByIDs(userIDs []uint) ([]UserInfo, error) {
	url := c.cfg.UserService
	if len(userIDs) == 0 {
		return nil, errors.New("no user IDs provided")
	}
	fmt.Println("UserService URL:", url)

	reqBody, err := json.Marshal(map[string][]uint{"user_id": userIDs})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to make POST request: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var res GetUsersResponse
	if err := json.Unmarshal(bodyBytes, &res); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return res.Data, nil
}
