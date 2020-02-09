package i18n

import "errors"

var (
	ErrGameExisted    = errors.New("game: Game already existed,type /close to close game at first.")
	ErrGameNotExisted = errors.New("game: Game has not created before, type /new to init a new game.")
)
