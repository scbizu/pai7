package core

import (
	"container/list"
	"errors"
)

type Kind uint8
type CardNumber int8

const (
	// KindRedHeart  ❤️
	KindRedHeart Kind = iota
	// KindBlackHeart 🖤
	KindBlackHeart
	// KindGrassFlower ♣️
	KindGrassFlower
	// KindCube ♦️
	KindCube
)

const (
	InitCartNumber CardNumber = 7
)

type MSets struct {
	sets    map[Kind]*list.List
	dropped map[string][]CardNumber
}

var (
	ErrInsertInvalid = errors.New("msets: insert failed")
	ErrInvalidCard   = errors.New("msets: validator: invalid card number")
	ErrKindNotFound  = errors.New("msets: find: kind not found")
)

func NewMSets() *MSets {
	return &MSets{
		sets: map[Kind]*list.List{
			KindRedHeart:    list.New(),
			KindBlackHeart:  list.New(),
			KindGrassFlower: list.New(),
			KindCube:        list.New(),
		},
		dropped: make(map[string][]CardNumber),
	}
}

func (ms *MSets) Insert(kind Kind, number CardNumber) error {
	if !isNumValid(number) {
		return ErrInvalidCard
	}
	l, ok := ms.sets[kind]
	if !ok {
		return ErrInsertInvalid
	}
	max := l.Back().Value
	if max == nil {
		max = CardNumber(7)
	}
	maxNum, ok := max.(CardNumber)
	if !ok {
		return ErrInsertInvalid
	}
	min := l.Front().Value
	if min == nil {
		min = CardNumber(7)
	}
	minNum, ok := min.(CardNumber)
	if !ok {
		return ErrInsertInvalid
	}
	if number > 7 && number == maxNum+1 {
		l.PushBack(number)
	}
	if number <= 7 && number == minNum-1 {
		l.PushFront(number)
	}

	return nil
}

func isNumValid(num CardNumber) bool {
	return num > 0 && num < 14
}

func (ms *MSets) Find(kind Kind) (CardNumber, CardNumber, error) {
	l, ok := ms.sets[kind]
	if !ok {
		return 0, 0, ErrKindNotFound
	}
	if l.Len() == 0 {
		return InitCartNumber, InitCartNumber, nil
	}
	return l.Front().Value.(CardNumber), l.Back().Value.(CardNumber), nil
}

func (ms *MSets) GetValidNums(kind Kind, nums []CardNumber) (map[Kind][]CardNumber, error) {
	res := make(map[Kind][]CardNumber)
	min, max, err := ms.Find(kind)
	if err != nil {
		return nil, err
	}
	var validNums []CardNumber
	for _, num := range nums {
		if num == min-1 || num == max+1 {
			validNums = append(validNums, num)
		}
	}
	res[kind] = validNums

	return res, nil
}
