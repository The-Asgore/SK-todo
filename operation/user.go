package operation

import (
	"SK-todo/auth"
	"SK-todo/model"
	respCode "SK-todo/resp-code"
	"SK-todo/serializer"
	"gorm.io/gorm"
)

type UserOP struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"admin"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"password"`
}

type UpdateUserOP struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"admin"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"password"`
}

type ShowUserTeamOP struct{}

type DeleteUserOP struct{}

type JoinTeamOP struct{}

type LeaveTeamOP struct{}

func (op *UserOP) Register() *serializer.Response {
	code := respCode.SUCCESS
	var user model.User
	var count int64
	model.DB.Model(&model.User{}).Where("user_name=?", op.UserName).First(&user).Count(&count)
	if count == 1 {
		code = respCode.ErrorExistUser
		return &serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}
	user.UserName = op.UserName

	if err := user.SetPassword(op.Password); err != nil {
		code = respCode.ErrorFailEncryption
		return &serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}

	if err := model.DB.Create(&user).Error; err != nil {
		code = respCode.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Msg:    respCode.GetMsg(code),
	}
}

func (op *UserOP) Login() serializer.Response {
	var user model.User
	code := respCode.SUCCESS
	if err := model.DB.Where("user_name=?", op.UserName).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			code = respCode.ErrorNotExistUser
			return serializer.Response{
				Status: code,
				Msg:    respCode.GetMsg(code),
			}
		}
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}
	if !user.CheckPassword(op.Password) {
		code = respCode.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}
	token, err := auth.GenerateToken(user.ID, op.UserName, 0)
	if err != nil {
		code = respCode.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    respCode.GetMsg(code),
	}
}

func (op *ShowUserTeamOP) ShowTeams(id uint) serializer.Response {
	var user model.User
	var team []model.Team
	code := respCode.SUCCESS
	if err := model.DB.Model(&user).Where("id=?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			code = respCode.ErrorNotExistUser
			return serializer.Response{
				Status: code,
				Msg:    respCode.GetMsg(code),
			}
		}
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}
	if err := model.DB.Model(&user).Association("Teams").Find(&team); err != nil {
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTeams(team),
		Msg:    respCode.GetMsg(code),
	}
}

func (op *UpdateUserOP) UpdateUser(id uint) serializer.Response {
	var user model.User
	code := respCode.SUCCESS
	if err := model.DB.Model(&user).Where("id=?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			code = respCode.ErrorNotExistUser
			return serializer.Response{
				Status: code,
				Msg:    respCode.GetMsg(code),
			}
		}
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}
	if err := model.DB.Model(&user).Where("id=?", id).Update("user_name", op.UserName).Error; err != nil {
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}

	if err := user.SetPassword(op.Password); err != nil {
		code = respCode.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}
	if err := model.DB.Model(&user).Where("id=?", id).Update("password_digest", user.PasswordDigest).Error; err != nil {
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    respCode.GetMsg(code),
	}
}

func (*DeleteUserOP) DeleteUser(id uint) serializer.Response {
	var user model.User
	code := respCode.SUCCESS
	if err := model.DB.Where("id=?", id).Delete(&user).Error; err != nil {
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    respCode.GetMsg(code),
	}
}

func (*JoinTeamOP) JoinTeam(id uint, teamId string) serializer.Response {
	var user model.User
	var team model.Team
	code := respCode.SUCCESS
	if err := model.DB.Model(&user).Where("id=?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			code = respCode.ErrorNotExistUser
			return serializer.Response{
				Status: code,
				Msg:    respCode.GetMsg(code),
			}
		}
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}
	if err := model.DB.Model(&team).Where("id=?", teamId).First(&team).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			code = respCode.ErrorNotExistTeam
			return serializer.Response{
				Status: code,
				Msg:    respCode.GetMsg(code),
			}
		}
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}
	if err := model.DB.Model(&user).Association("Teams").Append(&team); err != nil {
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    respCode.GetMsg(code),
	}
}

func (*LeaveTeamOP) LeaveTeam(id uint, teamId string) serializer.Response {
	var user model.User
	var team model.Team
	code := respCode.SUCCESS
	if err := model.DB.Model(&user).Where("id=?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			code = respCode.ErrorNotExistUser
			return serializer.Response{
				Status: code,
				Msg:    respCode.GetMsg(code),
			}
		}
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}
	if err := model.DB.Model(&team).Where("id=?", teamId).First(&team).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			code = respCode.ErrorNotExistTeam
			return serializer.Response{
				Status: code,
				Msg:    respCode.GetMsg(code),
			}
		}
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}
	if err := model.DB.Model(&user).Association("Teams").Delete(&team); err != nil {
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    respCode.GetMsg(code),
	}
}
