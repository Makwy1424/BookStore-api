package model

type Book struct {
	ID       string `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	ISBN     string `json:"isbn"`
	Genre    string `json:"genre"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	Desc     string `json:"desc"`
}
