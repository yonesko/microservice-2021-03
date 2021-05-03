package swagger

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	marshal, _ := json.Marshal(User{
		Username:  "1sFT0I0m67",
		FirstName: "6M1B1dtUN",
		LastName:  "3oEHpLF79",
		Email:     "76Hdp",
		Phone:     "b4PYX72rE2",
	})
	fmt.Println(string(marshal))
}
