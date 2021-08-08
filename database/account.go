package database

import "coursesheduling/model"

func GetRoleByWord(queryWord string) (roles []string) {
	rows, _ := appDB.Raw("select role from ?", AccountTable).Rows()
	for rows.Next(){
		var field string
		rows.Scan(&field)
		roles = append(roles,field)
	}
	return roles
}

func GetAccounts() (accounts []model.Account) {
	appDB.Find(&accounts)
	return accounts
}

func GetAccountByWord(queryWord string) (accounts []model.Account) {
	appDB.Where("user_name LIKE ?","%"+queryWord+"%").Find(&accounts)
	return accounts
}

func GetAccountByName(queryWord string) (account model.Account) {
	appDB.Where("user_name = ?",queryWord).Find(&account)
	return account
}
