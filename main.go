package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	details "github.com/DSATAM/go-microservices/details"
	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking application health")
	response := map[string]string{
		"status":    "up",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}
func roothHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the hmepage")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and running")

}

func detailhHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the Details")
	hostName, err := details.GetHostName()
	if err != nil {
		panic(err)
	}
	IP := details.GetIP()

	fmt.Println(hostName, IP)
	response := map[string]string{
		"hostname": hostName,
		"ip":       IP.String(),
	}
	json.NewEncoder(w).Encode(response)

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/", roothHandler)
	r.HandleFunc("/details", detailhHandler)
	log.Println("Server Started")
	log.Fatal(http.ListenAndServe(":80", r))
}

/* package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s using token %s\n", r.URL.Path, r.URL.Query().Get("token"))
}
func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//		fmt.Fprintf(w, "Hello, you've requested: %s using token %s\n", r.URL.Path, r.URL.Query().Get("token"))
	//})

	http.HandleFunc("/", rootHandler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Web Server has started")
	http.ListenAndServe(":80", nil)
} */

/* package main

import (
	"fmt"
	"unsafe"

	"rsc.io/quote"

	geo "github.com/DSATAM/go-microservices/geometry"
)

func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return
}

func main() {

	name := "DevOps"
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	fmt.Printf("Type of name %T and size is %d\n", name, unsafe.Sizeof(name))

	a, p := rectProps(1, 2)

	fmt.Printf("Area is %f and perimeter is %f\n", a, p)

	//var daysOfTheMonth map[string]int
	//daysOfTheMonth["Jan"] = 31

	var daysOfTheMonth = map[string]int{"Jan": 31, "Feb": 28}
	fmt.Println(daysOfTheMonth)

	area := geo.Area(1, 2)
	diag := geo.Diagonal(1, 2)
	fmt.Println(area, diag)

}
*/
