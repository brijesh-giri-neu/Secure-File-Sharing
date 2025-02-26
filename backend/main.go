package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
    "github.com/gorilla/mux"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/go-redis/redis/v8"
    "context"
)

var db *gorm.DB
var redisClient *redis.Client
var ctx = context.Background()

func init() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system environment variables.")
    }

    // Connect to PostgreSQL
    dsn := os.Getenv("DATABASE_URL")
    var err error
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Initialize Redis
    redisClient = redis.NewClient(&redis.Options{
        Addr: os.Getenv("REDIS_URL"),
    })
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(map[string]string{"message": "Secure File Transfer API is running!"})
    }).Methods("GET")

    r.HandleFunc("/upload", uploadFile).Methods("POST")
    r.HandleFunc("/download/{filename}", downloadFile).Methods("GET")
    r.HandleFunc("/share", shareFile).Methods("POST")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Println("Server running on port", port)
    http.ListenAndServe(":"+port, r)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
    file, _, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Invalid file upload", http.StatusBadRequest)
        return
    }
    defer file.Close()
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "File uploaded successfully"})
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    filename := vars["filename"]
    filePath := "./uploads/" + filename
    http.ServeFile(w, r, filePath)
}

func shareFile(w http.ResponseWriter, r *http.Request) {
    var request struct {
        Filename string `json:"filename"`
        User    string `json:"user"`
    }
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "File shared successfully", "filename": request.Filename, "sharedWith": request.User})
}
