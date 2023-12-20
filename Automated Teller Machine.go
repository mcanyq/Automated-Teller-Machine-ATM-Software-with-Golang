package main
import "fmt"

func main(){

	balance:= 10000

	for true{

fmt.Println(` 

********TRANSACTİONS********

1: WITHDRAW MONEY     
2: DEPOSIT            
3: BALANCE INQUIRY    
4: EXIT               

****************************

PLEASE SELECT THE TRANSACTION TYPE...

`)

var process int
fmt.Scanln(&process)

if process==1{
	fmt.Println("ENTER THE AMOUNT TO WITHDRAW: ")

	var withdrawn int
	fmt.Scanln(&withdrawn)

	if withdrawn<=balance{
		fmt.Println("YOUR WITHDRAWAL TRANSACTION HAS BEEN COMPLETED SUCCESSFULLY...")
		balance-=withdrawn
		fmt.Println("YOUR CURRENT BALANCE:", balance)
	}else{
		fmt.Println("YOUR BALANCE IS NOT ENOUGH...")

	}

}else if process==2{

	fmt.Println("ENTER THE AMOUNT YOU WANT TO DEPOSIT: ")

	var deposit int
    fmt.Scanln(&deposit)

	balance+=deposit

	fmt.Println("YOUR CURRENT BALANCE:", balance)


        }else if process==3{
		fmt.Println("YOUR CURRENT BALANCE:", balance)
		
	
	}else if process==4{

		fmt.Println("WE WISH YOU A NICE DAY...")

break

	}else{
		fmt.Println("PLEASE SELECT AN APPROPRIATE ACTION...")
	}

	}
}
