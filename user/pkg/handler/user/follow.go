package users

import (
	"net/http"

	"github.com/RipulHandoo/blogx/user/db/database"
	"github.com/RipulHandoo/blogx/user/pkg"
	"github.com/RipulHandoo/blogx/user/pkg/utils"
	"github.com/google/uuid"
)

func FollowUser(w http.ResponseWriter, req *http.Request, user database.User) {
	uuid_param := req.URL.Query().Get("toFollowId")
	uuid, paramParseErr := uuid.Parse(uuid_param)
	if paramParseErr != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, paramParseErr.Error())
		return
	}
	apiConfig := pkg.DbClient
	userFollowTuple, followerUpdateErr := apiConfig.FollowUser(req.Context(), database.FollowUserParams{
		FollowingID: uuid,
		FollowerID:  user.ID,
	})
	if followerUpdateErr != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, followerUpdateErr.Error())
		return
	}
	utils.ResponseJSON(w, http.StatusOK, userFollowTuple)
}