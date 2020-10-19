package main

import (
    "context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
	"strconv"
	"google.golang.org/grpc"
	pb "../proto"
)

const (
	address  = "localhost:50051"
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
		fmt.Println("pyme")
		go clienteDePyme(ordenp,3)

		time.Sleep(period * time.Second)
		
	}

}

func clienteDePyme(ordenp []string,period time.Duration){
	defer wg.Done()
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrdenServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tip := "normal" 

	if ordenp[5] == "1" {
		tip = "prioritario"
	}

	val,err := strconv.ParseInt(ordenp[2], 10, 64)

	if err != nil {
		fmt.Printf("no")
	}

	r, err := c.ReplyToOrder(ctx, &pb.SendToOrden{IdPaquete : ordenp[0],Tipo : tip ,Nombre :  ordenp[1],
		Valor : val,Origen : ordenp[3],Destino : ordenp[4]})
		fmt.Println(err)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	seg := r.GetSeguimiento() 

	
	//fmt.Printf("id: %s Producto %s Valor %s Tienda %s Destino %s Prioritario %s\n ", ordenp[0], ordenp[1], ordenp[2], 
	//		ordenp[3], ordenp[4], ordenp[5])
	for {

		time.Sleep(period * time.Second)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()

		r, err := c.GetState(ctx, &pb.ReplyFromOrden{Seguimiento : seg})
		fmt.Println(err)
        if err != nil {
        	fmt.Println("Entro en ERROR")
	    } else {
	    	fmt.Println("getstate funciono")
	    	est := r.GetEstado()
	    	fmt.Println(est)
	    	if est == "Recibido" || est == "No Recibido" {
	    		break
	    	}
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
		fmt.Println("retail")
	    wg.Add(1)
		go clienteDeRetail(ordenr,3)
		
		time.Sleep(period * time.Second)
		
	}

}

func clienteDeRetail(ordenr []string,period time.Duration){
	defer wg.Done()
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrdenServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	val,err := strconv.ParseInt(ordenr[2], 10, 64)

	if err != nil {
		fmt.Printf("no")
	}

	r, err := c.ReplyToOrder(ctx, &pb.SendToOrden{IdPaquete : ordenr[0],Tipo : "retail",Nombre :  ordenr[1],
		Valor : val,Origen : ordenr[3],Destino : ordenr[4]})
	fmt.Println(err)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	seg := r.GetSeguimiento() 

	//fmt.Printf("id: %s Producto %s Valor %s Tienda %s Destino %s\n ", ordenr[0], ordenr[1], ordenr[2], 
	//		ordenr[3], ordenr[4])
	for {
		time.Sleep(period * time.Second)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()
        
		r, err := c.GetState(ctx, &pb.ReplyFromOrden{Seguimiento : seg})
		fmt.Println(err)
        if err != nil {
        	fmt.Println("Entro en ERROR")
	    } else {
	    	fmt.Println("getstate funciono")
	    	est := r.GetEstado()
	    	fmt.Println(est)
	    	if est == "Recibido" || est == "No Recibido" {
	    		break
	    	}
	    }
	}
}

func main() {
	
	wg.Add(1)
	go ordenesDeRetail(12)

	time.Sleep(6 * time.Second)

	wg.Add(1)
	go ordenesDePymes(12)
	wg.Wait()
	
}