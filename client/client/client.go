package client

import (
	"fmt"
	"io"
	
	"net/http"
	"net/url"
)

func Client() {

	ip, port := input()

	adres := url.URL{
		Scheme: "http",
		Host: fmt.Sprintf("%s:%s", ip, port),
	}


for {

	var num string

	fmt.Println("введите число")

	fmt.Scan(&num)
//строка запроса 
	adres.RawQuery = url.Values{"num": {num}}.Encode()

// отправляем запрос
resp, err := http.Get(adres.String())
	if err != nil {
		fmt.Println("ошибка запроса", err)
		break
	}
 

 body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Ошибка чтения ответа:", err)
			resp.Body.Close()
			break
		}
		
		resp.Body.Close() 

		fmt.Println("Ответ сервера:", string(body))

}
	
}

const localhost string = "127.0.0.1"
const defaultPort string = "8080"


func input() (string, string) {
	var ip, port string

	fmt.Println("введите адрес")
	fmt.Scan(&ip)

	fmt.Println("введите порт")
	fmt.Scan(&port)

	if ip == ""{
		ip = localhost
	}
	if port == ""{
		port = defaultPort
	}

return ip, port
}