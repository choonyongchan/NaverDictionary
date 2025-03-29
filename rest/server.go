package rest

import (
	"errors"
	"naverdictionary/scraper"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures and returns a Gin router
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Define routes
	router.GET("/", welcome)                     // Welcome Page
	router.GET("/get", get)                      // Get Dictionary Info
	router.GET("/get/entryinfo", getentryinfo)   // Get Entry Info Raw
	router.GET("/get/searchinfo", getsearchinfo) // Get Search Info RaW
	router.GET("/get/message", getmessage)       // Get Message

	return router
}

// StartServer initializes and starts the server
func StartServer() {
	router := SetupRouter()
	router.Run(":8080") // Listen on port 8080
}

func welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the Naver Scraper API!",
	})
}

func extractword(c *gin.Context) (string, error) {
	wordhex := c.Query("word") // Get the "word" query parameter
	if wordhex == "" {
		return "", errors.New("empty 'word' parameter")
	}
	word := string(wordhex)
	if word == "" {
		return "", errors.New("empty Word")
	}
	return word, nil
}

// Returns the Dictionary Info
func get(c *gin.Context) {
	word, errword := extractword(c) // Extract the word from the query parameter
	if errword != nil {
		c.JSON(400, gin.H{
			"error": errword.Error(),
		})
		return
	}

	dictinfo, errget := scraper.Get(word) // Pass the word to the scraper
	if errget != nil {
		c.JSON(500, gin.H{
			"error": errget.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": dictinfo,
	})
}

// Returns the Raw Entry Info
func getentryinfo(c *gin.Context) {
	word, errword := extractword(c) // Extract the word from the query parameter
	if errword != nil {
		c.JSON(400, gin.H{
			"error": errword.Error(),
		})
		return
	}

	entryinfo, errentryinfo := scraper.GetEntryInfoRaw(word) // Pass the word to the scraper
	if errentryinfo != nil {
		c.JSON(500, gin.H{
			"error": errentryinfo.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": entryinfo,
	})
}

// Returns the Raw Search Info
func getsearchinfo(c *gin.Context) {
	word, errword := extractword(c) // Extract the word from the query parameter
	if errword != nil {
		c.JSON(400, gin.H{
			"error": errword.Error(),
		})
		return
	}

	searchinfo, errsearchinfo := scraper.GetSearchInfoRaw(word) // Pass the word to the scraper
	if errsearchinfo != nil {
		c.JSON(500, gin.H{
			"error": errsearchinfo.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": searchinfo,
	})
}

// Returns the Message
func getmessage(c *gin.Context) {
	word, errword := extractword(c) // Extract the word from the query parameter
	if errword != nil {
		c.JSON(400, gin.H{
			"error": errword.Error(),
		})
		return
	}

	message, errmessage := scraper.GetMessage(word) // Pass the word to the scraper
	if errmessage != nil {
		c.JSON(500, gin.H{
			"error": errmessage.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": message,
	})
}
