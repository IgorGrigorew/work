package server

import (
	"fmt"
	"net/http"
	"strconv"
)

const port string = ":5000"

var result []int

func Server() {

	fmt.Println("Start server")

	http.HandleFunc("/", handler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}

}


func handler(w http.ResponseWriter, r *http.Request) {

	

	switch r.Method {
	case "GET":
		
		numString := r.URL.Query().Get("num")
		num, err := strconv.Atoi(numString)
		if err != nil {
			http.Error(w, "ошибка конвертации", http.StatusBadRequest)
			return
		}
		result = append(result, num)
	
		fmt.Fprintf(w, "%d добавленна в %v", num, result)
	
		

	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
	
}