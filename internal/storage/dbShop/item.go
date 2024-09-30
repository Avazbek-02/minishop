package dbShop

import (
	"database/sql"
	"fmt"
	pb "github.com/minishop/genproto/shop"
	"github.com/minishop/internal/logger"
	"github.com/spf13/cast"
	"go.uber.org/zap"

	"github.com/google/uuid"
)

type ItemInt interface {
}

type ItemStr struct {
	DB *sql.DB
}

func NewItemStr(db *sql.DB) *ItemStr {
	return &ItemStr{
		DB: db,
	}
}

func (s *ItemStr) CreateItem(req *pb.CreateItemReq) (*pb.CreateItemRes, error) {
	log, _ := logger.NewLogger()
	query := `
		insert into shop(id,name,img_url,categorys,user_name,user_phone)
		values ($1, $2, $3, $4, $5,$6)
	`

	id := uuid.NewString()
	_, err := s.DB.Exec(query, id, req.Item.Name, req.Item.ImgUrl, req.Item.Category, req.Item.UserName, req.Item.UserPhone)
	if err != nil {
		log.Error("Error with create table", zap.Error(err))
		return nil, err
	}

	return &pb.CreateItemRes{Message: "Item Created"}, nil
}

func (s *ItemStr) UpdateItem(req *pb.UpdateItemReq) (*pb.UpdateItemRes, error) {

	log, _ := logger.NewLogger()

	query := "update shop set updated_at = now()"
	filter := ""
	count := 1

	if req.Updateitem.Name != "" {
		filter += req.Updateitem.Name + "=$" + cast.ToString(count) + ", "
		count++
	}

	if req.Updateitem.ImgUrl != "" {
		filter += req.Updateitem.ImgUrl + "=$" + cast.ToString(count) + ", "
		count++
	}

	if req.Updateitem.Category != "" {
		filter += req.Updateitem.Category + "=$" + cast.ToString(count) + ", "
		count++
	}

	if req.Updateitem.UserName != "" {
		filter += req.Updateitem.UserName + "=$" + cast.ToString(count) + ", "
		count++
	}

	if req.Updateitem.UserPhone != "" {
		filter += req.Updateitem.UserPhone + "=$" + cast.ToString(count) + ", "
		count++
	}
	if count == 1 {
		log.Error("Error updating item")
		return nil, nil
	}
	query += ", " + filter + "where id=$1 and deleted_at = 0"
	fmt.Println(query)

	_, err := s.DB.Exec(query, req.Updateitem.Name, req.Updateitem.ImgUrl, req.Updateitem.Category, req.Updateitem.UserName, req.Updateitem.UserPhone)
	if err != nil {
		log.Error("Error with update table", zap.Error(err))
		return nil, err
	}
	log.
	return &pb.UpdateItemRes{Message: "Item Updated"}, nil
}

func (s *ItemStr) DeleteItem(req *pb.DeleteItemReq) (*pb.DeleteItemRes, error) {
	query := `
		update shop set deleted_at = extract(epoch from now()) where id = $1 and deleted_at = 0
		`
	log, _ := logger.NewLogger()
	_, err := s.DB.Exec(query, req.Id)
	if err != nil {
		log.Error("Error with delete table", zap.Error(err))
		return nil, err
	}
	return &pb.DeleteItemRes{Message: "Item Deleted"}, nil
}

func (s *ItemStr) GetItem(item ItemInt) error {
	return nil
}

func (s *ItemStr) GetAllItem(item ItemInt) error {
	return nil
}
