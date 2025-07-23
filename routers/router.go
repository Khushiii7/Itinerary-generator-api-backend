package routers

import (
    "github.com/gin-gonic/gin"
    "Itenary_Backend_API/controllers"
)
//defines the routes
func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.POST("/generate-itinerary", controllers.GenerateItinerary)
    return r
}
