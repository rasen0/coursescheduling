package database

import "coursesheduling/model"

func GetRoleByWord(queryWord string) (roles []model.RoleItem) {
	appDB.Where("role LIKE %?%",queryWord).Find(&roles)
	return roles
}

func GetRole() (roles []string) {
	rows, _ := appDB.Raw("select role from ?", RoleTable).Rows()
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

func UpdateAccount(account model.Account) {
	appDB.Model(&model.Account{}).Where("user_name = ?",account.UserName).Updates(account)
}

func InsertAccount(account model.Account) {
	appDB.Create(account)
}

func InsertAccounts(accounts []model.Account) {
	appDB.Create(accounts)
}