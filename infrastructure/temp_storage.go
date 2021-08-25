package infrastructure

import (
	"backgammonclient/utils"
	"sync"
)

type UserData struct {
	Index int
	Username string
	Password string
	Token string
}

type AuthStorage struct {
	storage []UserData
	mutex sync.Mutex
}

func NewAuthStorage() *AuthStorage {
	return &AuthStorage{
		storage: make([]UserData, 0, 1024),
	}
}

func (a *AuthStorage) AddUser(username, password, token string) {
	a.mutex.Lock()
	a.storage = append(a.storage, UserData{Username: username, Password: password, Token: token})
	a.mutex.Unlock()
}

func (a *AuthStorage) GetRandomUser() (UserData, int) {
	var temp UserData
	i := utils.RandomInt(len(a.storage) - 1)
	a.mutex.Lock()
	temp = a.storage[i]
	a.mutex.Unlock()
	return temp, i

}

func (a *AuthStorage) UpdateToken(index int, token string) {
	a.mutex.Lock()
	a.storage[index].Token = token
	a.mutex.Unlock()
}