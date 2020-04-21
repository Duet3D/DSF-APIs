package move

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// Kinematics is a placeholder type for SerDe
type Kinematics map[string]interface{}

func toKinematics(src interface{}) (Kinematics, error) {
	b, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}
	var k Kinematics
	err = json.Unmarshal(b, &k)
	if err != nil {
		return nil, err
	}
	return k, nil
}

// GetName returns the KinematicsName of this Kinematics instance
func (k Kinematics) GetName() KinematicsName {
	return KinematicsName(k["name"].(string))
}

// AsBaseKinematics returns this instance as BaseKinematics
func (k Kinematics) AsBaseKinematics() (*BaseKinematics, error) {
	bk := &BaseKinematics{}
	err := mapstructure.Decode(k, bk)
	if err != nil {
		return nil, err
	}
	return bk, nil
}

// AsCoreKinematics returns this instance as CoreKinematics
func (k Kinematics) AsCoreKinematics() (*CoreKinematics, error) {
	name := k.GetName()
	switch name {
	case Cartesian:
	case CoreXY:
	case CoreXYU:
	case CoreXYUV:
	case CoreXZ:
	case MarkForged:
	default:
		return nil, fmt.Errorf("Not core kinematics: %s", name)
	}
	ck := &CoreKinematics{
		ForwardMatrix: DefaultForwardMatrix(),
		InverseMatrix: DefaultInverseMatrix(),
	}
	err := mapstructure.Decode(k, ck)
	if err != nil {
		return nil, err
	}

	return ck, nil
}

// AsDeltaKinematics returns this instance as DeltaKinematics
func (k Kinematics) AsDeltaKinematics() (*DeltaKinematics, error) {
	name := k.GetName()
	switch name {
	case Delta:
	case RotaryDelta:
	default:
		return nil, fmt.Errorf("Not delta kinematics: %s", name)
	}

	dk := &DeltaKinematics{}
	err := mapstructure.Decode(k, dk)
	if err != nil {
		return nil, err
	}

	return dk, nil
}

// AsHangprinterKinematics returns this instance as HangprinterKinematics
func (k Kinematics) AsHangprinterKinematics() (*HangprinterKinematics, error) {
	name := k.GetName()
	switch name {
	case Hangprinter:
	default:
		return nil, fmt.Errorf("Not Hangprinter kinematics: %s", name)
	}

	hk := &HangprinterKinematics{
		AnchorA:     DefaultAnchorA(),
		AnchorB:     DefaultAnchorB(),
		AnchorC:     DefaultAnchorC(),
		AnchorDz:    DefaultAnchorDz,
		PrintRadius: DefaultHangprinterPrintRadius,
	}
	err := mapstructure.Decode(k, hk)
	if err != nil {
		return nil, err
	}

	return hk, nil
}

// BaseKinematics holds information about the configured kinematics
type BaseKinematics struct {
	// Name of currently configured kinematics
	Name KinematicsName `json:"name"`
}

// AsKinematics converts this instance to Kinematics type
func (bk *BaseKinematics) AsKinematics() Kinematics {
	return Kinematics{
		"name": bk.Name,
	}
}

// DefaultForwardMatrix for CoreKinematics
func DefaultForwardMatrix() [][]float64 {
	return [][]float64{
		[]float64{1, 0, 0},
		[]float64{0, 1, 0},
		[]float64{0, 0, 1},
	}
}

// DefaultInverseMatrix for CoreKinematics
func DefaultInverseMatrix() [][]float64 {
	return [][]float64{
		[]float64{1, 0, 0},
		[]float64{0, 1, 0},
		[]float64{0, 0, 1},
	}
}

// CoreKinematics holds information about core kinematics
type CoreKinematics struct {
	BaseKinematics `mapstructure:",squash"`
	// ForwardMatrix is the regular movement matrix
	ForwardMatrix [][]float64 `json:"forwardMatrix"`
	// InverseMatrix is the inverted movement matrix
	InverseMatrix [][]float64 `json:"inverseMatrix"`
}

// AsKinematics converts this instance to Kinematics type
func (ck *CoreKinematics) AsKinematics() (Kinematics, error) {
	return toKinematics(ck)

	// TODO: remove once the above has proven to work
	return Kinematics{
		"name":          ck.Name,
		"forwardMatrix": ck.ForwardMatrix,
		"inverseMatrix": ck.InverseMatrix,
	}, nil
}

