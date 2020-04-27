package utils

import "fmt"

type CardNotFoundError struct{}

func (e *CardNotFoundError) Error() string {
	return "No valid card found either randomly or with your current query"
}

type CardNotURPairError struct {
	cardId   int64
	idolized bool
}

func (e *CardNotURPairError) Error() string {
	return fmt.Sprintf("Selected card: %v (Idolized?: %v) doesn't have a UR Pair", e.cardId, e.idolized)
}
