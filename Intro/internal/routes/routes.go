import (
	"net/http"
	"test"
	"fmt"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler) 
	mux.HandleFunc("/api/data", apiDataHandler)
	return mux
} 

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Intro Application!")
}
 
func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	data := test.GetSampleData()
	fmt.Fprintf(w, "Sample Data: %v", data)
}
