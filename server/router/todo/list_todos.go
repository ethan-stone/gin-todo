package todo

import (
	"net/http"
	"strconv"

	"github.com/ethan-stone/gin-todo/db"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func List(c *gin.Context) {
	skipQuery := c.Query("skip")
	limitQuery := c.Query("limit")

	var skip int

	if skipQuery == "" {
		skip = 0
	} else {
		skipParse, skipParseErr := strconv.Atoi(skipQuery)
		if skipParseErr != nil {
			log.Error(skipParseErr)
			c.JSON(http.StatusBadRequest, gin.H{"error": skipParseErr.Error()})
			return
		}
		skip = skipParse
	}

	var limit int

	if limitQuery == "" {
		limit = 50 
	} else {
		limitParse, limitParseErr := strconv.Atoi(limitQuery)
		if limitParseErr != nil {
			log.Error(limitParseErr)
			c.JSON(http.StatusBadRequest, gin.H{"error": limitParseErr.Error()})
			return
		}
		limit =limitParse 
	}

	var todos []db.Todo

	result := db.DB.Limit(limit).Offset(skip).Find(&todos)

	if result.Error != nil {
		log.Error(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	
	log.Infof("Todos retrieved")
	c.JSON(http.StatusOK, gin.H{"data": todos})
	return
}