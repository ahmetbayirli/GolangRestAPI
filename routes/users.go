package routes

import (
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
	"scroll2top.com/golang-rest-api/models"
	"scroll2top.com/golang-rest-api/utils"
)

// func getEvent(context *gin.Context) {
// 	eventId, err := strconv.ParseInt(context.Params.ByName("id"), 10, 64)

// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
// 		return
// 	}

// 	event, err := models.GetEventById(eventId)

// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server crush"})
// 	}
// 	context.JSON(http.StatusOK, event)
// }

// func getEvents(context *gin.Context) {
// 	events, err := models.GetAllEvents()
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server crush"})
// 	}
// 	context.JSON(http.StatusOK, events)
// }

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	// user.ID = 1
	// user.UserID = 1
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server crush"})
        return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User created ", "user": user})

}

func login(context *gin.Context) {
	var user models.User

    err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

    err = user.ValidateCredenentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate"})
        return
	}

    token, err := utils.GenerateToken(user.Email, user.ID)


	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate"})
        return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login succesful", "token": token})
}



// func deleteEvent(context *gin.Context) {
// 	eventId, err := strconv.ParseInt(context.Params.ByName("id"), 10, 64)

// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
// 		return
// 	}

// 	event, err := models.GetEventById(eventId)

// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server crush"})
// 	}

// 	err = event.Delete()

// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "Server crush"})
// 	}

// 	context.JSON(http.StatusOK, gin.H{"message": "Event deleted succesfully "})
// }
