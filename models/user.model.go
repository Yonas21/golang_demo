package models

import (
	db "example/hello/config"
	s "example/hello/serialize_data"
	"fmt"

	"github.com/pkg/errors"
)

type Product struct {
	Name      string `json:"name"`
	Price     int    `json:"price"`
	OrderedBy string `json:"ordered_by"`
}

func CreateUser(person *s.Person) (string, error) {

	database, err := db.Connect()

	if err != nil {
		return "", errors.WithMessage(err, "error with connection")
	}

	creation_stat := database.Create(&person)
	message := "row affected:" + fmt.Sprintf("%d", creation_stat.RowsAffected)
	return message, nil
}

func GetAllUsers() ([]s.Person, error) {

	var users []s.Person

	database, err := db.Connect()

	if err != nil {
		return nil, errors.WithMessage(err, "error with connection")
	}

	// get all data
	database.Order("age").Find(&users)
	return users, nil
}

// get data by id
func GetUserById(userId int) (s.Person, error) {
	var user s.Person

	database, err := db.Connect()

	if err != nil {
		return s.Person{}, errors.WithMessage(err, "error with connection")
	}

	database.AutoMigrate(&Product{})
	// get all data
	database.First(&user, userId)
	return user, nil

}

func SelectSome(name ...string) ([]s.Person, error) {
	var users []s.Person

	database, err := db.Connect()

	if err != nil {
		return nil, errors.WithMessage(err, "error with connection")
	}

	// get all data
	database.Limit(3).Select("name").Find(&users)
	return users, nil
}

func GetGrouped() ([]s.Person, error) {
	var users []s.Person

	database, err := db.Connect()

	if err != nil {
		return nil, errors.WithMessage(err, "error with connection")
	}

	// get all data
	database.Model(&s.Person{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "group").Find(&users)
	return users, nil
}

func UpdatePerson(id int, updatesting string) error {
	database, err := db.Connect()

	if err != nil {
		return errors.WithMessage(err, "error with connection")
	}
	database.Model(&s.Person{}).Where("id = ?", id).Update("name", updatesting)

	return nil

}

func DeletePerson(id int) error {

	var user s.Person

	database, err := db.Connect()

	if err != nil {
		return errors.WithMessage(err, "error with connection")
	}
	database.Where("id=?", id).Delete(&user)

	return nil

}
