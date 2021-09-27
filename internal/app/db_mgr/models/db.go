package models

import (
	"github.com/google/uuid"
)

type Station struct {
	ID       uuid.UUID `pg:"type:uuid,default:uuid_generate_v4()"`
	Capacity float32
	UsedCapacity float32
}

type Dock struct {
	ID       uuid.UUID `pg:"type:uuid,default:uuid_generate_v4()"`
	NumDockingPorts int
	Occupied int
	Weight float32
	StationID uuid.UUID
}

type Ship struct {
	ID       uuid.UUID `pg:"type:uuid,default:uuid_generate_v4()"`
	Status string
	Weight float32
}

type User struct {
	ID uuid.UUID `pg:"type:uuid,default:uuid_generate_v4()"`
	Username string
	Password string
	Role string
}