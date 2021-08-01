package models

import "lesson04/dao"

type Menu struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func (Menu)TableName() string {
	return "Menu"
}

func InsertData(menu Menu)(err error)  {
	err = dao.MyDB.Debug().Create(&menu).Error
	return err
}

func FindAllData(menuList *[] Menu)(err error)  {
	err = dao.MyDB.Debug().Find(menuList).Error
	return err
}

func FindOneData(id string, menu *Menu)(err error)  {
	err = dao.MyDB.Debug().Where("id=?",id).First(menu).Error
	return err
}

func UpdateData(menu *Menu)(err error)  {
	err = dao.MyDB.Debug().Save(menu).Error
	return err
}

func DeleteDataById(id string)(err error) {
	var menu Menu
	//Delete(caidan)指定表名的作用
	err = dao.MyDB.Debug().Where("id=?",id).Delete(&menu).Error
	return err
}