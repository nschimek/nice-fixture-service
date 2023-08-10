package model

import "time"

type Fixture struct {
	Id int `gorm:"primaryKey" json:"id"`
	Date time.Time `json:"date"`
	Venue FixtureVenue `gorm:"embedded;embeddedPrefix:venue_"`
	LeagueId int `json:"-"`
	Season int `json:"-"`
	LeagueSeason LeagueSeason `gorm:"foreignKey:LeagueId,Season" json:"leagueSeason"`
	StatusId string `json:"-"`
	Status FixtureStatus `gorm:"foreignKey:Id;references:StatusId" json:"status"`
	Teams FixtureTeams `gorm:"embedded; embeddedPrefix:team_" json:"teams"`
	Goals FixtureGoals `gorm:"embedded; embeddedPrefix:goals_" json:"goals"`
}

type FixtureVenue struct {
	Name string `json:"name"`
	City string `json:"city"`
}

type FixtureTeams struct {
	HomeId int `json:"-"`
	Home Team `gorm:"foreignKey:Id;references:HomeId" json:"home"`
	HomeResult string `json:"homeResult"`
	AwayId int `json:"-"`
	Away Team `gorm:"foreignKey:Id;references:AwayId" json:"away"`
	AwayResult string `json:"awayResult"`
}

type FixtureGoals struct {
	Home int `json:"home"`
	Away int `json:"away"`
}



