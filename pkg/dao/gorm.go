package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ColumnName string

type Column struct {
	ColumnName  ColumnName
	ColumnValue interface{}
}

type Eq Column
type Neq Column
type Gt Column
type Gte Column
type Lt Column
type Lte Column
type In Column
type Like Column

type OrderBy struct {
	ColumnName ColumnName
	Desc       *bool
}

type Sorter struct {
	Columns []OrderBy
}

func (sorter *Sorter) BuildStatement(tx *gorm.DB) {
	if sorter == nil {
		return
	}
	for _, order := range sorter.Columns {
		if order.Desc == nil {
			desc := true
			order.Desc = &desc
		}
		tx.Order(clause.OrderByColumn{Column: clause.Column{Name: string(order.ColumnName)}, Desc: *order.Desc})
	}
}

type Pagination struct {
	CurrentPage int64
	PageSize    int64
}

func (pagination Pagination) setDefault() Pagination {
	if pagination.CurrentPage < 0 {
		pagination.CurrentPage = 1
	}
	if pagination.PageSize > 1000 {
		pagination.PageSize = 1000
	} else if pagination.PageSize < 0 {
		pagination.PageSize = 10
	}
	return pagination
}

func (pagination Pagination) BuildStatement(tx *gorm.DB) Pagination {
	res := pagination.setDefault()
	offset := (pagination.CurrentPage - 1) * pagination.PageSize
	tx.Offset(int(offset)).Limit(int(pagination.PageSize))
	return res
}

type Filter struct {
	Eq []Eq
	// not in
	Neq  []Neq
	Gt   []Gt
	Gte  []Gte
	Lt   []Lt
	Lte  []Lte
	In   []In
	Like []Like
}

func (filter Filter) BuildStatement(tx *gorm.DB) {
	for _, eq := range filter.Eq {
		tx.Where(string(eq.ColumnName)+"= ?", eq.ColumnValue)
	}
	for _, gt := range filter.Gt {
		tx.Where(string(gt.ColumnName)+"> ?", gt.ColumnValue)
	}
	for _, gte := range filter.Gte {
		tx.Where(string(gte.ColumnName)+">= ?", gte.ColumnValue)
	}
	for _, in := range filter.In {
		tx.Where(string(in.ColumnName)+"in ?", in.ColumnValue)
	}
	for _, like := range filter.Like {
		tx.Where(string(like.ColumnName)+"like ?", like.ColumnValue)
	}
	for _, lt := range filter.Lt {
		tx.Where(string(lt.ColumnName)+"< ?", lt.ColumnValue)
	}
	for _, lte := range filter.Lte {
		tx.Where(string(lte.ColumnName)+"<= ?", lte.ColumnValue)
	}
	for _, neq := range filter.Neq {
		tx.Where(string(neq.ColumnName)+"not in ?", neq.ColumnValue)
	}
}

type Query struct {
	Filter     Filter
	Sorter     *Sorter
	Pagination Pagination
}

type Page[T any] struct {
	CurrentPage int64
	PageSize    int64
	List        []*T
	Total       int64
}

type IDAO[T any] interface {
	Create(ctx context.Context, data ...*T) (int64, error)
	Upsert(ctx context.Context, updateColumNames []string, hitColumnNames []string, data ...*T) (int64, error)
	FindById(ctx context.Context, id int64) (*T, error)
	FindByIds(ctx context.Context, ids []int64) ([]*T, error)
	Find(ctx context.Context, query Query) ([]*T, error)
	FindPage(ctx context.Context, query Query) (Page[T], error)
	Count(ctx context.Context, filter Filter) (int64, error)
	UpdateByModel(ctx context.Context, id int64, data *T) (int64, error)
	UpdateByMap(ctx context.Context, id int64, data map[string]interface{}) (int64, error)
	DeleteById(ctx context.Context, ids ...int64) (int64, error)
}

type DAO[T any] struct {
	db    *gorm.DB
	model T
}

func NewDAO[T any](db *gorm.DB) *DAO[T] {
	return &DAO[T]{
		db: db,
	}
}

func (dao *DAO[T]) Create(ctx context.Context, data ...*T) (int64, error) {
	tx := dao.db.Create(data)
	return tx.RowsAffected, tx.Error
}

// 注意要有唯一约束的hit字段才有效果，一般使用主键字段作为hit字段即可
// 若实体有autoUpdateTime，需要将此字段加入updateColumnName中
func (dao *DAO[T]) Upsert(ctx context.Context, updateColumnNames []string, hitColumnNames []string, data ...*T) (int64, error) {
	tx := dao.db.Model(&dao.model)
	columns := make([]clause.Column, 0)
	for _, name := range hitColumnNames {
		columns = append(columns, clause.Column{
			Name: name,
		})
	}

	tx.Clauses(clause.OnConflict{
		Columns:   columns,
		DoUpdates: clause.AssignmentColumns(updateColumnNames),
	}).Create(data)

	return tx.RowsAffected, tx.Error
}

func (dao *DAO[T]) FindById(ctx context.Context, id int64) (*T, error) {
	tx := dao.db.Model(&dao.model)
	res := new(T)
	if err := tx.Take(res, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (dao *DAO[T]) FindByIds(ctx context.Context, ids ...int64) ([]*T, error) {
	tx := dao.db.Model(&dao.model)
	res := make([]*T, 0)
	tx.Find(&res, "id in ?", ids)
	return res, tx.Error
}

func (dao *DAO[T]) Find(ctx context.Context, filter Filter, sorter *Sorter) ([]*T, error) {
	tx := dao.db.Model(&dao.model)
	filter.BuildStatement(tx)
	sorter.BuildStatement(tx)

	res := make([]*T, 0)
	if err := tx.Debug().Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (dao *DAO[T]) FindPage(ctx context.Context, query Query, count bool) (*Page[T], error) {
	tx := dao.db.Model(&dao.model)
	query.Filter.BuildStatement(tx)

	var total int64
	if count {
		if err := tx.Count(&total).Error; err != nil {
			return nil, err
		}
	}

	query.Sorter.BuildStatement(tx)
	pagination := query.Pagination.BuildStatement(tx)
	list := make([]*T, 0)
	if err := tx.Debug().Find(&list).Error; err != nil {
		return nil, err
	}

	return &Page[T]{
		CurrentPage: pagination.CurrentPage,
		PageSize:    pagination.PageSize,
		List:        list,
		Total:       total,
	}, nil
}

func (dao DAO[T]) Count(ctx context.Context, filter Filter) (int64, error) {
	tx := dao.db.Model(&dao.model)
	filter.BuildStatement(tx)

	var total int64
	if err := tx.Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (dao DAO[T]) UpdateByModel(ctx context.Context, id int64, data *T) (int64, error) {
	tx := dao.db.Model(&dao.model)
	tx.Where("id = ?", id).Updates(data)
	return tx.RowsAffected, tx.Error
}

func (dao DAO[T]) UpdateByColumnMap(ctx context.Context, id int64, columnMap map[string]interface{}) (int64, error) {
	tx := dao.db.Model(&dao.model)
	tx.Where("id = ?", id).Updates(columnMap)
	return tx.RowsAffected, tx.Error
}

func (dao DAO[T]) DeleteByIds(ctx context.Context, ids ...int64) (int64, error) {
	m := &dao.model
	dest := make([]*T, 0)
	tx := dao.db.Model(m)

	tx.Delete(&dest, ids)
	return tx.RowsAffected, tx.Error

}
func (dao DAO[T]) Native() *gorm.DB {
	return dao.db
}
