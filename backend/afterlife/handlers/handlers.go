package handlers

import (
	"github.com/byuoitav/afterlife"
	"go.uber.org/zap"
)

type Handlers struct {
	DataService afterlife.DataService
	*zap.Logger
}
