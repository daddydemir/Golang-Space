package internal

import (
	"context"
	"errors"
	users "ggateway/protogen/golang/user"
	"log"
)

type UserService struct {
	collection []*users.UserWithId
	userId     *int
	users.UnimplementedUsersServer
}

func (u *UserService) GetAllUsers(_ context.Context, _ *users.Empty) (*users.UserCollection, error) {

	log.Printf("users: %v", u.collection)

	return &users.UserCollection{Users: u.collection}, nil
}

func (u *UserService) GetUser(_ context.Context, id *users.UserId) (*users.User, error) {

	for _, user := range u.collection {
		if user.Id == id.GetId() {
			log.Printf("user: %v", user)
			return user.User, nil
		}
	}

	log.Printf("user not found: %v", id)
	return nil, errors.New("user not found")
}

func (u *UserService) AddUser(_ context.Context, user *users.User) (*users.Empty, error) {

	for _, i := range u.collection {
		if i.User.Email == user.Email {
			log.Printf("user already exists: %v", i.User)
			return &users.Empty{}, errors.New("user already exists")
		}
	}

	if u.userId == nil {
		u.userId = new(int)
	}

	*u.userId = *u.userId + 1

	u.collection = append(u.collection, &users.UserWithId{
		Id:   int32(*u.userId),
		User: user,
	})

	return &users.Empty{}, nil
}

func (u *UserService) UpdateUser(_ context.Context, user *users.UserWithId) (*users.User, error) {

	for _, i := range u.collection {
		if i.Id == user.Id {
			i.User = user.User
			log.Printf("user updated: %v", i.User)
			return i.User, nil
		}
	}

	return nil, errors.New("user not found")
}

func (u *UserService) DeleteUser(_ context.Context, user *users.UserId) (*users.Empty, error) {

	for index, i := range u.collection {
		if i.Id == user.Id {
			u.collection = append(u.collection[:index], u.collection[index+1:]...)
			log.Printf("user deleted: %v", user)
			return &users.Empty{}, nil
		}
	}

	return nil, errors.New("user not found")
}
