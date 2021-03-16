package server

import "github.com/gin-gonic/gin"

func (s *Server) getRoutes(r *gin.RouterGroup) {
	m := r.Group("/messages")
	m.POST("/message/", s.newMessage)
}
