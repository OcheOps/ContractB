package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProjectDetails struct {
    ID                 primitive.ObjectID `bson:"_id,omitempty"`
    PTBReference       string             `bson:"ptbReference"`
    ProjectName        string             `bson:"projectName"`
    AwardExecutionPeriod string           `bson:"awardExecutionPeriod"`
    ProjectSupervisor  string             `bson:"projectSupervisor"`
}