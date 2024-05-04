package todo

type User struct {
	Id          int    `json:"-"`
	Name        string `json:"name" validate:"required,min=3,max=20"`
	Username    string `json:"username" validate:"required,min=3,max=20"`
	Email       string `json:"email" validate:"required,min=3,max=20"`
	Password    string `json:"password" validate:"required,min=6"`
	Deleted     bool   `json:"delition"`
	TimeDeleted string `json:"time_delited"`
}

type Conditions struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
