// +build wireinject

package main

import (
	"github.com/farwydi/thunderstorm/domain"
	"github.com/google/wire"
)

func setup(domain.Config) (application, func(), error) {
	panic(wire.Build(
		newApplication,
	))
}
