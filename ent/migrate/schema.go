// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "sequence_number", Type: field.TypeInt32},
		{Name: "external_id", Type: field.TypeString},
		{Name: "dst", Type: field.TypeString},
		{Name: "message", Type: field.TypeString},
		{Name: "src", Type: field.TypeString},
		{Name: "state", Type: field.TypeInt},
		{Name: "smsc_message_id", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
		{Name: "provider_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "messages_provides_messages",
				Columns: []*schema.Column{MessagesColumns[10]},

				RefColumns: []*schema.Column{ProvidesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "messages_users_messages",
				Columns: []*schema.Column{MessagesColumns[11]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PricesColumns holds the columns for the "prices" table.
	PricesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "min", Type: field.TypeInt32},
		{Name: "max", Type: field.TypeInt32},
		{Name: "price", Type: field.TypeInt16, Unique: true},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// PricesTable holds the schema information for the "prices" table.
	PricesTable = &schema.Table{
		Name:        "prices",
		Columns:     PricesColumns,
		PrimaryKey:  []*schema.Column{PricesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProvidesColumns holds the columns for the "provides" table.
	ProvidesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "ip_adres", Type: field.TypeString, Unique: true},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// ProvidesTable holds the schema information for the "provides" table.
	ProvidesTable = &schema.Table{
		Name:        "provides",
		Columns:     ProvidesColumns,
		PrimaryKey:  []*schema.Column{ProvidesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// RatesColumns holds the columns for the "rates" table.
	RatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// RatesTable holds the schema information for the "rates" table.
	RatesTable = &schema.Table{
		Name:        "rates",
		Columns:     RatesColumns,
		PrimaryKey:  []*schema.Column{RatesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// RatePricesColumns holds the columns for the "rate_prices" table.
	RatePricesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
		{Name: "price_id", Type: field.TypeUUID, Nullable: true},
		{Name: "rate_id", Type: field.TypeUUID, Nullable: true},
	}
	// RatePricesTable holds the schema information for the "rate_prices" table.
	RatePricesTable = &schema.Table{
		Name:       "rate_prices",
		Columns:    RatePricesColumns,
		PrimaryKey: []*schema.Column{RatePricesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "rate_prices_prices_price_id",
				Columns: []*schema.Column{RatePricesColumns[3]},

				RefColumns: []*schema.Column{PricesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "rate_prices_rates_rate_id",
				Columns: []*schema.Column{RatePricesColumns[4]},

				RefColumns: []*schema.Column{RatesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "balance", Type: field.TypeInt16},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
		{Name: "rate_id", Type: field.TypeUUID, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "users_rate_prices_user",
				Columns: []*schema.Column{UsersColumns[4]},

				RefColumns: []*schema.Column{RatePricesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserMonthMessagesColumns holds the columns for the "user_month_messages" table.
	UserMonthMessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "month", Type: field.TypeTime},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
		{Name: "provider_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
	}
	// UserMonthMessagesTable holds the schema information for the "user_month_messages" table.
	UserMonthMessagesTable = &schema.Table{
		Name:       "user_month_messages",
		Columns:    UserMonthMessagesColumns,
		PrimaryKey: []*schema.Column{UserMonthMessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "user_month_messages_provides_provider_id",
				Columns: []*schema.Column{UserMonthMessagesColumns[4]},

				RefColumns: []*schema.Column{ProvidesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "user_month_messages_users_user_messages",
				Columns: []*schema.Column{UserMonthMessagesColumns[5]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MessagesTable,
		PricesTable,
		ProvidesTable,
		RatesTable,
		RatePricesTable,
		UsersTable,
		UserMonthMessagesTable,
	}
)

func init() {
	MessagesTable.ForeignKeys[0].RefTable = ProvidesTable
	MessagesTable.ForeignKeys[1].RefTable = UsersTable
	RatePricesTable.ForeignKeys[0].RefTable = PricesTable
	RatePricesTable.ForeignKeys[1].RefTable = RatesTable
	UsersTable.ForeignKeys[0].RefTable = RatePricesTable
	UserMonthMessagesTable.ForeignKeys[0].RefTable = ProvidesTable
	UserMonthMessagesTable.ForeignKeys[1].RefTable = UsersTable
}
