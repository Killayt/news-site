package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	rend(c.Writer, "Index")
}

func NewsHandler(c *gin.Context) {
	rend(c.Writer, "news")
}
func rend(w http.ResponseWriter, message string) {
	w.Write([]byte(message))
}
