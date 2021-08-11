package models

import (
	"connect_db_exc/entities"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserModel struct {
	DB         *mgo.Database
	Collection string
}

func (userModel UserModel) FindAllUsers() (users []entities.User, err error) {
	err = userModel.DB.C(userModel.Collection).Find(bson.M{}).All(&users)
	return
}

func (userModel UserModel) FindUser(id string) (user entities.User, err error) {
	err = userModel.DB.C(userModel.Collection).FindId(bson.ObjectIdHex(id)).One(&user)
	return
}

func (userModel UserModel) CreateUser(user *entities.User) (err error) {
	err = userModel.DB.C(userModel.Collection).Insert(&user)
	return
}

func (userModel UserModel) UpdateUser(user *entities.User) (err error) {
	err = userModel.DB.C(userModel.Collection).UpdateId(user.ID, &user)
	return
}

func (userModel UserModel) DeleteUser(user entities.User) (err error) {
	err = userModel.DB.C(userModel.Collection).Remove(user)
	return
}
