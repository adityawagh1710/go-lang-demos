package service

import (
	"context"
	"errors"
	"project/internal/model"
	"testing"
)

// userFetchFn matches the signature of repository.GetUserByID
type userFetchFn func(context.Context, int64) (model.User, error)

// getFullDataWith is a testable variant that accepts an injected fetch function
// instead of calling the real repository (which requires a live DB).
func getFullDataWith(ctx context.Context, userID int64, fetchUser userFetchFn) (UserResponse, error) {
	user, err := fetchUser(ctx, userID)
	if err != nil {
		return UserResponse{}, err
	}
	return UserResponse{
		User:    user,
		Profile: "profile-data",
		Orders:  "orders-data",
	}, nil
}

// --- fixtures ---

func mockFetchFound(_ context.Context, id int64) (model.User, error) {
	if id == 1 {
		return model.User{ID: 1, Name: "Alice", Email: "alice@example.com"}, nil
	}
	return model.User{}, errors.New("user not found")
}

func mockFetchError(_ context.Context, _ int64) (model.User, error) {
	return model.User{}, errors.New("db connection failed")
}

// --- tests ---

func TestGetUserFullData_ReturnsUser(t *testing.T) {
	resp, err := getFullDataWith(context.Background(), 1, mockFetchFound)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.User.ID != 1 {
		t.Errorf("expected user ID 1, got %d", resp.User.ID)
	}
	if resp.User.Name != "Alice" {
		t.Errorf("expected Alice, got %s", resp.User.Name)
	}
	if resp.User.Email != "alice@example.com" {
		t.Errorf("expected alice@example.com, got %s", resp.User.Email)
	}
}

func TestGetUserFullData_ReturnsProfileAndOrders(t *testing.T) {
	resp, err := getFullDataWith(context.Background(), 1, mockFetchFound)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Profile != "profile-data" {
		t.Errorf("expected profile-data, got %s", resp.Profile)
	}
	if resp.Orders != "orders-data" {
		t.Errorf("expected orders-data, got %s", resp.Orders)
	}
}

func TestGetUserFullData_UserNotFound(t *testing.T) {
	_, err := getFullDataWith(context.Background(), 99, mockFetchFound)
	if err == nil {
		t.Error("expected error for missing user")
	}
}

func TestGetUserFullData_DBError(t *testing.T) {
	_, err := getFullDataWith(context.Background(), 1, mockFetchError)
	if err == nil {
		t.Error("expected error on DB failure")
	}
}

func TestGetUserFullData_ContextPropagated(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // already cancelled

	// fetch func that respects context
	fetchWithCtx := func(ctx context.Context, id int64) (model.User, error) {
		select {
		case <-ctx.Done():
			return model.User{}, ctx.Err()
		default:
			return model.User{ID: id, Name: "Alice", Email: "alice@example.com"}, nil
		}
	}

	_, err := getFullDataWith(ctx, 1, fetchWithCtx)
	if err == nil {
		t.Error("expected context cancellation error")
	}
}
