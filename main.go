package main

import (
    "vietanh.com/gohero/mappings"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
   mappings.CreateUrlMappings()
   mappings.Router.Run(":8080")
}


