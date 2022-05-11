package main

import (
	"context"
	"fmt"

	productservice "github.com/evgeniy-dammer/serverinterceptorsgrpc/productservice/proto"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:1111", grpc.WithInsecure())

	if err != nil {
		fmt.Println(err)
	}

	defer connection.Close()

	productServ := productservice.NewProductServiceClient(connection)

	response1, err := productServ.FindAll(context.Background(), &productservice.FindAllRequest{})

	if err != nil {
		fmt.Println(err)
	} else {
		products := response1.Products

		fmt.Println("Product List")

		for _, product := range products {

			fmt.Println("Id: ", product.Id)
			fmt.Println("Name: ", product.Name)
			fmt.Println("Price: ", product.Price)
			fmt.Println("Quantity: ", product.Quantity)
			fmt.Println("Status: ", product.Status)
			fmt.Println("")
		}

		fmt.Println("========================")
	}

	response2, err := productServ.Search(context.Background(), &productservice.SearchRequest{Keyword: "vi"})

	if err != nil {
		fmt.Println(err)
	} else {
		products := response2.Products

		fmt.Println("Search Result")

		for _, product := range products {

			fmt.Println("Id: ", product.Id)
			fmt.Println("Name: ", product.Name)
			fmt.Println("Price: ", product.Price)
			fmt.Println("Quantity: ", product.Quantity)
			fmt.Println("Status: ", product.Status)
			fmt.Println("")
		}

		fmt.Println("========================")
	}
}
