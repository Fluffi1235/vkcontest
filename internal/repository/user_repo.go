package repository

import (
	"database/sql"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"log"
)

func (r Repository) RegistrationUser(messageinfo *model.MessageInfoText) {
	rows, err := r.db.Query("SELECT chatid, username, city, first_name, last_name, platform From Users where chatid = $1", messageinfo.MessageInfo.ChatID)
	if err != nil {
		log.Println(err, " dao RegistrationUser")
	}
	if rows.Next() {
		return
	}
	_, err = r.db.Exec("INSERT INTO Users(chatid, username, city, first_name, last_name, platform) VALUES($1, $2, $3, $4, $5, $6)",
		messageinfo.MessageInfo.ChatID, messageinfo.UserName, "москва", messageinfo.FirstName, messageinfo.LastName, messageinfo.MessageInfo.Platform)
	if err != nil {
		log.Println("Error inserting into dao RegistrationUser")
	}
}

func (r Repository) CityChange(city string, chatid int64) {
	_, err := r.db.Exec("UPDATE users set city = $1 where chatid = $2", city, chatid)
	if err != nil {
		log.Println("Error inserting into dao CityChange")
	}
}

func (r Repository) GetUserData(chatId int64, platform string) *sql.Rows {
	row, err := r.db.Query("SELECT * FROM users where chatid = $1 and platform = $2", chatId, platform)
	if err != nil {
		log.Println(err)
	}
	return row
}

func (r Repository) GetCityOfUser(chatId int64) *sql.Rows {
	row, err := r.db.Query("SELECT city FROM users where chatid = $1", chatId)
	if err != nil {
		log.Println(err, " dao GetCityOfUser")
	}
	return row
}
