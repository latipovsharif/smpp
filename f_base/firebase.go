package f_base

import (
	"smpp/ent"
	"smpp/rabbit"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"
)

func fbase(fb *Fbase, cacheMap *rabbit.CacheMap) {
	//lint:ignore SA4006 this value
	arrMes := []ent.Messages{}
	for {
		time.Sleep(1 * time.Second)
		iter := fb.FClient.Collection("message").Documents(fb.Ctx)
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalf("Failed to iterate: %v", err)
			}
			doc.DataTo(&fb.Msg)
			cacheMap.Mutex.RLock()
			arrMes = cacheMap.Hmap[uuid.MustParse(fb.Msg.UserID)]
			arrMes = append(arrMes, fb.Msg)
			cacheMap.Hmap[uuid.MustParse(fb.Msg.UserID)] = arrMes
			cacheMap.Mutex.RUnlock()
			_, err = fb.FClient.Collection("message").Doc(doc.Ref.ID).Delete(fb.Ctx)
			if err != nil {
				log.Errorf("cannot delete message in firebase %v", err)
			}
		}
	}
}
