package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type DriverId struct {
	Board uint64
	Port  uint64
}

func NewDriverId(board, port uint64) DriverId {
	return DriverId{board, port}
}

func NewDriverIdUint64(value uint64) DriverId {
	return DriverId{
		Board: (value >> 16) & 0xFFFF,
		Port:  value & 0xFFFF,
	}
}

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
		return d, errors.New("Driver value is invalid.")
	}
	return d, nil
}

func (d *DriverId) AsUint64() uint64 {
	return (d.Board << 16) | d.Port
}

func (d *DriverId) String() string {
	return fmt.Sprintf("%d.%d", d.Board, d.Port)
}
