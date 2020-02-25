package main
/**
Нужно на golang написать программу, которая:
1. Принимает считывает все переданные аргументы при старте и запускает web server на порту 8099.
2. На любой GET запрос отдаёт эти аргументы в виде JSON массива и кода ответа 200.
3. Принимает POST и PUT запросы, разбивает тело запроса на строки по Whitespace characters, и добавляет их
в массив с аргументами и возвращает код ответа 204 и пустое тело ответа.
4. В случае ошибки возвращает код ошибки 500 и пустой body http ответа.
5. Не выполняет никакой работы когда принимает запросы с любыми другими методами.

6. Задание должно быть оформлено в виде репозитория на github.

*/
import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

var args = os.Args[1:]

func handler(w http.ResponseWriter, r *http.Request) {
	//Если наш запрос - GET
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(args)
		//Вывод ошибки 500 если метод Encode не отработает
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	//Если наш запрос == Put или POST
	if r.Method == "PUT" || r.Method == "POST" {
		reqBody, err := ioutil.ReadAll(r.Body)
		//вывод ошибки 500 если Энкод не отработает
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		//Разделение строки на подстроки с помощью регулярки

		pattern := regexp.MustCompile(`\s+`)
		newArgs := pattern.Split(string(reqBody), -1)

		args = append(args, newArgs...)
		w.WriteHeader(http.StatusNoContent)
	}
}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8099", nil))
}
