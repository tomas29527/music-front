package models

type Banner struct {
	Id         int32  `db:"id"`
	Name       string `db:"name"`
	PicName    string `db:"pic_name"`
	Sort       string `db:"sort"`
	Url        string `db:"url"`
	DirectUrl  string `db:"direct_url"`
	Visible    int32  `db:"visible"`
	CreateTime string `db:"create_time"`
}

func QueryBannerList() (*[]Banner, error) {
	banner := []Banner{}
	e := mySqlDB.Select(&banner, "select  id, name, pic_name, sort, url,direct_url, visible, create_time from t_banner")
	if e != nil {
		return &banner, e
	}
	return &banner, nil
}
