package repository

import (
	"github.com/Fluffi1235/vkcontest/internal/model"
)

func (r Repository) SetUser(messageInfo *model.MessageInfoText) error {
	rows, err := r.db.Queryx("SELECT chatid, userName, city, firstName, lastName, platform From Users where chatid = $1", messageInfo.MessageInfo.ChatID)
	if err != nil {
		return err
	}
	if rows.Next() {
		return nil
	}
	query := "INSERT INTO Users(chatid, userName, city, firstName, lastName, platform) VALUES($1, $2, $3, $4, $5, $6)"
	_, err = r.db.Exec(query,
		messageInfo.MessageInfo.ChatID, messageInfo.UserName, "москва", messageInfo.FirstName, messageInfo.LastName, messageInfo.MessageInfo.Platform)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) UpdateCityOfUser(city string, chatId int64) error {
	query := "UPDATE users set city = $1 where chatId = $2"
	_, err := r.db.Exec(query, city, chatId)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) GetUserData(chatId int64, platform string) (*model.User, error) {
	userData := &model.User{}
	query := "SELECT chatId, userName, firstName, lastName, city, platform FROM users where chatid = $1 and platform = $2"
	err := r.db.Get(userData, query, chatId, platform)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (r Repository) GetCityOfUser(chatId int64) (string, error) {
	userData := &model.User{}
	query := "SELECT city FROM users where chatid = $1"
	err := r.db.Get(userData, query, chatId)
	if err != nil {
		return "", err
	}
	return userData.City, nil
}
