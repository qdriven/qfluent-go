package main

import (
	"qfluent-go/internal/do"
)

type Engine interface{}

type engineImplem struct {
}

func NewEngine(i *do.Injector) (Engine, error) {
	return &engineImplem{}, nil
}
