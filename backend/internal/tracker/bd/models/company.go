package models

type records map[int]string

func AllCompany() records {
	mapCompany := make(map[int]string)
	mapCompany[1] = "Ретн"
	mapCompany[2] = "Смарт Сервис"
	return mapCompany
}

func Company(id int) records {
	mapCompany := make(map[int]string)
	mapCompany[7] = "Корнеаль"
	return mapCompany
}
