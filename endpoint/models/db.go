package models

import (
	"EasyGridding/endpoint/log"
	"EasyGridding/endpoint/tool"
	"context"
	"gorm.io/gorm"
	"strings"
	"time"
)

type RecordDBService struct {
	db        *gorm.DB
	tableName string
}

func InitTables(g *gorm.DB) error {
	err := g.Exec(createRecordTable).Error
	if err != nil {
		return err
	}
	return nil
}

func NewDBService(db *gorm.DB, tableName string) *RecordDBService {
	r := &RecordDBService{
		db:        db,
		tableName: tableName,
	}
	InitTables(db)
	return r
}

func (d *RecordDBService) CreateRecord(ctx context.Context, r *Record) error {
	r.CTime = int(time.Now().Unix())
	r.MTime = r.CTime

	g := d.db.Table(d.tableName).Create(r)
	if g.Error != nil {
		log.Log.Printf("CreateRecord failed %+v", g.Error)
		return g.Error
	}
	return nil
}

func (d *RecordDBService) GetRecordList(ctx context.Context, filter []*tool.Filter, offset, limit int) ([]*Record, error) {

	var sqlStr []string
	var val []interface{}
	for _, v := range filter {
		if v.Name == "userId" {
			sqlStr = append(sqlStr, "user_id in (?)")
			val = append(val, v.Values)
		}
		if v.Name == "startTime" {
			sqlStr = append(sqlStr, "ctime >= ? ")
			val = append(val, v.Values)
		}
		if v.Name == "endTime" {
			sqlStr = append(sqlStr, "ctime <= ? ")
			val = append(val, v.Values)
		}
	}

	var data []*Record
	g := d.db.Table(d.tableName).Where(strings.Join(sqlStr, " and "), val...).Offset(offset).Limit(limit).Find(&data)
	if g.Error != nil {
		log.Log.Printf("GetRecordList failed %+v", g.Error)
		return nil, g.Error
	}

	return data, nil
}
