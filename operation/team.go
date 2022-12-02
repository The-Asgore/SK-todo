package operation

import (
	"SK-todo/model"
	respCode "SK-todo/resp-code"
	"SK-todo/serializer"
	"fmt"
)

type CreateTeamOP struct {
	TeamName string `form:"team_name" json:"team_name" binding:"required,min=3,max=15" example:"sleekflow"`
}

type UpdateTeamOP struct {
	TeamName string `form:"team_name" json:"team_name" binding:"min=3,max=15" example:"sleekflow"`
	User     []uint `form:"user" json:"user" example:"1,2,3"`
}

type DeleteTeamOP struct {
}

type ShowTeamMemberOP struct {
}

func (op *CreateTeamOP) CreateTeam() *serializer.Response {
	code := respCode.SUCCESS
	var team model.Team
	var count int64
	model.DB.Model(&model.Team{}).Where("team_name=?", op.TeamName).First(&team).Count(&count)
	if count == 1 {
		code = respCode.ErrorExistTeam
		return &serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}
	team.TeamName = op.TeamName

	if err := model.DB.Create(&team).Error; err != nil {
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

func (op *UpdateTeamOP) UpdateTeam(id string) *serializer.Response {
	code := respCode.SUCCESS
	var team model.Team
	var count int64
	model.DB.Model(&model.Team{}).Where("id=?", id).First(&team).Count(&count)
	if count != 1 {
		code = respCode.ErrorNotExistTeam
		return &serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}

	if op.TeamName != "" {
		model.DB.Model(&team).Update("team_name", op.TeamName)
	}

	var users []*model.User
	for _, uid := range op.User {
		var user model.User
		var count int64
		model.DB.Model(&model.User{}).Where("id=?", uid).First(&user).Count(&count)
		if count != 1 {
			code = respCode.ErrorNotExistUser
			return &serializer.Response{
				Status: code,
				Msg:    respCode.GetMsg(code),
			}
		}
		users = append(users, &user)
	}
	fmt.Println(users[0].ID)
	fmt.Println(users[1].ID)

	if op.User != nil {
		if err := model.DB.Model(&team).Association("User").Replace(users); err != nil {
			code = respCode.ErrorDatabase
			return &serializer.Response{
				Status: code,
				Msg:    respCode.GetMsg(code),
			}
		}
	}

	return &serializer.Response{
		Status: code,
		Msg:    respCode.GetMsg(code),
	}
}

func (op *DeleteTeamOP) DeleteTeam(id string) *serializer.Response {
	code := respCode.SUCCESS
	var team model.Team
	var count int64
	model.DB.Model(&model.Team{}).Where("id=?", id).First(&team).Count(&count)
	if count != 1 {
		code = respCode.ErrorNotExistTeam
		return &serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}

	if err := model.DB.Delete(&team).Error; err != nil {
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

func (op *ShowTeamMemberOP) ShowTeamMember(id string) *serializer.Response {
	code := respCode.SUCCESS
	var team model.Team
	var count int64
	model.DB.Model(&model.Team{}).Where("id=?", id).First(&team).Count(&count)
	if count != 1 {
		code = respCode.ErrorNotExistTeam
		return &serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}

	var users []model.User
	if err := model.DB.Model(&team).Association("User").Find(&users); err != nil {
		code = respCode.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
		}
	}

	return &serializer.Response{
		Status: code,
		Msg:    respCode.GetMsg(code),
		Data:   serializer.BuildUsers(users),
	}
}
