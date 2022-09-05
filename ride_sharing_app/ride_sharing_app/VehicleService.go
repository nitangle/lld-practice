package ride_sharing_app

import (
	"encoding/json"
	"fmt"
)

type IVehicleService interface {
	AddVehicle(vehicleDetails []byte) error
}

type VehicleService struct {
	Vehicles map[uint64][]Vehicle
}

func (vs *VehicleService) AddVehicle(vehicleDetails []byte) error {
	var v Vehicle
	err := json.Unmarshal(vehicleDetails, &v)
	if err != nil {
		fmt.Println("err while adding user details for Vehicles ", err)
		return err
	}
	if len(vs.Vehicles[v.ownedBy]) == 0 {
		vs.Vehicles[v.ownedBy] = []Vehicle{v}
	} else {
		vs.Vehicles[v.ownedBy] = append(vs.Vehicles[v.ownedBy], v)
	}

	return nil
}
