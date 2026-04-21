package service

import (
	"errors"
	"testing"

	"example4_fiber_clean_orm/models"
)

// mockUserRepo is an in-memory implementation of repository.UserRepository
type mockUserRepo struct {
	users  map[uint]*models.User
	nextID uint
}

func newMockRepo() *mockUserRepo {
	return &mockUserRepo{users: make(map[uint]*models.User), nextID: 1}
}

func (m *mockUserRepo) Create(user *models.User) error {
	user.ID = m.nextID
	m.nextID++
	m.users[user.ID] = user
	return nil
}

func (m *mockUserRepo) FindAll() ([]models.User, error) {
	list := make([]models.User, 0, len(m.users))
	for _, u := range m.users {
		list = append(list, *u)
	}
	return list, nil
}

func (m *mockUserRepo) FindByID(id uint) (*models.User, error) {
	u, ok := m.users[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return u, nil
}

func (m *mockUserRepo) Update(user *models.User) error {
	if _, ok := m.users[user.ID]; !ok {
		return errors.New("not found")
	}
	m.users[user.ID] = user
	return nil
}

func (m *mockUserRepo) Delete(id uint) error {
	if _, ok := m.users[id]; !ok {
		return errors.New("not found")
	}
	delete(m.users, id)
	return nil
}

// --- CreateUser ---

func TestCreateUser_Success(t *testing.T) {
	svc := NewUserService(newMockRepo())
	if err := svc.CreateUser(&models.User{Name: "Alice", Email: "alice@example.com"}); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCreateUser_MissingEmail(t *testing.T) {
	svc := NewUserService(newMockRepo())
	if err := svc.CreateUser(&models.User{Name: "Alice"}); err == nil {
		t.Error("expected error for missing email")
	}
}

func TestCreateUser_EmptyEmail(t *testing.T) {
	svc := NewUserService(newMockRepo())
	if err := svc.CreateUser(&models.User{Name: "Bob", Email: ""}); err == nil {
		t.Error("expected error for empty email")
	}
}

// --- GetUsers ---

func TestGetUsers_Empty(t *testing.T) {
	svc := NewUserService(newMockRepo())
	users, err := svc.GetUsers()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(users) != 0 {
		t.Errorf("expected 0 users, got %d", len(users))
	}
}

func TestGetUsers_ReturnsAll(t *testing.T) {
	svc := NewUserService(newMockRepo())
	svc.CreateUser(&models.User{Name: "Alice", Email: "alice@example.com"})
	svc.CreateUser(&models.User{Name: "Bob", Email: "bob@example.com"})

	users, err := svc.GetUsers()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}
}

// --- GetUser ---

func TestGetUser_Found(t *testing.T) {
	svc := NewUserService(newMockRepo())
	svc.CreateUser(&models.User{Name: "Alice", Email: "alice@example.com"})

	user, err := svc.GetUser(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.Name != "Alice" {
		t.Errorf("expected Alice, got %s", user.Name)
	}
	if user.Email != "alice@example.com" {
		t.Errorf("expected alice@example.com, got %s", user.Email)
	}
}

func TestGetUser_NotFound(t *testing.T) {
	svc := NewUserService(newMockRepo())
	if _, err := svc.GetUser(99); err == nil {
		t.Error("expected error for missing user")
	}
}

// --- UpdateUser ---

func TestUpdateUser_Success(t *testing.T) {
	svc := NewUserService(newMockRepo())
	svc.CreateUser(&models.User{Name: "Alice", Email: "alice@example.com"})

	if err := svc.UpdateUser(1, &models.User{Name: "Alice Updated", Email: "new@example.com"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	user, _ := svc.GetUser(1)
	if user.Name != "Alice Updated" {
		t.Errorf("expected Alice Updated, got %s", user.Name)
	}
	if user.Email != "new@example.com" {
		t.Errorf("expected new@example.com, got %s", user.Email)
	}
}

func TestUpdateUser_NotFound(t *testing.T) {
	svc := NewUserService(newMockRepo())
	if err := svc.UpdateUser(99, &models.User{Name: "Ghost", Email: "ghost@example.com"}); err == nil {
		t.Error("expected error for missing user")
	}
}

// --- DeleteUser ---

func TestDeleteUser_Success(t *testing.T) {
	svc := NewUserService(newMockRepo())
	svc.CreateUser(&models.User{Name: "Alice", Email: "alice@example.com"})

	if err := svc.DeleteUser(1); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, err := svc.GetUser(1); err == nil {
		t.Error("expected error after deletion")
	}
}

func TestDeleteUser_NotFound(t *testing.T) {
	svc := NewUserService(newMockRepo())
	if err := svc.DeleteUser(99); err == nil {
		t.Error("expected error for missing user")
	}
}

func TestDeleteUser_DoesNotAffectOthers(t *testing.T) {
	svc := NewUserService(newMockRepo())
	svc.CreateUser(&models.User{Name: "Alice", Email: "alice@example.com"})
	svc.CreateUser(&models.User{Name: "Bob", Email: "bob@example.com"})

	svc.DeleteUser(1)

	if _, err := svc.GetUser(2); err != nil {
		t.Error("Bob should still exist after deleting Alice")
	}
}
