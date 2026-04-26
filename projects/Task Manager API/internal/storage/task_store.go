package storage

import "task-manager/internal/models"

var Tasks = []models.Task{} //so here we are creating a exportable global variable
var NextID = 1              // this will where the task id start from and also it will increment later
