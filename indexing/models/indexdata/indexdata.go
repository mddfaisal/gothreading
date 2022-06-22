package indexdata

import (
	"encoding/json"
	"fmt"
	"indexing/db"
	"indexing/utils"
)

type IndexingData struct {
	Id       int64  `json:"id"`
	Name     string `json:"name" faker:"name"`
	Username string `json:"username" faker:"username"`
	Email    string `json:"email" faker:"email"`
	Phone    string `json:"phone" faker:"phone_number"`
	Uuid     string `json:"uuid" faker:"uuid_hyphenated"`
	Data     string `json:"user_data" faker:"paragraph"`
	Jwt      string `json:"jwt" faker:"jwt"`
	Sentence string `json:"sentence" faker:"sentence"`
}

func NewIndexingData(i int64) *IndexingData {
	d := new(IndexingData)
	d.Id = i
	return d
}

func (i *IndexingData) FindById() {
	query := `select a.name, a.email, a.phone, a.uuid, b.user_data, b.sentence 
				from users as a inner join usersdata as b on a.id=b.user_id where a.id = ` + fmt.Sprintf("%v", i.Id)
	result := db.Db().QueryRow(query)
	err := result.Scan(&i.Name, &i.Email, &i.Phone, &i.Uuid, &i.Data, &i.Sentence)
	utils.Panic(err)
}

func (i *IndexingData) AsJson() string {
	d, err := json.Marshal(i)
	utils.Panic(err)
	return string(d)
}
