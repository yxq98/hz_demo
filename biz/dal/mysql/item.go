package mysql

import (
	"context"

	"github.com/yxq98/hz_demo/biz/dal/mysql/model"
)

type ItemDAO interface {
	GetItemByID(ctx context.Context, id int64) (*model.Item, error)
}

func NewItemDAO(ctx context.Context) ItemDAO {
	return &itemDAOImpl{
		ms: GetMysql(),
	}
}

type itemDAOImpl struct {
	ms *Mysql
}

func (i *itemDAOImpl) GetItemByID(ctx context.Context, id int64) (*model.Item, error) {
	const query = `SELECT id, name FROM item WHERE id = ?`
	var item model.Item
	err := i.ms.r.GetContext(ctx, &item, query, id)
	return checkErr(&item, err)
}
