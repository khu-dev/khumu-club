// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ClubsColumns holds the columns for the "clubs" table.
	ClubsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "summary", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Size: 1000},
		{Name: "categories", Type: field.TypeJSON},
		{Name: "images", Type: field.TypeJSON, Nullable: true},
		{Name: "homepage", Type: field.TypeString, Nullable: true},
		{Name: "instagram", Type: field.TypeString, Nullable: true},
		{Name: "facebook", Type: field.TypeString, Nullable: true},
		{Name: "phone", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "recommended", Type: field.TypeBool, Default: false},
	}
	// ClubsTable holds the schema information for the "clubs" table.
	ClubsTable = &schema.Table{
		Name:        "clubs",
		Columns:     ClubsColumns,
		PrimaryKey:  []*schema.Column{ClubsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// LikeClubsColumns holds the columns for the "like_clubs" table.
	LikeClubsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "club_likes", Type: field.TypeInt, Nullable: true},
	}
	// LikeClubsTable holds the schema information for the "like_clubs" table.
	LikeClubsTable = &schema.Table{
		Name:       "like_clubs",
		Columns:    LikeClubsColumns,
		PrimaryKey: []*schema.Column{LikeClubsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "like_clubs_clubs_likes",
				Columns:    []*schema.Column{LikeClubsColumns[3]},
				RefColumns: []*schema.Column{ClubsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ClubsTable,
		LikeClubsTable,
	}
)

func init() {
	LikeClubsTable.ForeignKeys[0].RefTable = ClubsTable
}
