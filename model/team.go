package model

type TeamParams struct {
	Season int `form:"season" binding:"omitempty,gte=2008,lte=9999"`
	League int `form:"league" binding:"omitempty,gte=1"`
}

type Team struct {
	Id int `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Code string `json:"code"`
	Country string `json:"country"`
	National bool `json:"national"`
	LeagueSeason *TeamLeagueSeason `json:"leagueSeason"`
}

type TeamLeagueSeason struct {
	TeamId int `json:"-" gorm:"primaryKey"`
	LeagueId int `json:"leagueId" gorm:"primaryKey"`
	Season int `json:"season" gorm:"primaryKey"`
	MaxFixtureId int `json:"maxFixtureId"`
	TeamStats *TeamStats `gorm:"foreignKey:TeamId,LeagueId,Season,FixtureId;references:TeamId,LeagueId,Season,MaxFixtureId" json:"maxFixtureStats"`
}