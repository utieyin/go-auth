package main

import "log"

// AddUser add a test user to the db
func AddUser() (User, error) {
	user := User{
		Username: "test",
		Email:    "test@test.com",
		Password: "something",
	}
	err := a.DB.Debug().Model(&User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("Cannot create seed users table: '%v'", err)
	}
	log.Printf("Test user created")
	return user, nil
}
