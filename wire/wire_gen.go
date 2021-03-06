// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package wire

import (
	"awesomeProject/service"
)

// Injectors from wire.go:

func NewUserService(id int, name string) (*service.UserService, error) {
	userService := service.NewUserService(id, name)
	return userService, nil
}

func InitLogService(name string) (*service.LogService, error) {
	logService := service.NewLogService(name)
	return logService, nil
}

func InitConfigService(msg []string, ip string) (*service.ConfigService, error) {
	option := &service.Option{
		Message: msg,
		Ip:      ip,
	}
	configService, err := service.NewConfigService(option)
	if err != nil {
		return nil, err
	}
	return configService, nil
}
