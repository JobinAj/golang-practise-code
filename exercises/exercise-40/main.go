package main

import (
	"fmt"
	"log"
)

// User represents  a user with a id and first name
type User struct {
	ID    int
	First string
}

// Datastore defines an interface for storing retriviable data
// Any Type that implements this two methods is also a type of datastore
// this means any value of type mockdatastore is also a type of datastore
type Datastore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

// MockDatasotre is a temporary service that stores retrivable data
type MockDatastore struct {
	Users map[int]User
}

// This is the method to get user
func (md MockDatastore) GetUser(id int) (User, error) {
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf("user %v not found", id)
	}
	return user, nil
}

// Here we are checking whether the user already exist if doesn't exist we are entering the user info in the mockdatastore map.
func (md MockDatastore) SaveUser(u User) error {
	_, ok := md.Users[u.ID]
	if ok {
		return fmt.Errorf("The user %v already exist", u.ID)
	}
	md.Users[u.ID] = u
	return nil
}

// Here we are creating service type which has ds as variable  and datastore as type so datastore is a interface which service should satisfy two methods.
type Service struct {
	ds Datastore
}

// function service for talking with database for save useer
func (s Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}

func (s Service) SaveUser(u User) error {
	return s.ds.SaveUser(u)
}

func main() {
	db := MockDatastore{
		make(map[int]User),
	}
	srvc := Service{
		ds: db,
	}

	u1 := User{
		ID:    1,
		First: "Jobin",
	}

	err := srvc.SaveUser(u1)
	if err != nil {
		log.Fatalf("error %s", err)
	}
	u1Returned, err := srvc.GetUser(u1.ID)
	if err != nil {
		log.Fatalf("error %s", err)
	}
	fmt.Println(u1)
	fmt.Println(u1Returned)

}
