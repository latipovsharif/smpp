package rabbit

import (
	"context"
	"smpp/ent"
	"smpp/ent/user"
	"sync"
	"time"

	"github.com/google/uuid"
)

//CacheMap  CacheMap in cache
type CacheMap struct {
	Hmap  map[uuid.UUID][]ent.Messages
	Mutex *sync.RWMutex
}

type UserMessage struct {
	IDUser     uuid.UUID
	IDProvider uuid.UUID
}

func (s *Session) SendingMessage(c chan<- ent.Messages, cache *CacheMap) {
	ctx := context.Background()
	listMessage := []ent.Messages{}
	//lint:ignore SA4006 this value
	bulkMessages := make([]*ent.MessagesCreate, len(listMessage))
	for {
		cache.Mutex.Lock()
		for uid, v := range cache.Hmap {
			balans, price, count := s.chekBalans(ctx, uid, len(v))
			for _, v1 := range v {
				if balans-(price*float64(count)) > 0 && balans != 0.0 {
					v1.State = int(StateNew)
					count++
					c <- v1
				} else {
					v1.State = int(InvalidBalance)
				}
				listMessage = append(listMessage, v1)
			}
			bulkMessages = make([]*ent.MessagesCreate, len(listMessage))
			inserMessage(bulkMessages, listMessage, s.db)
			if count > 0 {
				s.db.User.Update().Where(user.ID(uid)).SetCount(count).Save(ctx)
			}
			delete(cache.Hmap, uid)
			listMessage = nil
		}
		cache.Mutex.Unlock()
		time.Sleep(1 * time.Second)
	}
}
