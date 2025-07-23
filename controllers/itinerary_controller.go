package controllers
import (
    "net/http"
    "github.com/gin-gonic/gin"
    "Itenary_Backend_API/models"
    "Itenary_Backend_API/services"
)
// It handles the request to generate an itinerary PDF
//expects a JSON body with the details and returns the path to the generated PDF
func GenerateItinerary(c *gin.Context) {
    var req models.ItineraryRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return;
    }
    pdfPath, err := services.GeneratePDF(req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "success", "pdf_path": pdfPath})
}
