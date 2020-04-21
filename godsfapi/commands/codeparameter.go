package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var ErrMissingParameter = errors.New("Parameter not found")

// CodeParameter represents a parsed parameter of a G/M/T-code
type CodeParameter struct {
	// Letter of the code parameter (e.g. P in M106 P2)
	Letter string
	// IsExpression indicates if this parameter is an expression that can be evaluated
	IsExpression bool
	// IsDriverId indicated if this parameter is a driver identifier
	IsDriverId bool
	// stringValue is the unparsed string representation of the code parameter
	stringValue string
	// IsString indicates if this parameter is a string
	IsString bool
	// parsedValue is the internal parsed representation of the string value
	parsedValue interface{}
}

// NewCodeParameter creates a new CodeParameter instance and parses value to a native data type if applicable
func NewCodeParameter(letter, value string, isString bool) (*CodeParameter, error) {
	cp := &CodeParameter{}
	err := cp.init(letter, value, isString)
	return cp, err
}

// NewSimpleCodeParameter instantiates a CodeParameter for the given letter and value
func NewSimpleCodeParameter(letter string, value interface{}) *CodeParameter {
	_, isString := value.(string)
	return &CodeParameter{
		Letter:      letter,
		parsedValue: value,
		stringValue: fmt.Sprintf("%v", value),
		IsString:    isString,
	}
}

// String prints out the parameter with quotes around the value if it is a string parameter
func (cp CodeParameter) String() string {
	if cp.IsString {
		return fmt.Sprintf(`%s"%s"`, cp.Letter, strings.ReplaceAll(cp.stringValue, `"`, `""`))
	}
	return fmt.Sprintf("%s%s", cp.Letter, cp.stringValue)
}

// Clone will create a copy of the this instance
func (cp *CodeParameter) Clone() *CodeParameter {
	cpc := *cp
	return &cpc
}

// ConvertDriverIds converts this parameter to a driver id or a list of such
func (cp *CodeParameter) ConvertDriverIds() error {
	if cp.IsExpression {
		return nil
	}

	d := make([]uint64, 0)

	// Split on the list-separator
	params := strings.Split(cp.stringValue, ":")
	for _, p := range params {
		// Split on the board-driver-separator
		s := strings.Split(p, ".")

		// It was just one value
		if len(s) == 1 {
			u, err := strconv.ParseUint(s[0], 10, 64)
			if err != nil {
				return errors.New(fmt.Sprintf("Failed to parse driver number from %s parameter", cp.Letter))
			}
			d = append(d, u)

			// board id was also given
		} else if len(s) == 2 {
			board, err := strconv.ParseUint(s[0], 10, 64)
			if err != nil {
				return errors.New(fmt.Sprintf("Failed to parse board number from %s parameter", cp.Letter))
			}
			port, err := strconv.ParseUint(s[1], 10, 64)
			if err != nil {
				return errors.New(fmt.Sprintf("Failed to parse driver number from %s parameter", cp.Letter))
			}
			driver := (board << 16) | port
			d = append(d, driver)
		} else {
			return errors.New(fmt.Sprintf("Driver value from %s parameter is invalid.", cp.Letter))
		}
	}

	if len(d) == 1 {
		cp.parsedValue = d[0]
	} else {
		cp.parsedValue = d
	}
	cp.IsDriverId = true

	return nil
}

// AsFloat64 returns the value as float64 if it was of this type or can be converted to one or an error otherwise
func (cp *CodeParameter) AsFloat64() (float64, error) {
	if cp == nil {
		return 0, ErrMissingParameter
	}
	switch v := cp.parsedValue.(type) {
	case float64:
		return v, nil
	case uint64:
		return float64(v), nil
	case int64:
		return float64(v), nil
	default:
		return 0, errors.New(fmt.Sprintf("Cannot convert %s parameter to float64 (value %s of type %T)", cp.Letter, cp.stringValue, cp.parsedValue))
	}
}

// AsInt64 returns the value as int64 if it was of this type or can be converted to one or an error otherwise
func (cp *CodeParameter) AsInt64() (int64, error) {
	if cp == nil {
		return 0, ErrMissingParameter
	}
	switch v := cp.parsedValue.(type) {
	case int64:
		return v, nil
	case uint64:
		if v <= math.MaxInt32 {
			return int64(v), nil
		}
	}
	return 0, errors.New(fmt.Sprintf("Cannot convert %s parameter to int64 (value %s of type %T)", cp.Letter, cp.stringValue, cp.parsedValue))
}

// AsUint64 returns the value as uint64 if it was of this type or can be converted to one or an error otherwise
func (cp *CodeParameter) AsUint64() (uint64, error) {
	if cp == nil {
		return 0, ErrMissingParameter
	}
	switch v := cp.parsedValue.(type) {
	case uint64:
		return v, nil
	case int64:
		if v >= 0 {
			return uint64(v), nil
		}
	}
	return 0, errors.New(fmt.Sprintf("Cannot convert %s parameter to uint64 (value %s of type %T)", cp.Letter, cp.stringValue, cp.parsedValue))
}

// AsBool returns the value as bool as returned by strconv.ParseBool()
func (cp *CodeParameter) AsBool() (bool, error) {
	if cp == nil {
		return false, ErrMissingParameter
	}
	return strconv.ParseBool(cp.stringValue)
}

// AsString returns the string representation of this parameter
func (cp *CodeParameter) AsString() string {
	return cp.stringValue
}

