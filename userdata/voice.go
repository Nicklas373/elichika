package userdata

import (
	"elichika/client"
	"elichika/utils"
)

// TODO(refactor): Move into subsystem
func (session *Session) UpdateVoice(naviVoiceMasterId int32, isNew bool) {
	userVoice := client.UserVoice{}
	exist, err := session.Db.Table("u_voice").Where("user_id = ? AND navi_voice_master_id = ?",
		session.UserId, naviVoiceMasterId).Get(&userVoice)
	utils.CheckErr(err)
	if exist {
		if userVoice.IsNew == isNew {
			return
		}
		userVoice.IsNew = isNew
	} else {
		userVoice = client.UserVoice{
			NaviVoiceMasterId: naviVoiceMasterId,
			IsNew:             isNew,
		}
	}
	session.UserModel.UserVoiceByVoiceId.Set(naviVoiceMasterId, userVoice)
}
func voiceFinalizer(session *Session) {
	for _, userVoice := range session.UserModel.UserVoiceByVoiceId.Map {
		affected, err := session.Db.Table("u_voice").Where("user_id = ? AND navi_voice_master_id = ?",
			session.UserId, userVoice.NaviVoiceMasterId).AllCols().Update(*userVoice)
		utils.CheckErr(err)
		if affected == 0 {
			GenericDatabaseInsert(session, "u_voice", *userVoice)
		}
	}
}
func init() {
	AddFinalizer(voiceFinalizer)
}
