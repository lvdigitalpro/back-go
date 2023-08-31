package contracts

import (
	"fmt"
	"testing"
)

func TestUserContract(t *testing.T) {

	user := UserContract{}
	userTest, err := user.NewUser("USER", "Test", "Test", "180.677.927-71", "test@test.com", "123456", "123456", nil, nil)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	fmt.Println(user.Name)

	if userTest.Validate() == false {
		t.Errorf("Error: %s", "User is not valid")
	}
}
