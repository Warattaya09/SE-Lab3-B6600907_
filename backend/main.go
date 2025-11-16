package models

import (
	"encoding/json"

	"gorm.io/gorm"
)
type Student struct {
 gorm.Model
 FirstName string `json:"first_name"`
 LastName string `json:"last_name"`
 Age int `json:"age"`
 Email string `json:"email"`
 Email1 json.RawMessage `json:"email111111111111111111111111111111111111111111111"`
 //test 1234567890
 ////////////////////

}




