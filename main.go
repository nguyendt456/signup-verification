package main

import "github.com/nguyendt456/signup-with-verification/deploy"

func main() {
	go deploy.StartGW()
	deploy.StartgRPCservice()
}
