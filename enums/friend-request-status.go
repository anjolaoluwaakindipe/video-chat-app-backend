package enums

import (
	"database/sql/driver"
	"fmt"
	"log"
)

type FriendRequestStatus int

const (
	Accepted FriendRequestStatus = iota + 1
	Denied
	Ignored
	Block
)

func (fs *FriendRequestStatus) Scan(value interface{}) error {
	v, error := driver.Int32.ConvertValue(value)
	if error !=nil {
		log.Fatal("Error when converting from db value to golang data Type")
	}
	result, ok := v.(int)
	if ok {
		*fs = FriendRequestStatus(result)

	}
	return nil
}

func (fs FriendRequestStatus) Value() (driver.Value, error) {
	if (int(fs)> int(Block) || int(fs)< int(Accepted)){
		return nil, fmt.Errorf("Not an acceptable friend Request Status")
	}
	return int(fs), nil
}

func (fs *FriendRequestStatus) String() string{
	return [...]string{"Accepted", "Denied", "Ignored", "Block"}[*fs - 1]
}
