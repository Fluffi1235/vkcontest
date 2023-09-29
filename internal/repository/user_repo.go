package repository

import (
	"database/sql"
	"github.com/Fluffi1235/vkcontest/internal/model"
	"log"
)

func (r Repository) RegistrUser(messageinfo *model.MessageInfoText) {
	rows, err := r.db.Query("SELECT * From Users where chatid = $1", messageinfo.Mi.ChatID)
	if err != nil {
		log.Println(err)
	}
	if rows.Next() {
		return
	}
	_, err = r.db.Exec("INSERT INTO Users(chatid, username, city, first_name, last_name, platform) VALUES($1, $2, $3, $4, $5, $6)",
		messageinfo.Mi.ChatID, messageinfo.UserName, "москва", messageinfo.FirstName, messageinfo.LastName, messageinfo.Mi.Platform)
	if err != nil {
		log.Println("Error inserting into dao RegistrUser")
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
		log.Println(err)
	}
	return row
}
