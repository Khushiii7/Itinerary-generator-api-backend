package main

import (
    "Itenary_Backend_API/routers"
)
//start the server
func main() {
    r := routers.SetupRouter()
    r.Run(":8080")
}
