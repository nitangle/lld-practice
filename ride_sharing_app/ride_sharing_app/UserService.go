package ride_sharing_app

import (
	"encoding/json"
	"fmt"
)

type IUserService interface {
	AddUser(userDetails []byte) error
}

type User struct {
	Name    string
	Id      int
	Offered int
	Taken   int
}

type UserService struct {
	Users map[uint64]User
	vs    IVehicleService
}

func (us *UserService) AddVehicle(vehicleDetails *Vehicle) {
	vd, err := json.Marshal(&vehicleDetails)
	if err != nil {
		fmt.Println("err while marshalling vehicleDetails")
	}
	us.vs.AddVehicle(vd)
}

func (us *UserService) AddUser(userDetails []byte) error {

	var u User
	err := json.Unmarshal(userDetails, &u)
	if err != nil {
		fmt.Println("err while adding user details for Driver ", err)
		return err
	}
	id := GenerateId()
	us.Users[id] = u

	return nil
}
