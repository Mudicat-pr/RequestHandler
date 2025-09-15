package storage

import "time"

type Database struct {
	Users struct {
		ID          int32
		Fullname    string
		Address     string
		Name        string
		Password    string
		Email       string
		Permissions string
		AppealsID   int32
	}

	Appeals struct {
		ID        int32
		Theme     string
		Body      string
		Status    string
		TariffID  int32
		CreatedAt time.Time
	}
	Tariffs struct {
		ID    int32
		Title string
		Body  string
	}
}
