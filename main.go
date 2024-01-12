/*
Senaryo, bankayadan hesap oluşturulacak. Vadeli hesap açılacak. Bankaya para yatırılacak. Bankadan Para çekilecek
*/
package main

import "fmt"

func main() {
	runBank()
}

func runBank() {
	fmt.Println("Bankamıza Hoş Geldiniz.")
	account := createABankAccount()
	termAccount := createATermAccount(account)
	fmt.Printf("%d hesabınıza, %d hesabınız açılmıştır \n", account, termAccount)
	deposit := depositMoneyInTheBank(termAccount)
	fmt.Printf("Hesabınıza Yatırılan Para : %d \n", deposit)
	withdrawnMoney, remainingMoney := withDrawMoneyFromTheBank(termAccount, deposit)
	fmt.Printf("Hesabınızda Çekilen Para : %d \nKalan Para : %d \n", withdrawnMoney, remainingMoney)
	fmt.Println("Tekrar bekleriz iyi günler.")

}

func createABankAccount() int {
	var account int
	fmt.Print("Lütfen numara giriniz : ")
	fmt.Scan(&account)
	return account
}
func createATermAccount(account int) int {
	var termAccount int
	fmt.Printf("%d hesabınıza, vadeli hesap için 1, vadesiz hesap oluşturmak için 2 yazınız ", account)
	fmt.Scan(&termAccount)
	return termAccount
}
func depositMoneyInTheBank(termAccount int) int {
	var deposit int
	fmt.Printf("%d hesabınıza yatırılacak parayı yazınız : ", termAccount)
	fmt.Scan(&deposit)
	return deposit
}
func withDrawMoneyFromTheBank(termAccount, accountDeposit int) (int, int) {
	var drawMoney int
	fmt.Printf("%d Hesabınızdan çekilecek parayı yazınız : ", termAccount)
	fmt.Scan(&drawMoney)
	calc := accountDeposit - drawMoney
	return drawMoney, calc
}
