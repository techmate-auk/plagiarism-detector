package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Serve static files (CSS, JS, images)
	r.Static("/static", "./static")

	// Route for the home page
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Plagiarism Detector",
		})
	})

	// Route to handle file upload
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"error": "File upload failed"})
			return
		}

		// Save the uploaded file (or process it directly)
		c.SaveUploadedFile(file, "./uploads/"+file.Filename)

		// Perform plagiarism detection (implement your logic here)
		// ...

		c.JSON(200, gin.H{"message": "File uploaded successfully"})
	})

	r.Run(":8080") // Start the server
}
