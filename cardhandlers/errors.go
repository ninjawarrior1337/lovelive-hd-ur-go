package cardhandlers

import "fmt"

type CardNotFoundError struct {
	cardId   int64
	idolized bool
}

func (e *CardNotFoundError) Error() string {
	return fmt.Sprintf("Selected card: %v (Idolized?: %v) is not a valid card, please choose another one", e.cardId, e.idolized)
}

type CardNotURPairError struct {
	cardId   int64
	idolized bool
}

func (e *CardNotURPairError) Error() string {
	return fmt.Sprintf("Selected card: %v (Idolized?: %v) doesn't have a UR Pair", e.cardId, e.idolized)
}
