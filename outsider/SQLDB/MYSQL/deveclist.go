package MYSQL

import (
	"context"
	"database/sql"
)

type deviceListdb struct {
	DB *sql.DB
}

type DeviceList struct {
	Id          int    `json:"id"`
	Os          string `json:"os"`
	Ip          int    `json:"ip"`
	Port        int    `json:"port"`
	Servicetype string `json:"type"`
	Sqlac       string `json:"ac"`
	Sqlpw       string `json:"pw"`
}

type DeviceListSql interface {
	SearchById(ctx context.Context, id int64) (DeviceList, error)
	SaveDevice(ctx context.Context, query string, args ...interface{}) (result []DeviceList, err error)
}

type DeviceListuse interface {
	Search(ctx context.Context, id int64) (DeviceList, error)
	SaveDevice(ctx context.Context, query string, args ...interface{}) (result []DeviceList, err error)
}

func NewDeviceList(db *sql.DB) DeviceListSql {
	return &deviceListdb{DB: db}
}

func (m *deviceListdb) SearchById(ctx context.Context, id int64) (DeviceList, error) {
	var devi DeviceList
	var err error
	return devi, err
}

func (m *deviceListdb) SaveDevice(ctx context.Context, query string, args ...interface{}) (result []DeviceList, err error)
