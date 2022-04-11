package entity

import "monografia/service"

type Entity struct {
	service service.Service
}

func New(srv service.Service) *Entity {
	return &Entity{service: srv}
}
