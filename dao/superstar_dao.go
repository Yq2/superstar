package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/yz124/superstar/models"
)

type SuperstarDao struct {
	engine *xorm.Engine
}

type SuperstarDaoo struct {
	engine *xorm.Engine
}
//dao
func NewSuperstarDao(engine *xorm.Engine) *SuperstarDao {
	return &SuperstarDao {
		engine:engine,
	}
}

func NewSuperstarDaoo(engin *xorm.Engine) *SuperstarDaoo {
	return &SuperstarDaoo{
		engine:engin,
	}
}

func (d *SuperstarDao) Get(id int) *models.StarInfo {
	data := &models.StarInfo{Id:id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *SuperstarDaoo) Get(id int) *models.StarInfo {
	data := &models.StarInfo{Id:id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *SuperstarDao) GetAll() []models.StarInfo {
	datalist := make([]models.StarInfo, 0)
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *SuperstarDaoo) GetAll() []models.StarInfo {
	dataList := make([]models.StarInfo,0)
	err := d.engine.Desc("id").Find(&dataList)
	if err != nil {
		return dataList
	} else {
		return []models.StarInfo{}
	}
}

func (d *SuperstarDao) Search(country string) []models.StarInfo {
	datalist := make([]models.StarInfo, 0)
	err := d.engine.Where("country=?", country).
		Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *SuperstarDaoo) Search(country string) []models.StarInfo {
	dataList := make([]models.StarInfo, 0)
	err := d.engine.Where("country=?", country).Desc("id").Find(&dataList)
	if err != nil {
		return dataList
	} else {
		return []models.StarInfo{}
	}
}

func (d *SuperstarDao) Delete(id int) error {
	data := &models.StarInfo{Id:id, SysStatus:1}
	//根据ID选择记录后更新
	//逻辑删除
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *SuperstarDaoo) Delete(id int) error {
	data := &models.StarInfo{Id:id,SysStatus:1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *SuperstarDao) Update(data *models.StarInfo, columns []string) error {
	//MustCols强制更新某些列，因为xorm遇到某些值为空时是不会更新的
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *SuperstarDaoo) Update(data *models.StarInfo, colunms []string) error {
	_, err := d.engine.Id(data.Id).MustCols(colunms...).Update(data)
	return err
}

func (d *SuperstarDao) Create(data *models.StarInfo) error {
	//插入一条记录
	_, err := d.engine.Insert(data)
	return err
}

func (d *SuperstarDaoo) Create(data *models.StarInfo) error {
	//插入一条记录
	_, err := d.engine.Insert(data)
	return err
}