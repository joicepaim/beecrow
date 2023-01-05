package main

import "fmt"

func main() {
	var H, M, O, N, i int64
	i = 0
	fmt.Scanf("%d\n", &N)
	for i < N {
		fmt.Scanf("%d %d %d\n", &H, &M, &O)
		if H < 10 {
			fmt.Print("0", H)
		} else {
			fmt.Print(H)
		}
		if M < 10 {
			fmt.Printf(":0%d - ", M)
		} else {
			fmt.Printf(":%d - ", M)
		}
		if O > 0 {
			fmt.Println("A porta abriu!")
		} else {
			fmt.Println("A porta fechou!")
		}
		i++
	}

}
