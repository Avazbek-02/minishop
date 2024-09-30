package dbShop

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	pb "github.com/minishop/genproto/shop"
	"github.com/minishop/internal/logger"
	"github.com/spf13/cast"
	"go.uber.org/zap"
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
	log.Named("Successfully item updated")
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

func (s *ItemStr) GetItem(req *pb.GetItemReq) (*pb.GetItemRes, error) {
	query := `select * from shop where id = $1`
	log, _ := logger.NewLogger()

	resp := &pb.GetItemRes{}

	err := s.DB.QueryRow(query, req.Id).Scan(&resp.Item.Id, &resp.Item.Name, &resp.Item.ImgUrl, &resp.Item.Category,
		&resp.Item.UserName, &resp.Item.UserPhone, &resp.Item.CreatedAt, &resp.Item.UpdatedAt)
	if err != nil {
		log.Error("Error with get table", zap.Error(err))
		return nil, err
	}
	log.Named("Successfully get item")

	return resp, nil

}

func (s *ItemStr) GetAllItem(req *pb.GetAllItemReq) (*pb.GetAllItemRes, error) {
	query := "select id,name,img_url,categorys,user_name,user_phone, created_at,updated_at from shop where deleted_at = 0"
	log, _ := logger.NewLogger()
	filter := ""
	count := 1
	if req.Id != "" && req.Id != "string" {
		filter += " and id =" + req.Id
	}
	if req.Name != "" && req.Name != "string" {
		filter += " and name =" + req.Name
	}
	if req.ImgUrl != "" && req.ImgUrl != "string" {
		filter += " and img_url =" + req.ImgUrl
	}
	if req.UserName != "" && req.UserName != "string" {
		filter += " and user_name =" + req.UserName
	}
	if req.UserPhone != "" && req.UserPhone != "string" {
		filter += "and user_phone =" + req.UserPhone
	}

	res := []*pb.ItemModel{}
	rows, err := s.DB.Query(query, filter)
	if err != nil {
		log.Error("Error with get table", zap.Error(err))
	}

	for rows.Next() {
		item := &pb.ItemModel{}
		err := rows.Scan(&item.Id, &item.Name, &item.ImgUrl, &item.Category, &item.UserName, &item.UserPhone, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			log.Error("Error with get table", zap.Error(err))
		}
		res = append(res, item)
	}
	return &pb.GetAllItemRes{Items: res}, err
}
