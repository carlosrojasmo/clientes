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
	address  = "10.10.28.10:50051"
)

var wg = &sync.WaitGroup{}


//Funcion para leer las ordenes de un archivo pymes.csv
func ordenesDePymes(period time.Duration){
	defer wg.Done()
	pymes, err := os.Open("pymes.csv") //Abrimos el archivo
	if err != nil{
		log.Fatalln("Couldn't open the csv files", err)
	}
    rp := csv.NewReader(pymes)
    rp.Read()
    for { //En cada iteracion leemos una orden
		ordenp, err := rp.Read()
		if err == io.EOF{
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		wg.Add(1)
		go clienteDePyme(ordenp,3) //Esta funcion se encarga de la comunicacion con logistica

		time.Sleep(period * time.Second) //Tiempo de espera antes de enviar la siguiente orden
		
	}

}

//Funcion que envia la informacion de la orden a logistica.
/*Posteriormente entra en un loop preguntando periodicamente
a logistica por el estado del paquete*/
func clienteDePyme(ordenp []string,period time.Duration){
	defer wg.Done()
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrdenServiceClient(conn)

	tip := "normal" //Por defecto suponemos que es tipo normal

	if ordenp[5] == "1" {
		tip = "prioritario" //Revisando los datos cambiamos a prioritario si corresponde
	}

	val,err := strconv.ParseInt(ordenp[2], 10, 64)

	if err != nil {
		fmt.Printf("no")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ReplyToOrder(ctx, &pb.SendToOrden{IdPaquete : ordenp[0],Tipo : tip ,Nombre :  ordenp[1],
		Valor : val,Origen : ordenp[3],Destino : ordenp[4]}) //Enviamos la orden y recibimos el codigo de seguimiento.
		fmt.Println(err)
	if err != nil {
		fmt.Println("No se pudo enviar la orden ",ordenp[0]," del tipo ",tip)
		fmt.Println("could not greet: ", err)
	} else {
		seg := r.GetSeguimiento() //Guardamos el codigo de seguimiento en seg.

		for {//Periodicamente cliente pregunta el estado de su paquete ,termina cuando este es Recibido o No Recibido.
			time.Sleep(period * time.Second)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			r, err := c.GetState(ctx, &pb.ReplyFromOrden{Seguimiento : seg})
			fmt.Println(err)
			if err != nil {
				fmt.Println("Fallo la consulta de seguimiento de: ",ordenp[0])
			} else {
				est := r.GetEstado()
				if est == "Recibido" || est == "No Recibido" {
                    fmt.Println("El paquete ",ordenp[0]," fue ",est)
					break
				}
	    	}
		}
	}
}

//Funcion para leer las ordenes de un archivo retail.csv
func ordenesDeRetail(period time.Duration){
	defer wg.Done()
	retail, err := os.Open("retail.csv")//Abrimos el archivo
	if err != nil {
		log.Fatalln("Couldn't open the csv files", err)
	}
	rr := csv.NewReader(retail)
	rr.Read()

	for {//En cada iteracion leemos una orden
		ordenr, err := rr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		
	    wg.Add(1)
		go clienteDeRetail(ordenr,3)//Esta funcion se encarga de la comunicacion con logistica
		
		time.Sleep(period * time.Second)//Tiempo de espera antes de enviar la siguiente orden
		
	}

}

//Funcion que envia la informacion de la orden a logistica.
/*Posteriormente entra en un loop preguntando periodicamente
a logistica por el estado del paquete*/
func clienteDeRetail(ordenr []string,period time.Duration){
	defer wg.Done()
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrdenServiceClient(conn)

	

	val,err := strconv.ParseInt(ordenr[2], 10, 64)

	if err != nil {
		fmt.Printf("no")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ReplyToOrder(ctx, &pb.SendToOrden{IdPaquete : ordenr[0],Tipo : "retail",Nombre :  ordenr[1],
		Valor : val,Origen : ordenr[3],Destino : ordenr[4]}) //Enviamos la orden y recibimos el codigo de seguimiento.
	fmt.Println(err)
	if err != nil {
		fmt.Println("No se pudo enviar la orden ",ordenr[0]," del tipo retail")
		fmt.Println("could not greet: ", err)
	} else {
		seg := r.GetSeguimiento() //Guardamos el codigo de seguimiento en seg.
		for {//Periodicamente cliente pregunta el estado de su paquete ,termina cuando este es Recibido o No Recibido.
			time.Sleep(period * time.Second)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			r, err := c.GetState(ctx, &pb.ReplyFromOrden{Seguimiento : seg})
			fmt.Println(err)
			if err != nil {
				fmt.Println("Fallo la consulta de seguimiento de: ",ordenr[0])
			} else {
				fmt.Println("getstate funciono")
				est := r.GetEstado()
				fmt.Println(est)
				if est == "Recibido" || est == "No Recibido" {
					fmt.Println("El paquete ",ordenr[0]," fue ",est)
					break
				}
	    	}
		}
	}
}
func main() {
	
	// Se inicia como una rutina los pedidos de las ordenes de retail
	wg.Add(1)
	go ordenesDeRetail(12) 

	time.Sleep(6 * time.Second)

    // Se inicia como una rutina los pedidos de las ordenes de pymes
	wg.Add(1)
	go ordenesDePymes(12)
	wg.Wait()
	fmt.Println("Fin de la ejecuccion de clientes.go")
}