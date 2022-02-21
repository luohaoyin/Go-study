package views

import (
	"Test/controller"
	"Test/db"
	"fmt"
)

func main(){
	if err := db.InitDB();err != nil {
		fmt.Printf("load config failed,err:#{err}\n")
		return
	}
	defer db.InitDB()
	controller.Menu()
}
