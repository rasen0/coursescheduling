package dao

import (
	"coursesheduling/database"
	"coursesheduling/model"
)

func QueryRoleByWord(queryWord string) (roles []string) {
	roles = database.GetRoleByWord(queryWord)
	return roles
}

func QueryAccounts() (accounts []model.Account) {
	accounts = database.GetAccounts()
	return accounts
}

func QueryAccountByWord(queryWord string) (accounts []model.Account) {
	accounts = database.GetAccountByWord(queryWord)
	return accounts
}

func QueryAccountByName(name string) (account model.Account){
	account = database.GetAccountByName(name)
	return account
}

func AddAccounts(accounts []model.Account) {
	database.InsertAccounts(accounts)
}

func AddAccount(account model.Account) {
	database.InsertAccount(account)
}

func UpdateAccount(account model.Account){
	database.UpdateAccount(account)
}