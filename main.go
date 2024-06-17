package main
import "fmt"

func main(){

	menu := map[string]float64{
		"Pizza": 40.00,
		"Suco": 12.50,
		"X-tudo": 22.90,
	}

	for k,v := range menu{
		fmt.Printf("%s custa R$%.2f\n", k, v)
	}

contanova := novaConta("Astrubal")
fmt.Println(contanova)

contanova2 := novaConta("Jubileu")
fmt.Println(contanova2)
}