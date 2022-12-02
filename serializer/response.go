package serializer

type ResponseUser struct {
	Status int    `json:"status"`
	Data   User   `json:"data"`
	Msg    string `json:"msg"`
	Error  string `json:"error"`
}

type ResponseUserTeams struct {
	Status int    `json:"status"`
	Data   []Team `json:"data"`
	Msg    string `json:"msg"`
	Error  string `json:"error"`
}

type ResponseTask struct {
	Status int    `json:"status"`
	Data   Task   `json:"data"`
	Msg    string `json:"msg"`
	Error  string `json:"error"`
}

type ResponseTeam struct {
	Status int    `json:"status"`
	Data   Team   `json:"data"`
	Msg    string `json:"msg"`
	Error  string `json:"error"`
}

type ResponseTeamMembers struct {
	Status int    `json:"status"`
	Data   []User `json:"data"`
	Msg    string `json:"msg"`
	Error  string `json:"error"`
}