// AsFloat64Slice converts this parameter to []float64 if it is a numeric type (or slice)
func (cp *CodeParameter) AsFloat64Slice() ([]float64, error) {
	if cp == nil {
		return nil, ErrMissingParameter
	}
	switch v := cp.parsedValue.(type) {
	case []float64:
		return v, nil
	case float64:
		return []float64{v}, nil
	case int64:
		return []float64{float64(v)}, nil
	case uint64:
		return []float64{float64(v)}, nil
	case []int64:
		fs := make([]float64, 0, len(v))
		for _, i := range v {
			fs = append(fs, float64(i))
		}
		return fs, nil
	case []uint64:
		fs := make([]float64, 0, len(v))
		for _, i := range v {
			fs = append(fs, float64(i))
		}
		return fs, nil
	}

	return nil, errors.New(fmt.Sprintf("Cannot convert %s parameter to []float64 (value %s of type %T)", cp.Letter, cp.stringValue, cp.parsedValue))
}

// AsInt64Slice converts this parameter to []int64 if it is a numeric type (or slice)
func (cp *CodeParameter) AsInt64Slice() ([]int64, error) {
	if cp == nil {
		return nil, ErrMissingParameter
	}
	switch v := cp.parsedValue.(type) {
	case []int64:
		return v, nil
	case float64:
		return []int64{int64(v)}, nil
	case int64:
		return []int64{int64(v)}, nil
	case uint64:
		return []int64{int64(v)}, nil
	case []float64:
		fs := make([]int64, 0, len(v))
		for _, i := range v {
			fs = append(fs, int64(i))
		}
		return fs, nil
	case []uint64:
		fs := make([]int64, 0, len(v))
		for _, i := range v {
			fs = append(fs, int64(i))
		}
		return fs, nil
	}

	return nil, errors.New(fmt.Sprintf("Cannot convert %s parameter to []int64 (value %s of type %T)", cp.Letter, cp.stringValue, cp.parsedValue))
}

// AsUint64Slice converts this parameter to []uint64 if it is a numeric type (or slice)
func (cp *CodeParameter) AsUint64Slice() ([]uint64, error) {
	if cp == nil {
		return nil, ErrMissingParameter
	}
	switch v := cp.parsedValue.(type) {
	case []uint64:
		return v, nil
	case float64:
		return []uint64{uint64(v)}, nil
	case int64:
		if v < 0 {
			goto Error
		}
		return []uint64{uint64(v)}, nil
	case uint64:
		return []uint64{uint64(v)}, nil
	case []float64:
		fs := make([]uint64, 0, len(v))
		for _, i := range v {
			fs = append(fs, uint64(i))
		}
		return fs, nil
	case []int64:
		fs := make([]uint64, 0, len(v))
		for _, i := range v {
			if i < 0 {
				goto Error
			}
			fs = append(fs, uint64(i))
		}
		return fs, nil
	}
Error:
	return nil, errors.New(fmt.Sprintf("Cannot convert %s parameter to []uint64 (value %s of type %T)", cp.Letter, cp.stringValue, cp.parsedValue))
}

// init will parse the string value of this parameter
func (cp *CodeParameter) init(letter, value string, isString bool) error {
	cp.Letter = letter
	cp.stringValue = value
	cp.IsString = isString
	if cp.IsString {
		cp.parsedValue = cp.stringValue
		return nil
	}
	value = strings.TrimSpace(value)
	if value == "" {
		cp.parsedValue = 0
	} else if strings.HasPrefix(value, "{") && strings.HasSuffix(value, "}") {
		cp.IsExpression = true
	} else if strings.Contains(value, ":") {
		cp.parseListValue(value)
	} else if i, err := strconv.ParseInt(value, 10, 64); err == nil {
		cp.parsedValue = i
	} else if u, err := strconv.ParseUint(value, 10, 64); err == nil {
		cp.parsedValue = u
	} else if f, err := strconv.ParseFloat(value, 64); err == nil {
		cp.parsedValue = f
	} else {
		cp.parsedValue = value
	}
	return nil
}

func (cp *CodeParameter) parseListValue(value string) {
	split := strings.Split(value, ":")
	success := true
	// FIXME: This is ugly
	if strings.Contains(value, ".") {

		// Try parsing as float64
		floats := make([]float64, 0)
		for _, s := range split {
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				success = false
				break
			}
			floats = append(floats, f)
		}
		if success {
			cp.parsedValue = floats
		}
	} else {

		// Try parsing as int64
		ints := make([]int64, 0)
		for _, s := range split {
			i, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				success = false
				break
			}
			ints = append(ints, i)
		}
		if success {
			cp.parsedValue = ints
		} else {

			// Try parsing as uint64
			success = true
			uints := make([]uint64, 0)
			for _, s := range split {
				u, err := strconv.ParseUint(s, 10, 64)
				if err != nil {
					success = false
					break
				}
				uints = append(uints, u)
			}
			if success {
				cp.parsedValue = uints
			}
		}
	}

	// In case it's not a numeric type keep it as a string
	if !success {
		cp.parsedValue = value
	}
}

func (cp *CodeParameter) MarshalJSON() ([]byte, error) {
	ss := make(map[string]interface{})
	ss["letter"] = cp.Letter
	ss["value"] = cp.stringValue
	if cp.IsString {
		ss["isString"] = 1
	} else {
		ss["isString"] = 0
	}

	return json.Marshal(ss)
}

func (cp *CodeParameter) UnmarshalJSON(data []byte) error {
	ss := make(map[string]interface{})
	err := json.Unmarshal(data, &ss)
	if err != nil {
		return err
	}
	var letter, value string
	var isString bool
	for k, v := range ss {
		if k == "letter" {
			letter = v.(string)
		} else if k == "value" {
			value = v.(string)
		} else if k == "isString" {
			isString = v.(float64) == 1
		}
	}
	return cp.init(letter, value, isString)
}
