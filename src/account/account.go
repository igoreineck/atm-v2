package account

type Account struct {
	Id       []byte  `json:"id"`
	Name     string  `json:"name"`
	Cpf      string  `json:"cpf"`
	Password string  `json:"password"`
	Age      uint8   `json:"age"`
	Balance  float64 `json:"balance"`
}

func CreateAccount(name string, cpf string, age uint8, password string) {
	uuid := generateUuid()
	account := getAccountInstance(uuid, name, cpf, age, password)
	accountDir := getAccountPath(uuid)
	fileDir := accountDir + "/info.json"

	createAccountDir(accountDir)
	createAccountFile(fileDir)
	writeAccountInfo(fileDir, account)
}
