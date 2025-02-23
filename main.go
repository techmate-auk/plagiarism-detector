package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// Function to generate n-grams
func generateNGrams(text string, n int) []string {
	words := strings.Fields(text)
	var ngrams []string
	for i := 0; i <= len(words)-n; i++ {
		ngrams = append(ngrams, strings.Join(words[i:i+n], " "))
	}
	return ngrams
}

// Function to calculate Jaccard similarity
func jaccardSimilarity(set1, set2 []string) float64 {
	set1Map := make(map[string]bool)
	set2Map := make(map[string]bool)

	for _, s := range set1 {
		set1Map[s] = true
	}
	for _, s := range set2 {
		set2Map[s] = true
	}

	intersection := 0
	union := 0

	for k := range set1Map {
		if set2Map[k] {
			intersection++
		}
		union++
	}
	for k := range set2Map {
		if !set1Map[k] {
			union++
		}
	}

	if union == 0 {
		return 0
	}
	return float64(intersection) / float64(union)
}

// Function to detect plagiarism
func detectPlagiarism(query string, corpus []string, n int, threshold float64) []string {
	var results []string
	queryNGrams := generateNGrams(query, n)

	for i, doc := range corpus {
		docNGrams := generateNGrams(doc, n)
		similarity := jaccardSimilarity(queryNGrams, docNGrams)

		if similarity >= threshold {
			results = append(results, fmt.Sprintf("Document %d (%.2f%% similar)", i+1, similarity*100))
		}
	}

	return results
}

// main function
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

		// Save the uploaded file temporarily
		tempFilePath := "./upload/" + file.Filename
		err = c.SaveUploadedFile(file, tempFilePath)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to save uploaded file"})
			return
		}

		// Read the content of the uploaded file
		queryContent, err := os.ReadFile(tempFilePath)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to read uploaded file"})
			return
		}

		// Load the corpus from files in a directory
		corpus := loadCorpus("./corpus")

		// Perform plagiarism detection
		results := detectPlagiarism(string(queryContent), corpus, 3, 0.2)

		// Delete the temporary file
		os.Remove(tempFilePath)

		// Respond with results
		if len(results) > 0 {
			c.JSON(200, gin.H{"message": "Plagiarism detected", "results": results})
		} else {
			c.JSON(200, gin.H{"message": "No plagiarism detected"})
		}
	})

	r.Run(":8080") // Start the server
}

// Function to load corpus documents from a directory
func loadCorpus(directory string) []string {
	var corpus []string

	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println("Error reading corpus directory:", err)
		return corpus
	}

	for _, file := range files {
		filePath := directory + "/" + file.Name()
		content, err := os.ReadFile(filePath)
		if err == nil {
			corpus = append(corpus, string(content))
		}
	}

	return corpus
}
