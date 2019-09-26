package main

func main()  {

	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}

func sendData()  {
	
}
