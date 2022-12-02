package serializer

import "SK-todo/model"

type Team struct {
	ID       uint   `json:"id" form:"id" example:"1"`
	TeamName string `json:"team_name" form:"team_name" example:"Sleekflow"`
	CreateAt int64  `json:"create_at" form:"create_at"`
}

func BuildTeam(team model.Team) Team {
	return Team{
		ID:       team.ID,
		TeamName: team.TeamName,
		CreateAt: team.CreatedAt.Unix(),
	}
}

func BuildTeams(teams []model.Team) (items []Team) {
	for _, team := range teams {
		item := BuildTeam(team)
		items = append(items, item)
	}
	return items
}
