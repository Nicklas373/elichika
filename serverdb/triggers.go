package serverdb

import (
	"elichika/generic"
	"elichika/model"
	"elichika/utils"

	"time"
)

// card grade up trigger is responsible for showing the pop-up animation when openning a card after getting a new copy
// or right after performing a limit break using items
func (session *Session) AddTriggerCardGradeUp(id int64, trigger *model.TriggerCardGradeUp) {
	if id == 0 {
		id = time.Now().UnixNano()
	}
	if trigger != nil {
		trigger.TriggerID = id
		trigger.UserID = session.UserStatus.UserID
	}
	session.TriggerCardGradeUps = append(session.TriggerCardGradeUps, id)
	session.TriggerCardGradeUps = append(session.TriggerCardGradeUps, trigger)
	if trigger != nil {
		dbTrigger := *trigger
		dbTrigger.BeforeLoveLevelLimit = dbTrigger.AfterLoveLevelLimit
		// db trigger when login have BeforeLoveLevelLimit = AfterLoveLevelLimit
		// if the 2 numbers are equal the level up don't show when we open the card.
		_, err := Engine.Table("s_user_trigger_card_grade_up").Insert(&dbTrigger)
		utils.CheckErr(err)
	} else {
		_, err := Engine.Table("s_user_trigger_card_grade_up").Where("trigger_id = ?", id).Delete(
			&model.TriggerCardGradeUp{})
		utils.CheckErr(err)
	}
}

func (session *Session) GetAllTriggerCardGradeUps() generic.ObjectByObjectIDWrite[*model.TriggerCardGradeUp] {
	triggers := generic.ObjectByObjectIDWrite[*model.TriggerCardGradeUp]{}
	err := Engine.Table("s_user_trigger_card_grade_up").
		Where("user_id = ?", session.UserStatus.UserID).Find(&triggers.Objects)
	utils.CheckErr(err)
	triggers.Length = len(triggers.Objects)
	return triggers
}

func (session *Session) AddTriggerBasic(id int64, trigger *model.TriggerBasic) {
	if id == 0 {
		id = time.Now().UnixNano()
	}
	if trigger != nil {
		trigger.TriggerID = id
		trigger.UserID = session.UserStatus.UserID
	}
	session.TriggerBasics = append(session.TriggerBasics, id)
	session.TriggerBasics = append(session.TriggerBasics, trigger)
	if trigger != nil {
		_, err := Engine.Table("s_user_trigger_basic").Insert(trigger)
		utils.CheckErr(err)
	} else {
		_, err := Engine.Table("s_user_trigger_basic").Where("trigger_id = ?", id).Delete(
			&model.TriggerBasic{})
		utils.CheckErr(err)
	}
}

func (session *Session) GetAllTriggerBasics() generic.ObjectByObjectIDWrite[*model.TriggerBasic] {
	triggers := generic.ObjectByObjectIDWrite[*model.TriggerBasic]{}
	err := Engine.Table("s_user_trigger_basic").
		Where("user_id = ?", session.UserStatus.UserID).Find(&triggers.Objects)
	utils.CheckErr(err)
	triggers.Length = len(triggers.Objects)
	return triggers
}

func (session *Session) AddTriggerMemberLoveLevelUp(id int64, trigger *model.TriggerMemberLoveLevelUp) {
	if id == 0 {
		id = time.Now().UnixNano()
	}
	if trigger != nil {
		trigger.TriggerID = id
		trigger.UserID = session.UserStatus.UserID
	}
	session.TriggerMemberLoveLevelUps = append(session.TriggerMemberLoveLevelUps, id)
	session.TriggerMemberLoveLevelUps = append(session.TriggerMemberLoveLevelUps, trigger)
	if trigger != nil {
		_, err := Engine.Table("s_user_trigger_member_love_level_up").Insert(trigger)
		utils.CheckErr(err)
	} else {
		_, err := Engine.Table("s_user_trigger_member_love_level_up").Where("trigger_id = ?", id).Delete(
			&model.TriggerMemberLoveLevelUp{})
		utils.CheckErr(err)
	}
}

func (session *Session) GetAllTriggerMemberLoveLevelUps() generic.ObjectByObjectIDWrite[*model.TriggerMemberLoveLevelUp] {
	triggers := generic.ObjectByObjectIDWrite[*model.TriggerMemberLoveLevelUp]{}
	err := Engine.Table("s_user_trigger_member_love_level_up").
		Where("user_id = ?", session.UserStatus.UserID).Find(&triggers.Objects)
	utils.CheckErr(err)
	triggers.Length = len(triggers.Objects)
	return triggers
}