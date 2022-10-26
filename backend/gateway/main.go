package gateway

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<p> %s </p>", "hi")
	})
	http.ListenAndServe("localhost:8080", nil)
}
