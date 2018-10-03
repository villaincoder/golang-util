package util

import (
	"fmt"
	"reflect"
	"time"
)

type Checker interface {
	Check() bool
	Error() string
}

type StringChecker struct {
	Value string
}

func (c *StringChecker) Check() bool {
	return c.Value != ""
}

func (c *StringChecker) Error() string {
	return "StringChecker value is empty"
}

type TimeChecker struct {
	Value time.Time
}

func (c *TimeChecker) Check() bool {
	return !IsInvalidTime(c.Value)
}

func (c *TimeChecker) Error() string {
	return "TimeChecker value is invalid"
}

type UnixMillisecondChecker struct {
	Value int64
}

func (c *UnixMillisecondChecker) Check() bool {
	return c.Value > 0
}

func (c *UnixMillisecondChecker) Error() string {
	return "UnixMillisecondChecker value is invalid"
}

type InterfaceChecker struct {
	Value interface{}
}

func (c *InterfaceChecker) Check() bool {
	return c.Value != nil && !reflect.ValueOf(c.Value).IsNil()
}

func (c *InterfaceChecker) Error() string {
	return "InterfaceChecker value is nil"
}

type StringArrayChecker struct {
	Value []string
	error string
}

func (c *StringArrayChecker) Check() bool {
	if c.Value == nil {
		c.error = "StringArrayChecker value is nil"
		return false
	}
	for index, value := range c.Value {
		if value == "" {
			c.error = fmt.Sprintf("StringArrayChecker item %d is empty", index)
			return false
		}
	}
	return true
}

func (c *StringArrayChecker) Error() string {
	return c.error
}

type IntegerChecker struct {
	Value int64
	Min   int64
	Max   int64
	error string
}

func (c *IntegerChecker) Check() bool {
	if c.Value < c.Min {
		c.error = fmt.Sprintf("IntegerChecker value %d < min %d", c.Value, c.Min)
		return false
	}
	if c.Value > c.Max {
		c.error = fmt.Sprintf("IntegerChecker value %d > max %d", c.Value, c.Max)
		return false
	}
	return true
}

func (c *IntegerChecker) Error() string {
	return c.error
}

type DecimalChecker struct {
	Value float64
	Min   float64
	Max   float64
	error string
}

func (c *DecimalChecker) Check() bool {
	if c.Value < c.Min {
		c.error = fmt.Sprintf("DecimalChecker value %f < min %f", c.Value, c.Min)
		return false
	}
	if c.Value > c.Max {
		c.error = fmt.Sprintf("DecimalChecker value %f > max %f", c.Value, c.Max)
		return false
	}
	return true
}

func (c *DecimalChecker) Error() string {
	return c.error
}

type UUIDChecker struct {
	Value string
}

func (c *UUIDChecker) Check() bool {
	return !IsInvalidUUID(c.Value)
}

func (c *UUIDChecker) Error() string {
	return "UUIDChecker value is invalid"
}

type UUIDArrayChecker struct {
	Value []string
	error string
}

func (c *UUIDArrayChecker) Check() bool {
	if c.Value == nil {
		c.error = "UUIDArrayChecker value is nil"
		return false
	}
	for index, value := range c.Value {
		if IsInvalidUUID(value) {
			c.error = fmt.Sprintf("UUIDArrayChecker item %d is invalid", index)
			return false
		}
	}
	return true
}

func (c *UUIDArrayChecker) Error() string {
	return c.error
}

func getCheckCallerName() (name string) {
	return GetCallerName(2)
}

func CheckArgs(fields map[string]Checker) (err error) {
	for key, checker := range fields {
		if checker == nil {
			err = fmt.Errorf("%s check %s error: checker is nil", getCheckCallerName(), key)
		} else if !checker.Check() {
			err = fmt.Errorf("%s check %s error: %s", getCheckCallerName(), key, checker.Error())
		}
	}
	return
}
