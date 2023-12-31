// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gormgen

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/lifezq/tickets/internal/model"
)

func newTicket(db *gorm.DB, opts ...gen.DOOption) ticket {
	_ticket := ticket{}

	_ticket.ticketDo.UseDB(db, opts...)
	_ticket.ticketDo.UseModel(&model.Ticket{})

	tableName := _ticket.ticketDo.TableName()
	_ticket.ALL = field.NewAsterisk(tableName)
	_ticket.ID = field.NewInt32(tableName, "id")
	_ticket.Numbers = field.NewString(tableName, "numbers")
	_ticket.Ymd = field.NewString(tableName, "ymd")

	_ticket.fillFieldMap()

	return _ticket
}

type ticket struct {
	ticketDo

	ALL     field.Asterisk
	ID      field.Int32
	Numbers field.String
	Ymd     field.String

	fieldMap map[string]field.Expr
}

func (t ticket) Table(newTableName string) *ticket {
	t.ticketDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t ticket) As(alias string) *ticket {
	t.ticketDo.DO = *(t.ticketDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *ticket) updateTableName(table string) *ticket {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt32(table, "id")
	t.Numbers = field.NewString(table, "numbers")
	t.Ymd = field.NewString(table, "ymd")

	t.fillFieldMap()

	return t
}

func (t *ticket) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *ticket) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 3)
	t.fieldMap["id"] = t.ID
	t.fieldMap["numbers"] = t.Numbers
	t.fieldMap["ymd"] = t.Ymd
}

func (t ticket) clone(db *gorm.DB) ticket {
	t.ticketDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t ticket) replaceDB(db *gorm.DB) ticket {
	t.ticketDo.ReplaceDB(db)
	return t
}

type ticketDo struct{ gen.DO }

type ITicketDo interface {
	gen.SubQuery
	Debug() ITicketDo
	WithContext(ctx context.Context) ITicketDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITicketDo
	WriteDB() ITicketDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITicketDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITicketDo
	Not(conds ...gen.Condition) ITicketDo
	Or(conds ...gen.Condition) ITicketDo
	Select(conds ...field.Expr) ITicketDo
	Where(conds ...gen.Condition) ITicketDo
	Order(conds ...field.Expr) ITicketDo
	Distinct(cols ...field.Expr) ITicketDo
	Omit(cols ...field.Expr) ITicketDo
	Join(table schema.Tabler, on ...field.Expr) ITicketDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITicketDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITicketDo
	Group(cols ...field.Expr) ITicketDo
	Having(conds ...gen.Condition) ITicketDo
	Limit(limit int) ITicketDo
	Offset(offset int) ITicketDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITicketDo
	Unscoped() ITicketDo
	Create(values ...*model.Ticket) error
	CreateInBatches(values []*model.Ticket, batchSize int) error
	Save(values ...*model.Ticket) error
	First() (*model.Ticket, error)
	Take() (*model.Ticket, error)
	Last() (*model.Ticket, error)
	Find() ([]*model.Ticket, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Ticket, err error)
	FindInBatches(result *[]*model.Ticket, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Ticket) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITicketDo
	Assign(attrs ...field.AssignExpr) ITicketDo
	Joins(fields ...field.RelationField) ITicketDo
	Preload(fields ...field.RelationField) ITicketDo
	FirstOrInit() (*model.Ticket, error)
	FirstOrCreate() (*model.Ticket, error)
	FindByPage(offset int, limit int) (result []*model.Ticket, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITicketDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FindByID(id int32) (result *model.Ticket, err error)
}

// where("`id`=@id")
func (t ticketDo) FindByID(id int32) (result *model.Ticket, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("`id`=? ")

	var executeSQL *gorm.DB
	executeSQL = t.UnderlyingDB().Where(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (t ticketDo) Debug() ITicketDo {
	return t.withDO(t.DO.Debug())
}

func (t ticketDo) WithContext(ctx context.Context) ITicketDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t ticketDo) ReadDB() ITicketDo {
	return t.Clauses(dbresolver.Read)
}

func (t ticketDo) WriteDB() ITicketDo {
	return t.Clauses(dbresolver.Write)
}

func (t ticketDo) Session(config *gorm.Session) ITicketDo {
	return t.withDO(t.DO.Session(config))
}

func (t ticketDo) Clauses(conds ...clause.Expression) ITicketDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t ticketDo) Returning(value interface{}, columns ...string) ITicketDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t ticketDo) Not(conds ...gen.Condition) ITicketDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t ticketDo) Or(conds ...gen.Condition) ITicketDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t ticketDo) Select(conds ...field.Expr) ITicketDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t ticketDo) Where(conds ...gen.Condition) ITicketDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t ticketDo) Order(conds ...field.Expr) ITicketDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t ticketDo) Distinct(cols ...field.Expr) ITicketDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t ticketDo) Omit(cols ...field.Expr) ITicketDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t ticketDo) Join(table schema.Tabler, on ...field.Expr) ITicketDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t ticketDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITicketDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t ticketDo) RightJoin(table schema.Tabler, on ...field.Expr) ITicketDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t ticketDo) Group(cols ...field.Expr) ITicketDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t ticketDo) Having(conds ...gen.Condition) ITicketDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t ticketDo) Limit(limit int) ITicketDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t ticketDo) Offset(offset int) ITicketDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t ticketDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITicketDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t ticketDo) Unscoped() ITicketDo {
	return t.withDO(t.DO.Unscoped())
}

func (t ticketDo) Create(values ...*model.Ticket) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t ticketDo) CreateInBatches(values []*model.Ticket, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t ticketDo) Save(values ...*model.Ticket) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t ticketDo) First() (*model.Ticket, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Ticket), nil
	}
}

func (t ticketDo) Take() (*model.Ticket, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Ticket), nil
	}
}

func (t ticketDo) Last() (*model.Ticket, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Ticket), nil
	}
}

func (t ticketDo) Find() ([]*model.Ticket, error) {
	result, err := t.DO.Find()
	return result.([]*model.Ticket), err
}

func (t ticketDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Ticket, err error) {
	buf := make([]*model.Ticket, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t ticketDo) FindInBatches(result *[]*model.Ticket, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t ticketDo) Attrs(attrs ...field.AssignExpr) ITicketDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t ticketDo) Assign(attrs ...field.AssignExpr) ITicketDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t ticketDo) Joins(fields ...field.RelationField) ITicketDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t ticketDo) Preload(fields ...field.RelationField) ITicketDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t ticketDo) FirstOrInit() (*model.Ticket, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Ticket), nil
	}
}

func (t ticketDo) FirstOrCreate() (*model.Ticket, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Ticket), nil
	}
}

func (t ticketDo) FindByPage(offset int, limit int) (result []*model.Ticket, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t ticketDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t ticketDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t ticketDo) Delete(models ...*model.Ticket) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *ticketDo) withDO(do gen.Dao) *ticketDo {
	t.DO = *do.(*gen.DO)
	return t
}
