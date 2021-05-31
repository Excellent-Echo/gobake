package user

import (
	"errors"
	"fmt"
	"go-bake/entity"
	"go-bake/helper"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Service interface {
	GetAllUser() ([]UserFormat, error)
	SaveNewUser(user entity.UserInput) (UserFormat, error)
	GetUserByID(ID string) (UserFormat, error)
	UpdateUserByID(ID string, dataInput entity.UpdateUserInput) (UserFormat, error)
	DeleteByUserID(ID string) (interface{}, error)
	LoginUser(input entity.LoginUserInput) (entity.User, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) LoginUser(input entity.LoginUserInput) (entity.User, error) {
	user, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %v not found", user.ID)
		return user, errors.New(newError)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return user, errors.New("Invalid password")
	}

	return user, nil
}

func (s *service) GetAllUser() ([]UserFormat, error) {
	//business logic
	users, err := s.repository.FindAll()
	var formatUsers []UserFormat

	for _, user := range users {
		formatUser := FormatUser(user)
		formatUsers = append(formatUsers, formatUser)
	}

	if err != nil {
		return formatUsers, err
	}

	return formatUsers, nil
}

func (s *service) SaveNewUser(user entity.UserInput) (UserFormat, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return UserFormat{}, err
	}

	var newUser = entity.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Username:  user.Username,
		Phone:     user.Phone,
		Address:   user.Address,
		Password:  string(genPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	createUser, err := s.repository.Create(newUser)
	formatUser := FormatUser(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil

}

func (s *service) GetUserByID(ID string) (UserFormat, error) {
	if err := helper.ValidateIDNumber(ID); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.FindByID(ID)

	if err != nil {
		return UserFormat{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %d not found", user.ID)

		return UserFormat{}, errors.New(newError)
	}

	formatUser := FormatUser(user)
	return formatUser, nil
}

func (s *service) UpdateUserByID(ID string, dataInput entity.UpdateUserInput) (UserFormat, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(ID); err != nil {
		return UserFormat{}, err
	}

	user, err := s.repository.FindByID(ID)

	if err != nil {
		return UserFormat{}, err

	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s not found", ID)
		return UserFormat{}, errors.New(newError)
	}

	if dataInput.FirstName != "" || len(dataInput.FirstName) != 0 {
		dataUpdate["first_name"] = dataInput.FirstName
	}

	if dataInput.LastName != "" || len(dataInput.LastName) != 0 {
		dataUpdate["last_name"] = dataInput.LastName
	}

	if dataInput.Email != "" || len(dataInput.Email) != 0 {
		dataUpdate["email"] = dataInput.Email
	}

	if dataInput.Username != "" || len(dataInput.Username) != 0 {
		dataUpdate["username"] = dataInput.Username
	}

	if dataInput.Address != "" || len(dataInput.Address) != 0 {
		dataUpdate["email"] = dataInput.Address
	}

	dataUpdate["updated_at"] = time.Now()

	// fmt.Println(dataUpdate)

	userUpdate, err := s.repository.UpdateByID(ID, dataUpdate)
	if err != nil {
		return UserFormat{}, err
	}

	formatUser := FormatUser(userUpdate)

	return formatUser, nil
}

func (s *service) DeleteByUserID(ID string) (interface{}, error) {
	if err := helper.ValidateIDNumber(ID); err != nil {
		return nil, err
	}

	user, err := s.repository.FindByID(ID)

	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %d not found", user.ID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteByID(ID)

	if status == "error" {
		return nil, errors.New("error delete")
	}

	msg := fmt.Sprintf("succes delete user ID : %s", ID)

	formatDelete := FormatDeleteUser(msg)

	return formatDelete, nil

}
