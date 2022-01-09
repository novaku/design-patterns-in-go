package framework

import "fmt"

type SecurityCode struct {
	Code int
}

func NewSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		Code: code,
	}
}

func (s *SecurityCode) CheckCode(incomingCode int) error {
	if s.Code != incomingCode {
		return fmt.Errorf("Security Code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}