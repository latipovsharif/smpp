package server

import (
	"os"
	"smpp/smsc"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/sirupsen/logrus"
)

type Server struct {
	l       *logrus.Logger
	e       *gin.Engine
	db      *pg.DB
	session *smsc.Session
}

// Run is entry point to application
func (s *Server) Run(logger *logrus.Logger, db *pg.DB, session *smsc.Session) {
	logger.Info("starting application")
	s.db = db
	s.l = logger
	s.e = gin.Default()
	s.getRoutes(s.e.Group("/v1"))
	s.session = session

	if err := s.e.Run(os.Getenv("SMS_HOST_ADDR")); err != nil {
		panic("cannot run server")
	}
}
