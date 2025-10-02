package dto

import "github.com/google/uuid"

type CarRequest struct {
	CarTypeID            uuid.UUID `json:"car_type_id" binding:"required,uuid"`
	LicenseNumber        string    `json:"license_number" binding:"required"`
	MachineFrameNumber   string    `json:"machine_frame_number" binding:"required"`
	Color                string    `json:"color" binding:"required"`
	LicenseNumberExpired string    `json:"license_number_expired" binding:"required"`
}
