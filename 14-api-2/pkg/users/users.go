package users

import "errors"

type Users struct {
	data []UserInfo
}

type UserInfo struct {
	id       int
	name     string
	password string
	admin    bool
}

func New() *Users {
	u := Users{}
	u.data = append(u.data,
		UserInfo{id: 1, name: "usr1", password: "qwerty", admin: true},
		UserInfo{id: 2, name: "usr2", password: "qwerty", admin: false},
	)
	return &u
}

func (u *Users) Search(name, psw string) (*UserInfo, error) {
	for _, usr := range u.data {
		if usr.name == name && usr.password == psw {
			return &usr, nil
		}
	}
	return nil, errors.New("Wrong user name or password")
}

func (usr *UserInfo) ID() int {
	return usr.id
}

func (usr *UserInfo) Admin() bool {
	return usr.admin
}
