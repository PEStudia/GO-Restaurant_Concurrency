package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	reset  = "\033[0m"
)

func main() {
	// Parametry symulacji restauracji
	numCustomers := 5
	numTables := 4
	numChefs := 2
	numWaiters := 1

	// Kanały do komunikacji
	// tables - Kanał reprezentujący dostępne stoliki
	// orders - Kanał do składania zamówień
	// preparedOrders - Kanał do gotowych zamówień
	// deliveredOrders - Kanał do dostarczonych zamówień
	tables := make(chan int, numTables)
	orders := make(chan int, numCustomers)
	preparedOrders := make(chan int, numCustomers)
	deliveredOrders := make(chan int, numCustomers)

	// Inicjalizacja stolików
	for i := 1; i <= numTables; i++ {
		tables <- i
	}

	// WaitGroup do oczekiwania na wyjście wszystkich klientów
	var wg sync.WaitGroup
	wg.Add(numCustomers)

	// Symulacja pracy kucharzy przygotowujących zamówienia
	for i := 1; i <= numChefs; i++ {
		go chef(i, orders, preparedOrders)
	}

	// Symulacja pracy kelnerów dostarczających zamówienia
	for i := 1; i <= numWaiters; i++ {
		go waiter(i, preparedOrders, deliveredOrders)
	}

	// Symulacja klientów wchodzących do restauracji
	for i := 1; i <= numCustomers; i++ {
		go customer(i, tables, orders, deliveredOrders, &wg)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second) // Losowe opóźnienie wejścia klienta
	}

	// Oczekiwanie na wyjście wszystkich klientów
	wg.Wait()
	fmt.Println("\n**** Restauracja jest zamykana. Wszyscy klienci opuścili lokal. ****")
}

// Symuluje klienta w restauracji
func customer(id int, tables chan int, orders chan int, deliveredOrders chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf(blue+"[-->] Klient %d wchodzi do restauracji\n"+reset, id)

	// Oczekiwanie na stolik
	fmt.Printf(yellow+"[-_-] Klient %d czeka na stolik\n"+reset, id)
	table := <-tables
	fmt.Printf(yellow+"[o_o] Klient %d siada przy stoliku %d\n"+reset, id, table)

	// Składanie zamówienia
	fmt.Printf(yellow+"[:/] Klient %d składa zamówienie\n"+reset, id)
	orders <- id

	// Oczekiwanie na dostarczenie jedzenia
	fmt.Printf(yellow+"[z_z] Klient %d czeka na jedzenie\n"+reset, id)
	<-deliveredOrders // Oczekiwanie na dostarczenie jedzenia

	// Symulacja jedzenia
	time.Sleep(time.Second * 2)
	fmt.Printf(green+"[:D] Klient %d skończył jeść\n"+reset, id)

	// Opuszczenie stolika
	fmt.Printf(blue+"[<--] Klient %d opuszcza stolik %d\n"+reset, id, table)
	tables <- table
}

// Symuluje kucharza przygotowującego zamówienie
func chef(id int, orders chan int, preparedOrders chan int) {
	for {
		order := <-orders
		fmt.Printf("Kucharz %d przygotowuje jedzenie dla klienta %d\n", id, order)
		time.Sleep(time.Second * 3) // Symulacja czasu przygotowywania
		fmt.Printf("Kucharz %d zakończył przygotowywanie jedzenia dla klienta %d\n", id, order)
		preparedOrders <- order // Powiadomienie, że zamówienie jest gotowe
	}
}

// Symuluje kelnera dostarczającego zamówienie
func waiter(id int, preparedOrders chan int, deliveredOrders chan int) {
	for {
		order := <-preparedOrders
		fmt.Printf("Kelner %d dostarcza jedzenie klientowi %d\n", id, order)
		deliveredOrders <- order // Powiadomienie, że jedzenie zostało dostarczone
	}
}
