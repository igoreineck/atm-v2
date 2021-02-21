package account

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

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

func generateUuid() []byte {
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}

	return uuid
}

func getAccountInstance(uuid []byte, name string, cpf string, age uint8, password string) Account {
	return Account{
		Id:       uuid,
		Name:     name,
		Cpf:      cpf,
		Age:      age,
		Password: password,
	}
}

func getAccountPath(uuid []byte) string {
	_, b, _, _ := runtime.Caller(0)
	srcDir := filepath.Dir(b)
	packageDir := filepath.Dir(srcDir)
	rootDir := filepath.Dir(packageDir)
	return rootDir + "/data/accounts/" + string(uuid)
}

func createAccountDir(accountPath string) {
	err := os.Mkdir(accountPath, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func createAccountFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
}

func writeAccountInfo(file string, accountInfo Account) {
	acc, err := json.MarshalIndent(accountInfo, "", "")
	if err != nil {
		log.Fatal(err)
	}

	_ = ioutil.WriteFile(file, acc, 0644)
}

// byteValue, _ := ioutil.ReadAll(jsonFile)
