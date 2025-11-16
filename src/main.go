package main

import (
	"log"
	"os"
	"strings"

	"github.com/electerm/electerm-sync-server-go/src/store"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"status": "error", "message": "No authorization header"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"status": "error", "message": "Invalid token"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		users := strings.Split(os.Getenv("JWT_USERS"), ",")
		authorized := false
		for _, user := range users {
			if user == userId {
				authorized = true
				break
			}
		}

		if !authorized {
			c.JSON(401, gin.H{"status": "error", "message": "Unauthorized!"})
			c.Abort()
			return
		}

		c.Set("userId", userId)
		c.Next()
	}
}

func handleSync(c *gin.Context) {
	userId := c.GetString("userId")

	if c.Request.Method == "GET" {
		data, err := store.SQLiteStore.Read(userId)
		if err != nil {
			c.String(404, "File not found")
			return
		}
		c.JSON(200, data)
		return
	}

	if c.Request.Method == "POST" {
		c.String(200, "test ok")
		return
	}

	if c.Request.Method == "PUT" {
		var data map[string]interface{}
		if err := c.BindJSON(&data); err != nil {
			c.JSON(400, gin.H{"status": "error", "message": "Invalid JSON"})
			return
		}

		err := store.SQLiteStore.Write(userId, data)
		if err != nil {
			c.JSON(500, gin.H{"status": "error", "message": "Failed to write data"})
			return
		}

		c.String(200, "ok")
		return
	}

	c.JSON(405, gin.H{"status": "error", "message": "Method not allowed"})
}

func setupRouter() *gin.Engine {
	r := gin.New()

	// Add logging middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Handle unsupported methods
	r.NoMethod(func(c *gin.Context) {
		c.JSON(405, gin.H{"status": "error", "message": "Method not allowed"})
	})

	// Handle not found routes
	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method != "GET" && c.Request.Method != "POST" && c.Request.Method != "PUT" {
			c.JSON(405, gin.H{"status": "error", "message": "Method not allowed"})
			return
		}
		c.JSON(404, gin.H{"status": "error", "message": "Not found"})
	})

	r.GET("/test", func(c *gin.Context) {
		c.String(200, "ok")
	})

	authorized := r.Group("/api")
	authorized.Use(authMiddleware())
	{
		authorized.Any("/sync", handleSync)
	}

	return r
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize SQLite store
	if err := store.SQLiteStore.Init(); err != nil {
		log.Fatalf("Failed to initialize SQLite store: %v", err)
	}
	defer store.SQLiteStore.Close()

	log.Println("SQLite store initialized successfully")

	r := setupRouter()

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtUsers := os.Getenv("JWT_USERS")

	log.Print("\n========================================")
	log.Printf("🚀 Server running at http://%s:%s", host, port)
	log.Print("========================================\n")

	log.Println("📝 Configuration Guide:")
	log.Println("----------------------------------------")
	log.Print("In electerm sync settings, set custom sync server with:\n")
	log.Printf("  API URL:    http://%s:%s/api/sync\n", host, port)

	log.Println("\n🔐 Authentication:")
	log.Println("----------------------------------------")
	if jwtSecret == "" {
		log.Println("  JWT_SECRET:    ⚠️  NOT SET")
	} else {
		log.Println("  JWT_SECRET:    ✓ SET")
	}
	if jwtUsers == "" {
		log.Println("  JWT_USERS:     ⚠️  NOT SET")
	} else {
		log.Printf("  JWT_USERS:     %s", jwtUsers)
	}
	log.Print("========================================\n")

	r.Run(host + ":" + port)
}
