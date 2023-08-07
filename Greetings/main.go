package main

import (
	"fmt"
	"net"
	"net/http"

	"greetings/asyncgreet"
	"greetings/datatypes"
	"greetings/embedding"
	"greetings/interfaceFactory"
	"greetings/selectCase"
	"greetings/slices"
	"greetings/switchCase"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/all", all)

	go bareBonesHttpServer(mux)

	addr := ":8080"
	fmt.Printf("Server listening on: %s \n", addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func bareBonesHttpServer(mux *http.ServeMux) {
	s := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	ln, e := net.Listen("tcp", s.Addr)
	if e != nil {
		fmt.Printf("%v", e)
	}

	fmt.Printf("Server listening on: %s \n", s.Addr)
	s.Serve(ln)
}

func all(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
	asyncgreet.Greet()
	datatypes.Types()
	datatypes.ComplexTypes()
	slices.SliceOptn()
	switchCase.SwitchCase()
	selectCase.GetFibonnaci()
	selectCase.GetGoPath()
	interfaceFactory.GetSound([]string{"Cat", "Dog", "Horse"})
	embedding.Embed()
}
