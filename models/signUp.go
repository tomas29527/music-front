package models

import (
	"music-front/vo"
	"time"
)

type SignUp struct {
	Id          uint32 `db:id`
	Name        string `db:"name"`
	Age         uint32 `db:"age"`
	Sex         uint32 `db:"sex"`
	Phone       string `db:"phone"`
	Create_time string
}

func (s *SignUp) InsertSignUp() error {
	sql := "insert into t_sign_up (name,age,sex,phone,create_time) values (?,?,?,?,?)"
	_, e := mySqlDB.Exec(sql, s.Name, s.Age, s.Sex, s.Phone, time.Now())
	return e
}

func NewSignUp(v *vo.SignUpVo) *SignUp {
	signUp := &SignUp{
		Name:  v.Name,
		Age:   v.Age,
		Sex:   v.Sex,
		Phone: v.Phone,
	}
	return signUp
}
