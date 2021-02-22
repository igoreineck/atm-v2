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

func ListDir(path string) []string {
	listedAccounts := []string{}

	directories, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, directory := range directories {
		if directory.IsDir() {
			directoryName := path + directory.Name()
			listedAccounts = append(listedAccounts, directoryName)
		}
	}

	return listedAccounts
}
