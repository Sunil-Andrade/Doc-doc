package main

import (
	"docdoc/document"
	"fmt"
)

func main() {
	// conn := db.ConnectDB()
	// _, err := db.GetDocumentByCode(context.Background(), conn, "XJHUDG")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	document.CreateDocumentFile("s")
	fmt.Println(document.GenerateCode())
}
