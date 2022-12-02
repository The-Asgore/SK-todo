package operation

import (
	"SK-todo/model"
	respCode "SK-todo/resp-code"
	"SK-todo/serializer"
	"strconv"
	"time"
)

type ShowTaskOP struct {
	UserID uint
	TeamID string
}

type DeleteTaskOP struct {
}

type UpdateTaskOP struct {
	Name        string `form:"name" json:"name" binding:"required,min=2,max=100"`
	Description string `form:"description" json:"description" binding:"max=1000"`
	Status      int    `form:"status" json:"status"`
}

type CreateTaskOP struct {
	Name        string `form:"name" json:"name" binding:"required,min=2,max=100"`
	Description string `form:"description" json:"description" binding:"max=1000"`
	Status      int    `form:"status" json:"status"`
	DueDate     int64  `form:"due_date" json:"due_date"`
}

type SearchTaskOP struct {
	Info string `form:"info" json:"info"`
}

type ListTasksOP struct {
	Index  string `form:"index" json:"index"`
	Order  string `form:"order" json:"order"`
	Limit  string `form:"limit" json:"limit"`
	Start  string `form:"start" json:"start"`
	Status string `form:"status" json:"status"`
}

func CreateTask(task model.Task) serializer.Response {
	code := respCode.SUCCESS
	err := model.DB.Create(&task).Error
	if err != nil {
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    respCode.GetMsg(code),
	}
}

func (op *CreateTaskOP) CreatePersonalTask(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	task := model.Task{
		User:        &user,
		UserID:      user.ID,
		Owner:       0,
		Name:        op.Name,
		Description: op.Description,
		Status:      op.Status,
		StartTime:   time.Now().Unix(),
		DueDate:     op.DueDate,
	}
	return CreateTask(task)
}

func (op *CreateTaskOP) CreateTeamTask(teamID string) serializer.Response {
	var team model.Team
	model.DB.First(&team, teamID)
	task := model.Task{
		Team:        &team,
		TeamID:      team.ID,
		Owner:       1,
		Name:        op.Name,
		Description: op.Description,
		Status:      op.Status,
		StartTime:   time.Now().Unix(),
		DueDate:     op.DueDate,
	}
	return CreateTask(task)
}

func (op *ListTasksOP) ListPersonalTask(id uint) serializer.Response {
	var tasks []model.Task
	var total int64
	limitNum, _ := strconv.Atoi(op.Limit)
	startNum, _ := strconv.Atoi(op.Start)
	if op.Status != "1" && op.Status != "0" {
		model.DB.Model(model.Task{}).Preload("User").Where("user_id = ? and owner=0", id).Count(&total).Order(op.Index + " " + op.Order).
			Limit(limitNum).Offset((startNum - 1) * limitNum).
			Find(&tasks)
	} else {
		model.DB.Model(model.Task{}).Preload("User").Where("user_id = ? and status = ? and owner=0", id, op.Status).Count(&total).Order(op.Index + " " + op.Order).
			Limit(limitNum).Offset((startNum - 1) * limitNum).
			Find(&tasks)
	}

	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(total))
}

func (op *ListTasksOP) ListTeamTask(teamID string) serializer.Response {
	var tasks []model.Task
	var total int64
	limitNum, _ := strconv.Atoi(op.Limit)
	startNum, _ := strconv.Atoi(op.Start)
	if op.Status != "1" && op.Status != "0" {
		model.DB.Model(model.Task{}).Preload("Team").Where("team_id = ? and owner=1", teamID).Count(&total).Order(op.Index + " " + op.Order).
			Limit(limitNum).Offset((startNum - 1) * limitNum).
			Find(&tasks)
	} else {
		model.DB.Model(model.Task{}).Preload("Team").Where("team_id = ? and status = ? and owner=1", teamID, op.Status).Count(&total).Order(op.Index + " " + op.Order).
			Limit(limitNum).Offset((startNum - 1) * limitNum).
			Find(&tasks)
	}

	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(total))
}

func (op *ShowTaskOP) ShowPersonalTask(id string) serializer.Response {
	var task model.Task
	code := respCode.SUCCESS
	err := model.DB.Model(&task).Where("user_id = ?", op.UserID).First(&task, id).Error
	if err != nil {
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    respCode.GetMsg(code),
	}
}

func (op *ShowTaskOP) ShowTeamTask(id string) serializer.Response {
	var task model.Task
	code := respCode.SUCCESS
	err := model.DB.Model(&task).Where("team_id = ?", op.TeamID).First(&task, id).Error
	if err != nil {
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    respCode.GetMsg(code),
	}
}

func DeleteTask(task model.Task) serializer.Response {
	code := respCode.SUCCESS
	err := model.DB.Delete(&task).Error
	if err != nil {
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    respCode.GetMsg(code),
	}
}

func (op *DeleteTaskOP) DeletePersonalTask(id string, operatorID uint) serializer.Response {
	var task model.Task
	code := respCode.SUCCESS
	err := model.DB.First(&task, id).Where("user_id = ? and owner=0", operatorID).Error
	if err != nil {
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return DeleteTask(task)
}

func (op *DeleteTaskOP) DeleteTeamTask(id string, teamID string) serializer.Response {
	var task model.Task
	code := respCode.SUCCESS
	err := model.DB.First(&task, id).Where("team_id = ? and owner=1", teamID).Error
	if err != nil {
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return DeleteTask(task)
}

func UpdateTask(task model.Task) serializer.Response {
	code := respCode.SUCCESS
	err := model.DB.Save(&task).Error
	if err != nil {
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    respCode.GetMsg(code),
		Data:   "Success",
	}
}

func (op *UpdateTaskOP) UpdatePersonalTask(id string, operatorID uint) serializer.Response {
	var task model.Task
	model.DB.Model(model.Task{}).Where("id = ? and user_id = ? and owner=0", id, operatorID).First(&task)
	task.Description = op.Description
	task.Status = op.Status
	task.Name = op.Name
	return UpdateTask(task)
}

func (op *UpdateTaskOP) UpdateTeamTask(id string, teamID string) serializer.Response {
	var task model.Task
	model.DB.Model(model.Task{}).Where("id = ? and team_id = ? and owner=1", id, teamID).First(&task)
	task.Description = op.Description
	task.Status = op.Status
	task.Name = op.Name
	return UpdateTask(task)
}

func (op *SearchTaskOP) SearchPersonalTask(id uint) serializer.Response {
	var tasks []model.Task
	code := respCode.SUCCESS
	//model.DB.Where("user_id=?", id).Preload("User").First(&tasks)
	err := model.DB.Model(&tasks).Where("name LIKE ? OR description LIKE ?",
		"%"+op.Info+"%", "%"+op.Info+"%").Where("user_id = ? and owner=0", id).Find(&tasks).Error
	if err != nil {
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    respCode.GetMsg(code),
		Data:   serializer.BuildTasks(tasks),
	}
}

func (op *SearchTaskOP) SearchTeamTask(teamID string) serializer.Response {
	var tasks []model.Task
	code := respCode.SUCCESS
	//model.DB.Where("team_id=?", teamID).Preload("Team").First(&tasks)
	err := model.DB.Model(&tasks).Where("name LIKE ? OR description LIKE ?",
		"%"+op.Info+"%", "%"+op.Info+"%").Where("team_id = ? and owner=1", teamID).Find(&tasks).Error
	if err != nil {
		code = respCode.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    respCode.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    respCode.GetMsg(code),
		Data:   serializer.BuildTasks(tasks),
	}
}
