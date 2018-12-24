// package main

// import (
// 	"fmt"
// 	"math"
// )

// // type accountHolder struct {
// // 	account_no  int
// // 	name        string
// // 	address     string
// // 	mobile_no   int64
// // 	account_bal float64
// // }
// type Writer interface {
// 	Write([]byte) (int, error)
// }

// type ConsoleWriter struct{}

// func (cw ConsoleWriter) Write(data []byte) (int, error) {
// 	n, err := fmt.Println(string(data))
// 	return n, err
// }

// type IntCounter int

// func (ic *IntCounter) increment() int {
// 	*ic++
// 	return int(*ic)
// }

// type Incrementer interface {
// 	Increment() int
// }

// func main() {

// 	var w Writer = ConsoleWriter{}
// 	w.Write([]byte("Hello Go"))

// 	myInt := IntCounter(0)
// 	var inc Incrementer = &myInt
// 	for i := 0; i < 11; i++ {
// 		fmt.Println(inc.Increment())
// 	}

// 	fmt.Printf("hello, world\n")
// 	fmt.Println(pow(2, 4, 10))

// 	// ah := accountHolder{}
// 	// _, err := fmt.Scan(&ah.account_no, &ah.mobile_no, &ah.account_bal)
// 	// ah.name, _ = bufio.NewReader(os.Stdin).ReadString('\n')
// 	// ah.address, _ = bufio.NewReader(os.Stdin).ReadString('\n')
// 	// _, err := fmt.Scan(&ah.mobile_no)
// 	// _, err := fmt.Scan(&ah.account_bal)
// 	state_population := map[string]int{
// 		"Madhya Pradesh": 320982,
// 		"Karnataka":      280298,
// 		"Uttar Pradesh":  230982,
// 		"Tamilnadu":      298024,
// 		"Maharashtra":    328789,
// 		"Chhattisgrah":   980982,
// 	}
// 	fmt.Println(state_population)
// 	delete(state_population, "Karnataka")
// 	fmt.Println(state_population)
// }

// func pow(x, n, lim float64) float64 {
// 	if a := math.Pow(x, n); a < lim {
// 		return a
// 	}
// 	return lim
// }

// func interest_cals(account_bal float64) (float64, string) {

// 	if account_bal > 0 {
// 		interest_bal := account_bal * .10
// 		account_bal = account_bal + interest_bal
// 		return account_bal, "0"
// 	}
// 	return 0, "Your account balance is zero"
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}
// var count = 0
// var m = sync.RWMutex{}

// func main() {

// 	for i := 0; i < 10; i++ {
// 		wg.Add(2)
// 		go sayHello()
// 		go printCount()
// 		wg.Wait()
// 	}

// }

// func sayHello() {
// 	m.RLock()
// 	fmt.Println("Hello")
// 	m.RUnlock()
// 	wg.Done()
// }

// func printCount() {
// 	m.Lock()
// 	count++
// 	fmt.Println(count)
// 	m.Unlock()
// 	wg.Done()
// }

package main

import (
	"fmt"
	"sync"
)

var count = 0
var wg = sync.WaitGroup{}

func main() {
	// router := mux.NewRouter()
	// router.HandleFunc("/people", getPeople).Methods("GET")
	ch := make(chan int)
	wg.Add(2)
	go func() {
		for i := range ch {
			// j := <-ch
			fmt.Println(i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()

	wg.Wait()
}

// func getPeople() string {
// 	return "done api call"
// }
