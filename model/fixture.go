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
	TeamHomeId int `json:"-"`
	TeamHome *Team `gorm:"foreignKey:Id;references:TeamHomeId" json:"homeTeam"`
	TeamHomeResult string `json:"homeTeamResult"`
	TeamAwayId int `json:"-"`
	TeamAway *Team `gorm:"foreignKey:Id;references:TeamAwayId" json:"awayTeam"`
	TeamAwayResult string `json:"awayTeamResult"`
	TeamHomeStats *TeamStats `gorm:"foreignKey:TeamId,LeagueId,Season,NextFixtureId;references:TeamHomeId,LeagueId,Season,Id" json:"homeTeamStats"`
	TeamAwayStats *TeamStats `gorm:"foreignKey:TeamId,LeagueId,Season,NextFixtureId;references:TeamAwayId,LeagueId,Season,Id" json:"awayTeamStats"`
	TeamHomeTLS *TeamLeagueSeason `gorm:"foreignKey:TeamId,LeagueId,Season;references:TeamHomeId,LeagueId,Season" json:"homeTeamTLS"`
	TeamAwayTLS *TeamLeagueSeason `gorm:"foreignKey:TeamId,LeagueId,Season;references:TeamAwayId,LeagueId,Season" json:"awayTeamTLS"`
	Goals *FixtureGoals `gorm:"embedded; embeddedPrefix:goals_" json:"goals"`
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



