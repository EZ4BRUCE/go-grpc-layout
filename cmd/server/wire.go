//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/EZ4BRUCE/go-grpc-layout/configs/conf"
	"github.com/EZ4BRUCE/go-grpc-layout/internal/biz"
	"github.com/EZ4BRUCE/go-grpc-layout/internal/data"
	"github.com/EZ4BRUCE/go-grpc-layout/internal/server"
	"github.com/EZ4BRUCE/go-grpc-layout/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Consul, *conf.Global, log.Logger) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
