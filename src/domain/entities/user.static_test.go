package entities

import "testing"

func TestUser_NewUserIr(t *testing.T) {

	test := User{}
	_, err := test.NewUser(RoleUser, "Test", "Test", "00000000000", "test@test.com", "123456", "123456", nil, nil)
	if err == nil {
		t.Errorf("Error: %s", err)
	}

}

func TestUser_NewUserPassword(t *testing.T) {

	test := User{}
	_, err := test.NewUser(RoleUser, "Test", "Test", "180.677.927-71", "", "123456", "123455", nil, nil)

	if err == nil {
		t.Errorf("Error: %s", err)
	}
}

func TestUser_NewUserNrle(t *testing.T) {

	test := User{}
	enterpriseName := "LV Digital"
	nrle := "50.081.719/0001-71"
	_, err := test.NewUser(RoleEnterprise, "Test", "Test", "180.677.927-71", "test@test.com", "123456", "123456", &enterpriseName, &nrle)

	if err == nil {
		t.Errorf("Error: %s", err)
	}

	nrle = "50081719000171"
	_, err = test.NewUser(RoleEnterprise, "Test", "Test", "180.677.927-71", "test@test.com", "123456", "123456", &enterpriseName, &nrle)

	if err == nil {
		t.Errorf("Error: %s", err)
	}

	nrle = "50.081.719/0001-72"

	_, err = test.NewUser(RoleEnterprise, "Test", "Test", "180.677.927-71", "test@test.com", "123456", "123456", &enterpriseName, &nrle)

	if err != nil {
		t.Errorf("Error: %s", err)
	}

}

func TestUser_AssignUser(t *testing.T) {

	test := User{}
	enterpriseName := "LV Digital"
	nrle := "50.081.719/0001-72"

	user, err := test.NewUser(RoleEnterprise, "Test", "Test", "180.677.927-71", "test@test.com", "123456", "123456", &enterpriseName, &nrle)

	if err != nil {
		t.Errorf("Error: %s", err)
	}
	userTest := User{}
	userTest.AssignUser(*user)
	if userTest.ID != user.ID {
		t.Errorf("Error: %s", err)
	}
}
