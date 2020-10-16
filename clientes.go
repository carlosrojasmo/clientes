package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

func ordenesDePymes(period time.Duration){
	defer wg.Done()
	pymes, err := os.Open("pymes.csv")
	if err != nil{
		log.Fatalln("Couldn't open the csv files", err)
	}
    rp := csv.NewReader(pymes)
    rp.Read()
    for {
		ordenp, err := rp.Read()
		if err == io.EOF{
			break
		}
		if err != nil {
			log.Fatal(err)
		}

        wg.Add(1)
		go clienteDePyme(ordenp,0)

		time.Sleep(period * time.Second)
		
	}

}

func clienteDePyme(ordenp []string,period time.Duration){
	defer wg.Done()
	fmt.Printf("id: %s Producto %s Valor %s Tienda %s Destino %s Prioritario %s\n ", ordenp[0], ordenp[1], ordenp[2], 
			ordenp[3], ordenp[4], ordenp[5])
	for {
		time.Sleep(period * time.Second)
		if true {
			break
		}
	}
}

func ordenesDeRetail(period time.Duration){
	defer wg.Done()
	retail, err := os.Open("retail.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv files", err)
	}
	rr := csv.NewReader(retail)
	rr.Read()

	for {
		ordenr, err := rr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

	    wg.Add(1)
		go clienteDeRetail(ordenr,0)
		
		time.Sleep(period * time.Second)
		
	}

}

func clienteDeRetail(ordenr []string,period time.Duration){
	defer wg.Done()
	fmt.Printf("id: %s Producto %s Valor %s Tienda %s Destino %s\n ", ordenr[0], ordenr[1], ordenr[2], 
			ordenr[3], ordenr[4])
	for {
		time.Sleep(period * time.Second)
		if true {
			break
		}
	}
}

func main() {
	
	wg.Add(1)
	go ordenesDeRetail(4)

	time.Sleep(2 * time.Second)

	wg.Add(1)
	go ordenesDePymes(4)

	wg.Wait()
	
}