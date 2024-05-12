package db

import (
	"bitoOA/internal/logger"
)

type service struct {
	table map[string]*User
	log   *logger.Logger
}

func New(log *logger.Logger) Service {
	return &service{
		table: map[string]*User{},
		log:   log.With("service", "db"),
	}
}
