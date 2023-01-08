package services

import "awesomeProject/internal/entities"

type TableService interface {
	CreateTable() int
	GetAllTables() []entities.Table
}

func (c entities.Table) CreateTable() int {

}

func (t entities.Table) GetAllTables() []entities.Table {

}
