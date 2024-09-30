package dbShop

import (
	"database/sql"
)

type ItemInt interface {
}

type ItemStr struct {
	DB *sql.DB
}

//func (s *ItemStr) CreateItem(req *pb.CreateItemReq) (*pb.CreateItemRes, error) {
//	query := `
//		insert into shop(id,name,img_url,cat)
//	`
//
//	return nil
//}

func (s *ItemStr) UpdateItem(item ItemInt) error {
	return nil
}

func (s *ItemStr) DeleteItem(item ItemInt) error {
	return nil
}

func (s *ItemStr) GetItem(item ItemInt) error {
	return nil
}

func (s *ItemStr) GetAllItem(item ItemInt) error {
	return nil
}
