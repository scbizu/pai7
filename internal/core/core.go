package core

import (
	"bytes"
	"container/list"
	"errors"
	"fmt"
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
	InitCardNumber CardNumber = 7
	// MaxCardNumber == K
	MaxCardNumber CardNumber = 13
	// MinCardNumber == A
	MinCardNumber CardNumber = 1
)

type MSets struct {
	sets    map[Kind]*list.List
	dropped map[string][]CardNumber
}

var (
	ErrInsertInvalid = errors.New("msets: insert failed")
	ErrInvalidCard   = errors.New("msets: validator: invalid card number")
	ErrKindNotFound  = errors.New("msets: find: kind not found")
	ErrSkipNoFirst7  = errors.New("msets: validator: no 7 set")
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
	if l.Len() == 0 {
		l.PushBack(InitCardNumber)
		return nil
	}
	max := l.Back().Value
	if max == nil {
		max = InitCardNumber
	}
	maxNum, ok := max.(CardNumber)
	if !ok {
		return ErrInsertInvalid
	}
	min := l.Front().Value
	if min == nil {
		min = InitCardNumber
	}
	minNum, ok := min.(CardNumber)
	if !ok {
		return ErrInsertInvalid
	}
	if number > InitCardNumber && number == maxNum+1 {
		l.PushBack(number)
	}
	if number <= InitCardNumber && number == minNum-1 {
		l.PushFront(number)
	}

	return nil
}

func isNumValid(num CardNumber) bool {
	return num > 0 && num < MaxCardNumber+1
}

func (ms *MSets) Find(kind Kind) (CardNumber, CardNumber, error) {
	l, ok := ms.sets[kind]
	if !ok {
		return 0, 0, ErrKindNotFound
	}
	if l.Len() == 0 {
		return 0, 0, nil
	}
	return l.Front().Value.(CardNumber), l.Back().Value.(CardNumber), nil
}

func sum(nums []CardNumber) int64 {
	var res int64
	for _, num := range nums {
		res += int64(num)
	}
	return res
}

func (ms MSets) PrintDropped() string {
	var dropped bytes.Buffer
	dropped.WriteString("Dropped: \n")
	for name, nums := range ms.dropped {
		dropped.WriteString(fmt.Sprintf("%s: %d \n", name, sum(nums)))
	}
	return dropped.String()
}

func (ms *MSets) PrintStatus() string {
	var status bytes.Buffer
	for kind, l := range ms.sets {
		var kindStatus bytes.Buffer
		for node := l.Front(); node != nil; node = node.Next() {
			labelStr := CardNumberLabel[node.Value.(CardNumber)]
			if node != l.Back() {
				kindStatus.WriteString(fmt.Sprintf("%s -> ", labelStr))
			} else {
				kindStatus.WriteString(labelStr)
			}
		}
		lableKindStr := CardKindLabel[kind]
		status.WriteString(fmt.Sprintf("Kind: %s, List: %s\n", lableKindStr, kindStatus.String()))
	}
	// status.WriteString("Dropped: \n")
	// for kind, numbers := range ms.dropped {
	// 	status.WriteString(fmt.Sprintf("Kind: %s, Numbers:%v\n", CardKindLabel[kind], numbers))
	// }
	return status.String()
}

func (ms *MSets) IsFirst7Set() bool {
	l := ms.sets[KindBlackHeart]
	return l.Len() >= 1
}

func (ms *MSets) GetValidNums(kind Kind, nums []CardNumber) (map[Kind][]CardNumber, error) {
	res := make(map[Kind][]CardNumber)
	min, max, err := ms.Find(kind)
	if err != nil {
		return nil, err
	}
	if !ms.IsFirst7Set() {
		// allows first 7
		if kind == KindBlackHeart {
			for _, num := range nums {
				if num == InitCardNumber {
					res[kind] = []CardNumber{InitCardNumber}
					return res, nil
				}
			}
		} else {
			return nil, ErrSkipNoFirst7
		}
	}
	var validNums []CardNumber
	for _, num := range nums {
		if (num == min-1 || num == max+1) && (max > 0 && min > 0) {
			validNums = append(validNums, num)
		}
		if num == InitCardNumber {
			validNums = append(validNums, num)
		}
	}
	res[kind] = append(res[kind], validNums...)

	return res, nil
}

func (ms *MSets) AddDroppedNum(key string, num CardNumber) {
	ms.dropped[key] = append(ms.dropped[key], num)
}
