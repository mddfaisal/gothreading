package indexdata

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"indexing/db"
	"indexing/utils"
)

type IndexingData struct {
	Id       sql.NullInt64  `json:"id"`
	Name     sql.NullString `json:"name" faker:"name"`
	Username sql.NullString `json:"username" faker:"username"`
	Email    sql.NullString `json:"email" faker:"email"`
	Phone    sql.NullString `json:"phone" faker:"phone_number"`
	Uuid     sql.NullString `json:"uuid" faker:"uuid_hyphenated"`
	Data     sql.NullString `json:"user_data" faker:"paragraph"`
	Jwt      sql.NullString `json:"jwt" faker:"jwt"`
	Sentence sql.NullString `json:"sentence" faker:"sentence"`
}

func NewIndexingData(i int64) *IndexingData {
	d := new(IndexingData)
	d.Id.Scan(i)
	return d
}

func (i *IndexingData) values() map[string]interface{} {
	return map[string]interface{}{
		"id":        i.Id.Int64,
		"name":      i.Name.String,
		"username":  i.Username.String,
		"email":     i.Email.String,
		"phone":     i.Phone.String,
		"uuid":      i.Uuid.String,
		"user_data": i.Username.String,
		"jwt":       i.Jwt.String,
		"sentence":  i.Sentence.String,
	}
}

func (i *IndexingData) FindById() {
	id, err := i.Id.Value()
	utils.Panic(err)
	query := `select a.name, a.email, a.phone, a.uuid, a.username, b.user_data, b.sentence, b.jwt
				from users as a left join usersdata as b on a.id=b.user_id where a.id = ` + fmt.Sprintf("%v", id)
	result := db.Db().QueryRow(query)
	err = result.Scan(&i.Name, &i.Email, &i.Phone, &i.Uuid, &i.Username, &i.Data, &i.Sentence, &i.Jwt)
	if err != nil {
		fmt.Println(query)
	}
	utils.Panic(err)
}

func (i *IndexingData) AsJson() string {
	b, err := json.Marshal(i.values())
	utils.Panic(err)
	return string(b)
}
