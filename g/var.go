package g

import (
	"github.com/barryz/rmqmonitor/utils"
)

// StatsDB stats management database
type StatsDB struct {
	CurrentLocate   string `json:"current_locate"`
	PreviousLocate  string `json:"previous_locate"`
	LastChangeTime  string `json:"last_change"`
	LastCollectTime string `json:"last_collect"`
}

// NewStatsDB create an new stats management database cache
func NewStatsDB() *StatsDB {
	return &StatsDB{}
}

// SetCurrentLocate setting current database location
func (s *StatsDB) SetCurrentLocate(locate string) {
	if s.CurrentLocate != locate {
		s.PreviousLocate = s.CurrentLocate
		s.CurrentLocate = locate
		s.LastChangeTime = utils.GetCurrentDateTime()
	} else {
		// do nothing
	}

	s.LastCollectTime = utils.GetCurrentDateTime()

}
