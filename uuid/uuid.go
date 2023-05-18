package uuid

import (
	"database/sql/driver"
	"github.com/google/uuid"
)

// GormUUID uuid as binary(16) datatype
type GormUUID uuid.UUID

// StringToGormUUID
func StringToGormUUID(s string) (GormUUID, error) {
	id, err := uuid.Parse(s)
	return GormUUID(id), err
}

// String representation of Binary16
func (u *GormUUID) String() string {
	return uuid.UUID(*u).String()
}

// GormDataType sets type to binary(16)
func (u *GormUUID) GormDataType() string {
	return "binary(16)"
}

func (u *GormUUID) MarshalJSON() ([]byte, error) {
	s := uuid.UUID(*u)
	str := "\"" + s.String() + "\""
	return []byte(str), nil
}

func (u *GormUUID) UnmarshalJSON(by []byte) error {
	s, err := uuid.ParseBytes(by)
	*u = GormUUID(s)
	return err
}

// Scan how data are received from DB
func (u *GormUUID) Scan(value interface{}) error {
	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	*u = GormUUID(parseByte)
	return err
}

// Value how data are save into DB
func (u *GormUUID) Value() (driver.Value, error) {
	return uuid.UUID(*u).MarshalBinary()
}
