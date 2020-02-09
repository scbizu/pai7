package i18n

import (
	"errors"
)

func Err(err error) error {
	switch err {
	case ErrGameExisted:
		return errors.New("游戏已经存在,请创建人先用 /close 关闭上一局游戏")
	case ErrGameNotExisted:
		return errors.New("游戏还未创建，输入 /new 创建")
	default:
		return err
	}
}
