package handler

import "gateway-service/pb"

type LessonHandler struct {
	lessonService pb.LessonServiceClient
}

func NewLessonHandler(lessonService pb.LessonServiceClient) *LessonHandler {
	return &LessonHandler{
		lessonService: lessonService,
	}
}

