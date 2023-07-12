package model

type LeagueParams struct {
	Season int `form:"season" binding:"omitempty,gte=2008,lte=9999"`
}

type League struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Type string `json:"type" gorm:"type:enum('league', 'cup');default:'league'"`
	Country LeagueCountry  `json:"country" gorm:"embedded;embeddedPrefix:country_"`
	Season LeagueSeason `json:"season"`
	Audit   `json:"-"`
}

type LeagueCountry struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type LeagueSeason struct {
	LeagueId int       `json:"-" gorm:"primaryKey"`
	Season   int       `json:"year" gorm:"primaryKey"`
	Start    CivilTime `json:"start"`
	End      CivilTime `json:"end"`
	Current  bool      `json:"current"`
	Audit    `json:"-"`
}
