package main


import (
	"fmt"
	"io"
	"net/http"
)


func main() {
	//Greet(os.Stdout, "Elodie")
	http.ListenAndServe(":5000",http.HandlerFunc(MyGreetHandler))
}
func MyGreetHandler(w http.ResponseWriter,r *http.Request) {
	Greet(w, "wssssssord")
}

func Greet(writer io.Writer, name string){
	fmt.Fprintf(writer,"Hello,%s",name)
}
