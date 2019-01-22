package main
import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprinttf(w, "hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8000", nil)
}

