package models

type Recommend struct {
	Id         uint32 `db:"id"`
	VideoId    uint32 `db:"video_id"`
	VideoTitle string `db:"video_title"`
	Title      string `db:"title"`
	SecTitle   string `db:"sec_title"`
	Sort       uint32 `db:"sort"`
	CreateTime string `db:"create_time"`
	Visible    uint32 `db:"visible"`
	Type       uint32 `db:"type"`
}

func QueryRecommendList() (*[]Recommend, error) {
	recommend := []Recommend{}
	e := mySqlDB.Select(&recommend, "select  id, video_id, video_title, title, sec_title,sort, create_time, visible,type from t_hot_recommend")
	if e != nil {
		return &recommend, e
	}
	return &recommend, nil
}
