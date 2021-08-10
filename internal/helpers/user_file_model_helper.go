package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/model"
	"io/ioutil"
	"os"
	"time"
)

var formattingCharacters = []byte{',', '\n'}

func Create(modelName string, v interface{}) error {

	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(fmt.Sprintf("./datastore/%s.txt", modelName), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	f, _ := file.Stat()
	if f.Size() != 0 {
		bytes = append(formattingCharacters, bytes...)
	}
	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil

}

func Get(modelName string, email *string) (*model.User, error) {
	users, err := GetAll(modelName)
	if err != nil {
		return nil, err
	}

	for _, t := range users {
		if t.Email == *email {
			return &t, err
		}
	}

	return nil, err
}

func GetAll(modelName string) ([]model.User, error) {
	var users []model.User

	data, err := ioutil.ReadFile(fmt.Sprintf("./datastore/%s.txt", modelName))
	if err != nil {
		return nil, err
	}

	JsonMarkup(&data)

	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func Update(modelName string, user *model.User) error {
	dataUpdate := false
	users, err := GetAll(modelName)
	if err != nil {
		return err
	}

	for i, u := range users {
		if u.ID == user.ID {
			if u != *user {
				dataUpdate = true
			}
			users[i] = *user
			if dataUpdate {
				err = Save(modelName, &users)
				if err != nil {
					return err
				}
			}
			return nil
		}
	}
	return errors.New("user not found")
}

func Delete(modelName string, id int) error {
	users, err := GetAll(modelName)
	if err != nil {
		return err
	}

	for i, t := range users {
		if t.ID == int32(id) {
			users[i].Delete = time.Now().String()
			err = Save(modelName, &users)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("user not found")
}

func Save(modelName string, users *[]model.User) error {
	var bytes []byte
	for j, i := range *users {
		data, _ := json.Marshal(i)
		bytes = append(bytes, data...)

		if j < len(*users)-1 {
			bytes = append(bytes, formattingCharacters...)
		}
	}

	err := ioutil.WriteFile(fmt.Sprintf("./datastore/%s.txt", modelName), bytes, 0600)
	if err != nil {
		return err
	}
	return nil
}
func JsonMarkup(data *[]byte) {
	*data = bytes.TrimSuffix(*data, []byte(","))

	*data = append([]byte{'['}, *data...)
	*data = append(*data, ']')
}
