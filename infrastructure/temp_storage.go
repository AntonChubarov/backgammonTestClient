package infrastructure

import (
	"sync"
)

type UserData struct {
	Username string
	Password string
	Token string
}

type AuthStorage struct {
	storage map[string]UserData
	mutex sync.Mutex
}

func NewAuthStorage() *AuthStorage {
	return &AuthStorage{
		storage: make(map[string]UserData),
	}
}

func (a *AuthStorage) AddUser(username, password, token string) {
	a.mutex.Lock()
	a.storage[username] = UserData{Username: username, Password: password, Token: token}
	a.mutex.Unlock()
}

func (a *AuthStorage) GetRandomUser() UserData {
	var temp UserData
	a.mutex.Lock()
	for _, user := range a.storage {
		temp = user
		break
	}
	a.mutex.Unlock()
	return temp

}

func (a *AuthStorage) UpdateToken(username, token string) {
	a.mutex.Lock()
	temp := a.storage[username]
	temp.Token = token
	a.storage[username] = temp
	a.mutex.Unlock()
}