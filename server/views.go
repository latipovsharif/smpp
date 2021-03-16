package server

import (
	"smpp/rabbit"

	"github.com/gin-gonic/gin"
)

func (s *Server) newMessage(c *gin.Context) {
	m := rabbit.Message{}
	if err := c.BindJSON(&m); err != nil {
		badRequest(c, "all fields required")
		return
	}

	_, err := s.db.Model(&m).Insert()
	if err != nil {
		internalError(c, "cannot insert message")
		return
	}

	accepted(c)
}
