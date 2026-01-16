// repository pattern

package main

import (
	"errors"
)

// Defining a technical abstraction/dependency that provides data access for UserService below
type UserRepository interface {
    Save(user MyUser) error
    FindByID(id int) (MyUser, error)
}

type MyUser struct {
    Email string
    ID int
}

type InMemoryUserRepo struct {
    UserList map[int]MyUser
}

// pointers are used to ensure shared state and consistent behavior across method calls
func (r *InMemoryUserRepo) Save(u MyUser) error {
    if u.Email == "" {
        return errors.New("User email cannot be empty")
    }
    r.UserList[u.ID] = u
    return nil
}

// pointers are used to ensure shared state and consistent behavior across method calls
func (r *InMemoryUserRepo) FindByID(ID int) (MyUser, error) {

    foundUser, ok := r.UserList[ID]

    if !ok {
        return MyUser{}, errors.New("User not found")
    }
    return foundUser, nil 
}

/* Defining a service that defines business behavior (rules like duplicate email, workflows, coordination)
UserRepository remains the technical abstraction while UserService allows to add more validation
and expand business logic.
UserService uses UserRepository to access Save() and FindById
*/
type UserService struct {
		repo UserRepository
	}

func (s *UserService) Register(user MyUser) error {
    // Adding a new business logic of checking email existance before registering a user
    existing, _ := s.repo.FindByID(user.ID)
    if existing.ID != 0 {
        return errors.New("email already in use")
    }

    err := s.repo.Save(user)
    if err != nil {
        return err
    }
    return nil
}

func (s *UserService) GetUserById(ID int) (MyUser, error) {
    existing, err := s.repo.FindByID(ID)
    if existing.ID != 0 {
        return MyUser{}, err
    }
    return existing, nil
}