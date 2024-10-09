package repository

import (
	"chess/internal/model"
	utiltime "chess/pkg/helper/time"
	"context"
	json "github.com/goccy/go-json"
	"time"
)

type RankRepository interface {
	GetRank(ctx context.Context) ([]*model.RankList, error)
	PostRank(ctx context.Context, params *model.UserRankParam) (bool, error)
}

func NewRankRepository(
	repository *Repository,
) RankRepository {
	return &rankRepository{
		Repository: repository,
	}
}

type rankRepository struct {
	*Repository
}

func (r *rankRepository) GetRank(ctx context.Context) ([]*model.RankList, error) {
	var rankList []*model.RankList
	// 记录列表
	sql := "SELECT region, count(1) total FROM rank where date = ? group by region order by total desc limit 3"
	res := r.db.Raw(sql, time.Now().Format("2006-01-02")).Scan(&rankList)
	if res.Error != nil {
		r.logger.Debug(res.Error.Error())
		return nil, res.Error
	}
	for k := range rankList {
		rankList[k].Rank = k + 1
	}
	return rankList, nil
}

func (r *rankRepository) PostRank(ctx context.Context, params *model.UserRankParam) (bool, error) {
	var rank model.Rank
	rank.AccountName = params.AccountName
	rank.Region = params.Region
	rank.Date = time.Now().Format("2006-01-02")
	rank.CreatedAt = utiltime.GetNowDateTime()
	res := r.db.Create(&rank)
	if res.Error != nil {
		paramsByte, _ := json.Marshal(params)
		r.logger.Debug(res.Error.Error() + ", bot info: " + string(paramsByte))
		return false, res.Error
	}
	return true, nil
}
