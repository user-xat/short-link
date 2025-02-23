package stat

import (
	"time"

	"github.com/user-xat/short-link/pkg/db"

	"gorm.io/datatypes"
)

type StatsBy struct {
	Period string `json:"period"`
	Sum    int    `json:"sum"`
}

type StatRepository struct {
	Database *db.Db
}

func NewStatRepository(database *db.Db) *StatRepository {
	return &StatRepository{
		Database: database,
	}
}

func (repo *StatRepository) AddClick(linkId uint) {
	var stat Stat
	currentDate := datatypes.Date(time.Now())
	repo.Database.Find(&stat, "link_id = ? and date = ?", linkId, currentDate)
	if stat.ID == 0 {
		repo.Database.DB.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks++
		repo.Database.DB.Save(&stat)
	}
}

func (repo *StatRepository) Get(by string, from, to time.Time) []StatsBy {
	var stats []StatsBy
	var selectQuery string
	switch by {
	case GroupByDay:
		selectQuery = "to_char(date, 'YYYY-MM-DD') as period, sum(clicks)"
	case GroupByMonth:
		selectQuery = "to_char(date, 'YYYY-MM') as period, sum(clicks)"
	}
	repo.Database.Table("stats").
		Select(selectQuery).
		Where("deleted_at is null and date between ? and ?", from, to).
		Group("period").
		Order("period").
		Scan(&stats)
	return stats
}
