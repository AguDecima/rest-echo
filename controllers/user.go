package controllers

import (
	"net/http"
	"rest-echo/models"

	"github.com/labstack/echo"
)

// GetUser is...
func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	var user models.User

	if dbc := models.DB.Find(&user, id); dbc.Error != nil {
		var res map[string]interface{}
		res["message"] = "No se encontro el registro"
		res["status"] = false
		return c.JSON(http.StatusOK, res)
	}

	return c.JSON(http.StatusOK, user)

}

// Save is...
func Save(c echo.Context) error {
	// Get name and email
	//name := c.FormValue("name")
	//email := c.FormValue("email")
	user := &models.User{Name: "Agustin", LastName: "Decima", DNI: "38652010"}
	models.DB.Save(&user)
	return c.JSON(http.StatusCreated, user)
}
