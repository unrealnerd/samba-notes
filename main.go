package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	
	router.LoadHTMLGlob("templates/**/*")

	router.GET("/", rootIndex())
	router.GET("/notes", notesIndex())

	router.Run(":8080")

}

// func main() {

// 	http.HandleFunc("/", helllo)

// 	http.ListenAndServe(":8080", nil)

// }

// func helllo(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello back for request: %s", r.URL)
// }

func qhIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "quoteheader/index", gin.H{
			"PageTitle": "Hello1",
		})
	}
}

func rootIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "common/index", gin.H{
			"PageTitle": "Hello1",
		})
	}
}

func notesIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "notes/index", gin.H{
			"PageTitle": "Notes",
		})
	}
}
