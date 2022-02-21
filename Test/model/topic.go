package model

type MessageInfo struct {
	Message string `db:message`
	Username string `db:username`
}
type Topic struct {
	Content string `db:content`
	Name string `db:name`
	Num string `db:num`
	Author string `db:author`
}
type Message struct {
	Username string `db:name`
	Message string `db:message`
	Num string `db:num`
	Like int `db:like`

}