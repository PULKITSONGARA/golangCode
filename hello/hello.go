package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type accountHolder struct {
	account_no  int
	name        string
	address     string
	mobile_no   int64
	account_bal float64
}

func main() {
	fmt.Printf("hello, world\n")
	// fmt.Println(pow(2, 4, 10))
	// reader := bufio.NewReader(os.Stdin)
	ah := accountHolder{}
	// var x int
	_, err := fmt.Scan(&ah.account_no, &ah.mobile_no, &ah.account_bal)
	ah.name, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	ah.address, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	// _, err := fmt.Scan(&ah.mobile_no)
	// _, err := fmt.Scan(&ah.account_bal)
}

func pow(x, n, lim float64) float64 {
	if a := math.Pow(x, n); a < lim {
		return a
	}
	return lim
}

func interest_cals(account_bal float64) (float64, string) {

	if account_bal > 0 {
		interest_bal := account_bal * .10
		account_bal = account_bal + interest_bal
		return account_bal, "0"
	}
	return 0, "Your account balance is zero"
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	//reading an integer
// 	var age int
// 	fmt.Println("What is your age?")
// 	_, err := fmt.Scan(&age)

// 	//reading a string
// 	reader := bufio.NewReader(os.Stdin)
// 	var name string
// 	fmt.Println("What is your name?")
// 	name, _ = reader.ReadString('\n')

// 	fmt.Println("Your name is ", name, " and you are age ", age)
// 	fmt.Println(err)
// }
