package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// DriverId represents a driver identification
type DriverId struct {
	// Board of this driver identifier
	Board uint64
	// Port of this driver identifier
	Port uint64
}

// NewDriverId creates a new DriverId from the given board and port values
func NewDriverId(board, port uint64) DriverId {
	return DriverId{board, port}
}

// NewDriverIdUint64 creates a new DriverId from the given bit-masked board and port value
func NewDriverIdUint64(value uint64) DriverId {
	return DriverId{
		Board: (value >> 16) & 0xFFFF,
		Port:  value & 0xFFFF,
	}
}

// NewDriverIdString creates a new DriverId from the given string
func NewDriverIdString(value string) (DriverId, error) {
	d := DriverId{}
	if strings.TrimSpace(value) == "" {
		return d, nil
	}

	s := strings.Split(value, ".")

	// It was just one value
	if len(s) == 1 {
		u, err := strconv.ParseUint(s[0], 10, 64)
		if err != nil {
			return d, errors.New("Failed to parse driver number")
		}
		d.Port = u

		// board id was also given
	} else if len(s) == 2 {
		board, err := strconv.ParseUint(s[0], 10, 64)
		if err != nil {
			return d, errors.New("Failed to parse board number")
		}
		port, err := strconv.ParseUint(s[1], 10, 64)
		if err != nil {
			return d, errors.New("Failed to parse driver number")
		}
		d.Board = board
		d.Port = port
	} else {
		return d, errors.New("Driver value is invalid")
	}
	return d, nil
}

// AsUint64 converts this instance to uint64
func (d *DriverId) AsUint64() uint64 {
	return (d.Board << 16) | d.Port
}

func (d *DriverId) String() string {
	return fmt.Sprintf("%d.%d", d.Board, d.Port)
}
