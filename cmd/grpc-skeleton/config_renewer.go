package main

import (
	"context"
	"go.citilink.cloud/citizap/factory"
)

// ConfigsRenewer обновитель конфигов
type ConfigsRenewer interface {
	// RenewConfig обновляет конфиги
	RenewConfig(ctx context.Context) error
}

// EmptyConfigsRenewer обновляет конфиги
type EmptyConfigsRenewer struct {
	loggerFactory factory.Factory
}

// NewEmptyConfigsRenewer инициализирует и возвращает EmptyConfigsRenewer
func NewEmptyConfigsRenewer(loggerFactory factory.Factory) *EmptyConfigsRenewer {
	return &EmptyConfigsRenewer{
		loggerFactory: loggerFactory,
	}
}

// RenewConfig обновляет конфиги
func (r *EmptyConfigsRenewer) RenewConfig(ctx context.Context) error {
	logger := r.loggerFactory.Create(ctx)
	logger.Info("renew configs initialized")

	return nil
}
