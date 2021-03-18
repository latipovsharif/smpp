// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"smpp/ent/price"
	"smpp/ent/rate"
	"smpp/ent/rateprice"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// RatePrice is the model entity for the RatePrice schema.
type RatePrice struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt time.Time `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt time.Time `json:"update_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RatePriceQuery when eager-loading is set.
	Edges    RatePriceEdges `json:"edges"`
	price_id *uuid.UUID
	rate_id  *uuid.UUID
}

// RatePriceEdges holds the relations/edges for other nodes in the graph.
type RatePriceEdges struct {
	// IDRate holds the value of the id_rate edge.
	IDRate *Rate
	// IDPrice holds the value of the id_price edge.
	IDPrice *Price
	// User holds the value of the user edge.
	User []*User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// IDRateOrErr returns the IDRate value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RatePriceEdges) IDRateOrErr() (*Rate, error) {
	if e.loadedTypes[0] {
		if e.IDRate == nil {
			// The edge id_rate was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: rate.Label}
		}
		return e.IDRate, nil
	}
	return nil, &NotLoadedError{edge: "id_rate"}
}

// IDPriceOrErr returns the IDPrice value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RatePriceEdges) IDPriceOrErr() (*Price, error) {
	if e.loadedTypes[1] {
		if e.IDPrice == nil {
			// The edge id_price was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: price.Label}
		}
		return e.IDPrice, nil
	}
	return nil, &NotLoadedError{edge: "id_price"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e RatePriceEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[2] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RatePrice) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case rateprice.FieldCreateAt, rateprice.FieldUpdateAt:
			values[i] = &sql.NullTime{}
		case rateprice.FieldID:
			values[i] = &uuid.UUID{}
		case rateprice.ForeignKeys[0]: // price_id
			values[i] = &uuid.UUID{}
		case rateprice.ForeignKeys[1]: // rate_id
			values[i] = &uuid.UUID{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type RatePrice", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RatePrice fields.
func (rp *RatePrice) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case rateprice.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				rp.ID = *value
			}
		case rateprice.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				rp.CreateAt = value.Time
			}
		case rateprice.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				rp.UpdateAt = value.Time
			}
		case rateprice.ForeignKeys[0]:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field price_id", values[i])
			} else if value != nil {
				rp.price_id = value
			}
		case rateprice.ForeignKeys[1]:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field rate_id", values[i])
			} else if value != nil {
				rp.rate_id = value
			}
		}
	}
	return nil
}

// QueryIDRate queries the "id_rate" edge of the RatePrice entity.
func (rp *RatePrice) QueryIDRate() *RateQuery {
	return (&RatePriceClient{config: rp.config}).QueryIDRate(rp)
}

// QueryIDPrice queries the "id_price" edge of the RatePrice entity.
func (rp *RatePrice) QueryIDPrice() *PriceQuery {
	return (&RatePriceClient{config: rp.config}).QueryIDPrice(rp)
}

// QueryUser queries the "user" edge of the RatePrice entity.
func (rp *RatePrice) QueryUser() *UserQuery {
	return (&RatePriceClient{config: rp.config}).QueryUser(rp)
}

// Update returns a builder for updating this RatePrice.
// Note that you need to call RatePrice.Unwrap() before calling this method if this RatePrice
// was returned from a transaction, and the transaction was committed or rolled back.
func (rp *RatePrice) Update() *RatePriceUpdateOne {
	return (&RatePriceClient{config: rp.config}).UpdateOne(rp)
}

// Unwrap unwraps the RatePrice entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rp *RatePrice) Unwrap() *RatePrice {
	tx, ok := rp.config.driver.(*txDriver)
	if !ok {
		panic("ent: RatePrice is not a transactional entity")
	}
	rp.config.driver = tx.drv
	return rp
}

// String implements the fmt.Stringer.
func (rp *RatePrice) String() string {
	var builder strings.Builder
	builder.WriteString("RatePrice(")
	builder.WriteString(fmt.Sprintf("id=%v", rp.ID))
	builder.WriteString(", create_at=")
	builder.WriteString(rp.CreateAt.Format(time.ANSIC))
	builder.WriteString(", update_at=")
	builder.WriteString(rp.UpdateAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// RatePrices is a parsable slice of RatePrice.
type RatePrices []*RatePrice

func (rp RatePrices) config(cfg config) {
	for _i := range rp {
		rp[_i].config = cfg
	}
}