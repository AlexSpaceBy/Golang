package structs

type Group struct{
	Id   string `json:"Id"`
	Name string `json:"name"`
	Uuid string `json:"uuid"`
	Tasks []string `json:"tasks"`
}

type Interval struct{
	Start []string `json:"start"`
	Stop []string `json:"stop"`
}

type Task struct{
	Id    string `json:"Id"`
	Name  string `json:"name"`
	Uuid  string `json:"uuid"`
	Group string `json:"group"`
	Timeframes []Interval `json:"timeframes"`
}

type GroupTask struct{
	Id string `json:"Id"`
	Group_id string `json:"group_id"`
	Task_id string `json:"task_id"`
}

type Timeframe struct{
	Id string `json:"Id"`
	Start string `json:"start"`
	Stop string `json:"stop"`
	Uuid string `json:"uuid"`
	Task string `json:"task"`
}

