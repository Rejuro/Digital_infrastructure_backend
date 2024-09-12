package main
    
import (
	"net/http"
	"log"
	"app/controllers"
)
    
func main() {
    http.HandleFunc("/", controllers.Sayhello) // Устанавливаем роутер

	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
