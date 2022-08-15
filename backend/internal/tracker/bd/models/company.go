package models

type Records map[int]string

func AllCompany() Records {
	mapCompany := make(map[int]string)
	mapCompany[1] = "Ретн"
	mapCompany[2] = "Смарт Сервис"
	return mapCompany
}

func Company(id int) Records {
	mapCompany := make(map[int]string)
	mapCompany[7] = "Корнеаль"
	return mapCompany
}
