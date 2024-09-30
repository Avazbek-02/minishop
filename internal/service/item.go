package service

import "database/sql"

type ItemService interface {
}

type itemServiceStr struct {
	DB *sql.DB
}

func (s *itemServiceStr) CreateItem(item string) error {
	return nil
}

func (s *itemServiceStr) UpdateItem(item string) error {
	return nil
}

func (s *itemServiceStr) DeleteItem(item string) error {
	return nil
}

func (s *itemServiceStr) GetItem(item string) error {
	return nil
}

func (s *itemServiceStr) GetAllItem(item string) error {
	return nil
}
