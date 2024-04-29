package models

type ProjectProgress struct {
    ProjectName     string   `bson:"projectName"`
    TasksAccomplished []string `bson:"tasksAccomplished"`
    PendingTasks    []string `bson:"pendingTasks"`
    Constraints     []string `bson:"constraints"`
    Remarks         string   `bson:"remarks"`
}