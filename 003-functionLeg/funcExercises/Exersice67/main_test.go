package main

import "testing"

func TestGetUser(t *testing.T) {
	db := &MockDatastore{
		Users: map[int]User{
			003: {
				ID:    003,
				First: "Amr",
			},
		},
	}

	service := Service{
		ds: db,
	}

	user, err := service.GetUser(003)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	if user.First != "Amr" {
		t.Errorf("got %s, want %s", user.First, "Amr")
	}

}
