package dao

import (
	"context"
	"github.com/bilibili/kratos/pkg/cache/memcache"
	"shiji_server/internal/model"
	"testing"
)

func TestDaoStruct_AddUser(t *testing.T) {
	type fields struct {
		db       *sql.DB
		mc       *memcache.Memcache
		mcExpire int32
	}
	type args struct {
		c    context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DaoStruct{
				db:       tt.fields.db,
				mc:       tt.fields.mc,
				mcExpire: tt.fields.mcExpire,
			}
			if err := d.AddUser(tt.args.c, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDaoStruct_GetUser(t *testing.T) {
	type fields struct {
		db       *sql.DB
		mc       *memcache.Memcache
		mcExpire int32
	}
	type args struct {
		c    context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRe  string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DaoStruct{
				db:       tt.fields.db,
				mc:       tt.fields.mc,
				mcExpire: tt.fields.mcExpire,
			}
			gotRe, err := d.GetUser(tt.args.c, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRe != tt.wantRe {
				t.Errorf("GetUser() gotRe = %v, want %v", gotRe, tt.wantRe)
			}
		})
	}
}