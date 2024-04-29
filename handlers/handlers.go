package handlers

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "time"
	"os"
    "github.com/OcheOps/ContractB/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)

var client *mongo.Client

func init() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    mongoURI := LoadMongoURI()
    var err error
    client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatal(err)
    }
}

func LoadMongoURI() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        log.Fatal("MONGO_URI is not set in the .env file")
    }

    return mongoURI
}


func ReportHandler(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    projectDetailsCollection := client.Database("your-database").Collection("projectDetails")
    projectProgressCollection := client.Database("your-database").Collection("projectProgress")

    var projectDetails []models.ProjectDetails
    cursor, err := projectDetailsCollection.Find(ctx, bson.M{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    if err = cursor.All(ctx, &projectDetails); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var projectProgress []models.ProjectProgress
    cursor, err = projectProgressCollection.Find(ctx, bson.M{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    if err = cursor.All(ctx, &projectProgress); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    report := struct {
        ProjectDetails  []models.ProjectDetails
        ProjectProgress []models.ProjectProgress
    }{
        ProjectDetails:  projectDetails,
        ProjectProgress: projectProgress,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(report)
}

func CreateProjectDetailsHandler(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    projectDetailsCollection := client.Database("your-database").Collection("projectDetails")

    var projectDetails models.ProjectDetails
    if err := json.NewDecoder(r.Body).Decode(&projectDetails); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result, err := projectDetailsCollection.InsertOne(ctx, projectDetails)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result.InsertedID)
}

func CreateProjectProgressHandler(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    projectProgressCollection := client.Database("your-database").Collection("projectProgress")

    var projectProgress models.ProjectProgress
    if err := json.NewDecoder(r.Body).Decode(&projectProgress); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result, err := projectProgressCollection.InsertOne(ctx, projectProgress)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result.InsertedID)
}