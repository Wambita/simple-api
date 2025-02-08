package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// User structure
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	dataFile = "data.json"
	mutex    = &sync.Mutex{}
)

// loadData loads users from the JSON file
func loadData() map[string]User {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]User) // Return empty map if file doesn't exist
		}
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var data map[string]User
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return make(map[string]User)
	}
	return data
}

// saveData saves the user data to the JSON file
func saveData(data map[string]User) {
	file, err := os.Create(dataFile) // Overwrite existing file
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty-print JSON
	if err := encoder.Encode(data); err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}

func main() {
	router := SetupRouter()
	router.Run(":8080")
}

// SetupRouter sets up the API routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		user.ID = uuid.New().String()

		mutex.Lock()
		dataStore := loadData() // Load existing data
		dataStore[user.ID] = user
		saveData(dataStore) // Save back to file
		mutex.Unlock()

		c.JSON(http.StatusOK, gin.H{"id": user.ID})
	})

	router.GET("/get/:id", func(c *gin.Context) {
		id := c.Param("id")

		mutex.Lock()
		dataStore := loadData() // Load data from file
		user, exists := dataStore[id]
		mutex.Unlock()

		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	return router
}
