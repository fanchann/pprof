package function

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"

	"github.com/google/uuid"
)

type Users struct {
	ID       uuid.UUID
	Username string
	Email    string
}

func WriteFileJson(n int) error {
	var users []Users
	for i := 0; i < n; i++ {
		user := Users{
			ID:       uuid.New(),                         // mem consume 9MB
			Username: fmt.Sprintf("user %d", i),          // mem consume 24.50MB
			Email:    fmt.Sprintf("user%d@gmail.com", i), // mem consume 26.50MB
		}
		users = append(users, user) // mem consume 255.58MB
	}

	jsonByte, _ := json.Marshal(&users) // 393.56MB

	return ioutil.WriteFile("gen/users.json", jsonByte, fs.ModePerm)
}

func ReadFileJson() ([]Users, int, error) {
	dataByte, err := ioutil.ReadFile("../gen/users.json") // mem consume 97.07MB
	if err != nil {
		return nil, 0, err
	}

	var users []Users
	if err := json.Unmarshal(dataByte, &users); err != nil { // mem consume 300.31MB
		return nil, 0, err
	}
	return users, len(users), nil
}
