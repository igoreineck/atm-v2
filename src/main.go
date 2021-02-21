package main

import "github.com/igoreineck/atm-v2/src/account"

func main() {
	// printf("DIGITE UMA DAS OPÇÕES ABAIXO: \n");
	// printf("-------------------------------------------------------\n");
	// printf("Digite 1 - Para criar uma conta.\n");
	// printf("Digite 2 - Para efetuar o login no sistema.\n");
	// printf("Digite 3 - Para acessar como Admin.\n");
	// printf("Digite 4 - Para Sair do sistema.\n");
	// printf("-------------------------------------------------------\n");

	account.CreateAccount("Igor", "1234567", 12, "123")
}
