// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "name", Type: field.TypeString},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:        "categories",
		Columns:     CategoriesColumns,
		PrimaryKey:  []*schema.Column{CategoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ClubsColumns holds the columns for the "clubs" table.
	ClubsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "summary", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Size: 1000},
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
	// CategoryClubsColumns holds the columns for the "category_clubs" table.
	CategoryClubsColumns = []*schema.Column{
		{Name: "category_id", Type: field.TypeString},
		{Name: "club_id", Type: field.TypeInt},
	}
	// CategoryClubsTable holds the schema information for the "category_clubs" table.
	CategoryClubsTable = &schema.Table{
		Name:       "category_clubs",
		Columns:    CategoryClubsColumns,
		PrimaryKey: []*schema.Column{CategoryClubsColumns[0], CategoryClubsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "category_clubs_category_id",
				Columns:    []*schema.Column{CategoryClubsColumns[0]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "category_clubs_club_id",
				Columns:    []*schema.Column{CategoryClubsColumns[1]},
				RefColumns: []*schema.Column{ClubsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CategoriesTable,
		ClubsTable,
		LikeClubsTable,
		CategoryClubsTable,
	}
)

func init() {
	LikeClubsTable.ForeignKeys[0].RefTable = ClubsTable
	CategoryClubsTable.ForeignKeys[0].RefTable = CategoriesTable
	CategoryClubsTable.ForeignKeys[1].RefTable = ClubsTable
}
