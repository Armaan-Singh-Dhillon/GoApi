package producthelpers

import (
	"fmt"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi i am  a product")
}
