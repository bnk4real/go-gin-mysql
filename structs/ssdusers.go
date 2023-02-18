package structs

func (ssdusers) TableName() string {
	return "tableName"
}

type ssdusers struct {
	UserId      string
	FulllNameTh string
}
