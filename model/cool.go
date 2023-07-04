package model

import (
	"database/sql"
	"time"
)

type Cool struct {
	ID					int
	CoolLeaderID		int				`json:"coolLeaderID"`
	CoolLeaderName		string			`json:"coolLeaderName"`
	Name				string			`json:"coolName"`
	MemberCount			int				`json:"memberCount"`
	Address				string			`json:"address"`
	AreaID				int				`json:"areaID"`
	AreaName			string			`json:"areaName"`
	IsActive			bool			`json:"isActive"`
	CreatedAt 			time.Time		`json:"createdAt"`
	UpdatedAt 			time.Time		`json:"updatedAt"`
	DeletedAt			sql.NullTime	`json:"deletedAt"`
}

type CreateCoolRequest struct {
	Name				string			`json:"coolName"`
	CoolLeaderID		int				`json:"coolLeaderID"`
	IsActive			bool			`json:"isActive"`
	Address				string			`json:"address"`
	AreaID				int				`json:"areaID"`
}