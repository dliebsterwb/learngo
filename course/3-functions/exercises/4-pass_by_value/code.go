package main

import "fmt"

func main() {
	sendsSoFar := 430
	const sendsToAdd = 25
	incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages")
	fmt.Println("Betch were looking for", incrementSends(sendsSoFar, sendsToAdd))
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("Much Better!! ", sendsSoFar)
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	return sendsSoFar + sendsToAdd
}
