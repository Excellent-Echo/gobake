package user

import (
	"go-bake/entity"
	"time"
)

type UserFormat struct {
	ID        uint32 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Address   string `json:"address"`
	Phone     uint   `json:"phone"`
	Email     string `json:"email"`
}

type DeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatUser(user entity.User) UserFormat {
	var formatUser = UserFormat{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Phone:     user.Phone,
		Address:   user.Address,
		Email:     user.Email,
	}
	return formatUser

}

func FormatDeleteUser(message string) DeleteFormat {
	var deleteFormat = DeleteFormat{
		Message:    message,
		TimeDelete: time.Now(),
	}
	return deleteFormat
}
