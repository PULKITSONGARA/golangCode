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
	// var map_name string
	// var key string
	fmt.Printf("hello, world\n")
	fmt.Println(pow(2, 4, 10))

	ah := accountHolder{}
	_, err := fmt.Scan(&ah.account_no, &ah.mobile_no, &ah.account_bal)
	ah.name, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	ah.address, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	// _, err := fmt.Scan(&ah.mobile_no)
	// _, err := fmt.Scan(&ah.account_bal)
	state_population := map[string]int{
		"Madhya Pradesh": 320982,
		"Karnataka":      280298,
		"Uttar Pradesh":  230982,
		"Tamilnadu":      298024,
		"Maharashtra":    328789,
		"Chhattisgrah":   980982,
	}
	del_map_entry(state_population, "Karnataka")
	fmt.Println(state_population, err)
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

func del_map_entry(map_name map[string]int, key string) {
	delete(map_name, key)

}
