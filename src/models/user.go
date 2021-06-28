package models

import "github.com/kamva/mgm/v3"

// UserModel used to create users with mongo
type UserModel struct {
	mgm.DefaultModel `bson:",inline"`
	Username         string `json:"username" bson:"username"`
	Password         string `json:"password" bson:"password"`
	Email            string `json:"email" bson:"email"`
}

func (model *UserModel) CollectionName() string {
	return "users"
}
