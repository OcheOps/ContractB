package handlers

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "time"

    "github.com/your-project/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
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
    // Load MongoDB URI from .env file
    // ...
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