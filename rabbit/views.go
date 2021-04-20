package rabbit

import (
	"context"
	"smpp/ent"
	"smpp/ent/price"
	"smpp/ent/user"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func (s *Session) chekBalans(ctx context.Context, userID uuid.UUID, countMessge int) (float64, float64, int32) {
	usr, err := s.db.User.Query().
		Where(user.ID(userID)).
		First(ctx)
	if err != nil {
		log.Errorf("cannot select User %v", err)
		return 0, 0, 0
	}
	messagePrice, err := s.db.Price.Query().
		Where(price.MinLTE(usr.Count+int32(countMessge)), price.MaxGT(usr.Count+int32(countMessge))).
		First(ctx)
	if err != nil {
		log.Errorf("cannot select Price %v", err)
		return 0, 0, 0
	} else if usr.Balance >= messagePrice.Price {
		return usr.Balance, messagePrice.Price, usr.Count
	}
	return 0, 0, 0
}

func inserMessage(bulkMessages []*ent.MessagesCreate, listMessage []ent.Messages, db *ent.Client) {
	ctx := context.Background()
	listUserMes := []UserMessage{}
	for index, message := range listMessage {
		bulkMessages[index] = db.Messages.Create().
			SetSequenceNumber(message.SequenceNumber).
			SetExternalID(message.ExternalID).
			SetDst(message.Dst).
			SetMessage(message.Message).
			SetSrc(message.Src).
			SetState(message.State).
			SetSmscMessageID(message.SmscMessageID).
			SetProviderIDID(uuid.MustParse(message.ProviderID)).
			SetUserIDID(uuid.MustParse(message.UserID))
		listUserMes = append(listUserMes, UserMessage{
			IDUser:     uuid.MustParse(message.UserID),
			IDProvider: uuid.MustParse(message.ProviderID),
		})
	}
	_, err := db.Messages.CreateBulk(bulkMessages...).Save(ctx)
	if err != nil {
		log.Errorf("cannot creat Messages %v", err)
	}
	inserUserMess(listUserMes, db)

}

func inserUserMess(listUserMes []UserMessage, db *ent.Client) {
	ctx := context.Background()
	bulk := make([]*ent.UserMonthMessageCreate, len(listUserMes))
	for index, v := range listUserMes {
		bulk[index] = db.UserMonthMessage.Create().
			SetProviderIDID(v.IDProvider).
			SetUserIDID(v.IDUser)
	}
	_, err := db.UserMonthMessage.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		log.Errorf("cannot creat UserMonthMessage %v", err)
	}
}
