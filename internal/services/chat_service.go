// /internal/services/chat_service.go
package services

import (
	"math/rand"

	"github.com/google/uuid"
)

type ChatService struct {
	waitingUsers []uuid.UUID
}

func NewChatService() *ChatService {
	return &ChatService{waitingUsers: []uuid.UUID{}}
}

func (cs *ChatService) PairUser(userID uuid.UUID) uuid.UUID {
	if len(cs.waitingUsers) > 0 {
		partner := cs.waitingUsers[rand.Intn(len(cs.waitingUsers))]
		cs.waitingUsers = remove(cs.waitingUsers, partner)
		return partner
	}
	cs.waitingUsers = append(cs.waitingUsers, userID)
	return uuid.Nil
}

func remove(slice []uuid.UUID, s uuid.UUID) []uuid.UUID {
	for i, other := range slice {
		if other == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