// DeltaKinematics holds information about delta kinematics
type DeltaKinematics struct {
	BaseKinematics `mapstructure:",squash"`
	// DeltaRadius in mm
	DeltaRadius float64 `json:"deltaRadius"`
	// HomedHeight in mm
	HomedHeight float64 `json:"homedHeight"`
	// PrintRadius in mm
	PrintRadius float64 `json:"printRadius"`
	// Towers holds information about the delta Towers
	Towers []DeltaTower `json:"towers"`
	// XTilt is how much Z needs to be raised for each unit of movement
	// in the +X direction
	XTilt float64 `json:"xTilt"`
	// YTilt is how much Z needs to be raised for each unit of movement
	// in the +Y direction
	YTilt float64 `json:"yTilt"`
}

// AsKinematics converts this instance to Kinematics type
func (dk *DeltaKinematics) AsKinematics() (Kinematics, error) {
	return toKinematics(dk)

	// TODO: remove once the above has proven to work
	ts := make([]Kinematics, len(dk.Towers))
	for _, t := range dk.Towers {
		tk := Kinematics{
			"angleCorrection":   t.AngleCorrection,
			"Diagonal":          t.Diagonal,
			"endstopAdjustment": t.EndstopAdjustment,
			"xPos":              t.XPos,
			"yPos":              t.YPos,
		}
		ts = append(ts, tk)
	}
	return Kinematics{
		"name":        dk.Name,
		"deltaRadius": dk.DeltaRadius,
		"homedHeight": dk.HomedHeight,
		"printRadius": dk.PrintRadius,
		"towers":      ts,
		"xTilt":       dk.XTilt,
		"yTilt":       dk.YTilt,
	}, nil
}

// DeltaTower properties
type DeltaTower struct {
	// AngleCorrection represents tower position correction (in degrees)
	AngleCorrection float64 `json:"angleCorrection"`
	// Diagonal rod length (in mm)
	Diagonal float64 `json:"diagonal"`
	// EndstopAdjustment is the deviation of the ideal endstop position (in mm)
	EndstopAdjustment float64 `json:"endstopAdjustment"`
	// XPos is the X coordinate of this tower (in mm)
	XPos float64 `json:"xPos"`
	// YPos is the Y coordinate of this tower (in mm)
	YPos float64 `json:"yPos"`
}

const (
	// DefaultAnchorDz for HangprinterKinematics
	DefaultAnchorDz = 3000.0
	// DefaultHangprinterPrintRadius is the default radius for hangprinter
	DefaultHangprinterPrintRadius = 1500.0
)

// DefaultAnchorA for HangprinterKinematics
func DefaultAnchorA() []float64 { return []float64{0, -2000, -100} }

// DefaultAnchorB for HangprinterKinematics
func DefaultAnchorB() []float64 { return []float64{2000, 1000, -100} }

// DefaultAnchorC for HangprinterKinematics
func DefaultAnchorC() []float64 { return []float64{-2000, 1000, -100} }

// HangprinterKinematics properties
type HangprinterKinematics struct {
	BaseKinematics `mapstructure:",squash"`
	// AnchorA of the hangprinter
	AnchorA []float64 `json:"anchorA"`
	// AnchorB of the hangprinter
	AnchorB []float64 `json:"anchorB"`
	// AnchorC of the hangprinter
	AnchorC []float64 `json:"anchorC"`
	// AnchorDz of the hangprinter
	AnchorDz float64 `json:"anchorDz"`
	// PrintRadius in mm
	PrintRadius float64 `json:"printRadius"`
}

// AsKinematics converts this instance to Kinematics type
func (hk *HangprinterKinematics) AsKinematics() (Kinematics, error) {
	return toKinematics(hk)

	// TODO: remove once the above has proven to work
	return Kinematics{
		"name":        hk.Name,
		"anchorA":     hk.AnchorA,
		"anchorB":     hk.AnchorB,
		"anchorC":     hk.AnchorC,
		"anchorDz":    hk.AnchorDz,
		"printRadius": hk.PrintRadius,
	}, nil
}

// KinematicsName represents the supported kinmatics types
type KinematicsName string

const (
	// Cartesian kinematics
	Cartesian KinematicsName = "cartesian"
	// CoreXY kinematics
	CoreXY = "coreXY"
	// CoreXYU is a CoreXY kinematics with extra U axis
	CoreXYU = "coreXYU"
	// CoreXYUV is a CoreXY kinematics with extra UV axes
	CoreXYUV = "coreXYUV"
	// CoreXZ kinmatics
	CoreXZ = "coreXZ"
	// MarkForged kinematics
	MarkForged = "markForged"
	// FiveBarScara kinematics
	FiveBarScara = "FiveBarScara"
	// Hangprinter kinematics
	Hangprinter = "Hangprinter"
	// Delta kinematics
	Delta = "delta"
	// Polar kinematics
	Polar = "Polar"
	// RotaryDelta kinematics
	RotaryDelta = "Rotary delta"
	// Scara kinematics
	Scara = "Scara"
	// Unknown kinematics
	Unknown = "unknown"
)
