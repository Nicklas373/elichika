package userdata

import (
	"elichika/utils"
)

func liveMvFinalizer(session *Session) {
	for _, userLiveMv := range session.UserModel.UserLiveMvDeckByID.Objects {
		affected, err := session.Db.Table("u_live_mv_deck").Where("user_id = ? AND live_master_id = ?",
			session.UserStatus.UserID, userLiveMv.LiveMasterID).AllCols().Update(userLiveMv)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_live_mv_deck").Insert(userLiveMv)
			utils.CheckErr(err)
		}
	}
	for _, userLiveMvCustom := range session.UserModel.UserLiveMvDeckCustomByID.Objects {
		affected, err := session.Db.Table("u_live_mv_deck_custom").Where("user_id = ? AND live_master_id = ?",
			session.UserStatus.UserID, userLiveMvCustom.LiveMasterID).AllCols().Update(userLiveMvCustom)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_live_mv_deck_custom").Insert(userLiveMvCustom)
			utils.CheckErr(err)
		}
	}
}
func init() {
	addFinalizer(liveMvFinalizer)
	addGenericTableFieldPopulator("u_live_mv_deck", "UserLiveMvDeckByID")
	addGenericTableFieldPopulator("u_live_mv_deck_custom", "UserLiveMvDeckCustomByID")
}