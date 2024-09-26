package service

import "course-service/data"

var repo *data.Models

func InitService(m *data.Models) {
	repo = m
}
