package services

import (
	"github.com/yz124/superstar/dao"
	"github.com/yz124/superstar/models"
	"github.com/yz124/superstar/datasource"
)

//对外暴露
//service层就是对Model层的封装
type SuperstarService interface {
	GetAll() []models.StarInfo
	Search(country string) []models.StarInfo
	Get(id int) *models.StarInfo
	Delete(id int) error
	Update(user *models.StarInfo, columns []string) error
	Create(user *models.StarInfo) error
}

type SuperstarServicee interface {
	GetAll() []models.StarInfo
	Search(country string) []models.StarInfo
	Get(id int) *models.StarInfo
	Delete(id int) error
	Update(user *models.StarInfo, columns []string) error
	Create(user *models.StarInfo) error
}
//内部结构
type superstarService struct {
	dao *dao.SuperstarDao
}

type superstarServicee struct {
	dao *dao.SuperstarDao
}

func NewSuperstarService() SuperstarService {
	return &superstarService {
		//引擎来自于主Master库
		dao: dao.NewSuperstarDao(datasource.InstanceMaster()),
	}
}

func NewSuperstarServicee() SuperstarServicee {
	return &superstarServicee{
		dao:dao.NewSuperstarDao(datasource.InstanceMasterr()),
	}
}

func (s *superstarService)GetAll() []models.StarInfo {
	return s.dao.GetAll()
}

func (s *superstarServicee) GetAll() []models.StarInfo {
	return s.dao.GetAll()
}

func (s *superstarService)Search(country string) []models.StarInfo {
	return s.dao.Search(country)
}

func (s *superstarServicee) Search(country string) []models.StarInfo {
	return s.dao.Search(country)
}

func (s *superstarService)Get(id int) *models.StarInfo {
	return s.dao.Get(id)
}

func (s *superstarServicee) Get(id int) *models.StarInfo {
	return s.dao.Get(id)
}

func (s *superstarService)Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *superstarServicee) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *superstarService)Update(user *models.StarInfo, columns []string) error {
	return s.dao.Update(user, columns)
}

func (s *superstarServicee) Update(user *models.StarInfo, columns []string) error {
	return s.dao.Update(user, columns)
}

func (s *superstarService)Create(user *models.StarInfo) error {
	return s.dao.Create(user)
}

func (s *superstarServicee) Create(user *models.StarInfo) error {
	return s.dao.Create(user)
}
