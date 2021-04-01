package rabbit

import (
	"context"
	"fmt"
	"smpp/ent"
	"smpp/ent/price"
	"smpp/ent/user"
	"sync"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

//CacheMap  CacheMap in cache
type CacheMap struct {
	Hmap  map[uuid.UUID][]ent.Messages
	Mutex *sync.RWMutex
}

func (s *Session) SendingMessage(c chan<- ent.Messages, cache *CacheMap) {
	ctx := context.Background()
	listMessage := []ent.Messages{}
	for {
		cache.Mutex.Lock()
		for uuid, v := range cache.Hmap {
			balans, price, count := s.chekBalans(ctx, uuid)
			fmt.Println("Count |||||||", count)
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
			bulk := make([]*ent.MessagesCreate, len(listMessage))
			s.db.Messages.CreateBulk(bulk...).Save(ctx)
			if count != 0 {
				s.db.User.Update().Where(user.ID(uuid)).SetCount(count).Save(ctx)
			}
			delete(cache.Hmap, uuid)
			listMessage = nil
		}
		cache.Mutex.Unlock()
		time.Sleep(2 * time.Second)
	}

}

func (s *Session) chekBalans(ctx context.Context, userID uuid.UUID) (float64, float64, int32) {
	usr, err := s.db.User.Query().
		Where(user.IDEQ(userID)).
		First(ctx)
	if err != nil {
		log.Errorf("cannot select User %v", err)
		return 0, 0, 0
	}
	messagePrice, err := s.db.Price.Query().
		Where(price.MinLTE(usr.Count), price.MaxGT(usr.Count)).
		First(ctx)
	if err != nil {
		return 0, 0, 0
	} else if usr.Balance >= messagePrice.Price {
		return usr.Balance, messagePrice.Price, usr.Count
	}
	return 0, 0, 0

}