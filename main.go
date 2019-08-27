package main

import (
	"Queue/conf"
	"Queue/connection"
	"fmt"
	"os"
	queueMemroy "queue-memory/conf"
	queue "service-queue"
)

func main() {
	var opt int
	var number int
	var svc queue.Service
	config := "redis"
	switch config {
	case "redis":
		connect := connection.NewRedis()
		svc = conf.New(connect)

	// case "mgodb":
	// 	var ctx, db = mgodb.Connect()

	case "memory":
		svc = queueMemroy.New()
	}

	for {
		fmt.Println("=========================")
		fmt.Println("          QUEUE          ")
		fmt.Println("=========================")
		fmt.Printf("1. Push\n2. Pop\n3. PopAll\n4. Key\n5. Len\n6. Exit\nChoose: ")
		fmt.Scanln(&opt)

		switch opt {
		case 1:
			fmt.Print("Enter your data: ")
			fmt.Scanln(&number)
			svc.Push(number)
			svc.Contains(number)
		case 2:
			data := svc.Pop()
			fmt.Println("Remove :", data)
		case 3:
			fmt.Print("PopAll :")
			for _, each := range svc.Keys() {
				fmt.Print(each)
				svc.Pop()
			}
			fmt.Println()
		case 4:
			key := svc.Keys()
			fmt.Println("Key :", key)
		case 5:
			len := svc.Len()
			fmt.Println("Count: ", len)
		case 6:

			os.Exit(1)
		default:
			fmt.Println("There's no option")
		}
	}

}
