package service

import (
	"github.com/TremblingV5/DouTok/applications/publish/dal/query"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Do query.IVideoDo
var Video = query.VideoStruct
