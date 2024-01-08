package userdata

import (
	"elichika/client"
	"elichika/gamedata"
	"elichika/utils"
)

// fetch a card, use the value in diff is present, otherwise fetch from db
func (session *Session) GetUserCard(cardMasterId int32) client.UserCard {
	pos, exist := session.UserCardMapping.SetList(&session.UserModel.UserCardByCardId).Map[int64(cardMasterId)]
	if exist {
		return session.UserModel.UserCardByCardId.Objects[pos]
	}
	card := client.UserCard{}
	exist, err := session.Db.Table("u_card").
		Where("user_id = ? AND card_master_id = ?", session.UserId, cardMasterId).Get(&card)
	utils.CheckErr(err)

	if !exist {
		gamedata := session.Ctx.MustGet("gamedata").(*gamedata.Gamedata)
		masterCard := gamedata.Card[cardMasterId]
		card = client.UserCard{
			CardMasterId:        cardMasterId,
			Level:               1,
			MaxFreePassiveSkill: masterCard.PassiveSkillSlot,
			Grade:               -1, // check this for new card
			ActiveSkillLevel:    1,
			PassiveSkillALevel:  1,
			PassiveSkillBLevel:  1,
			PassiveSkillCLevel:  1,
			AcquiredAt:          int32(session.Time.Unix()),
			IsNew:               true,
		}
	}
	return card
}

func (session *Session) UpdateUserCard(card client.UserCard) {
	session.UserCardMapping.SetList(&session.UserModel.UserCardByCardId).Update(card)
}

func cardFinalizer(session *Session) {
	for _, card := range session.UserModel.UserCardByCardId.Objects {
		affected, err := session.Db.Table("u_card").
			Where("user_id = ? AND card_master_id = ?", session.UserId, card.CardMasterId).AllCols().Update(card)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_card", card)
		}
	}
}

// insert all the cards
func (session *Session) InsertCards(cards []client.UserCard) {
	session.UserModel.UserCardByCardId.Objects = append(session.UserModel.UserCardByCardId.Objects, cards...)
}

func init() {
	addFinalizer(cardFinalizer)
	addGenericTableFieldPopulator("u_card", "UserCardByCardId")
}
