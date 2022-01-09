package main

import (
	"fmt"

	"design-patterns/behavioral-patterns/template-method/framework"
)

func main() {

	smsOTP := &framework.Sms{}
	o := framework.Otp{
		IOtp: smsOTP,
	}
	err := o.GenAndSendOTP(smsOTP,4)
	if err != nil {
		fmt.Printf("Error send SMS OTP, error: %+v", err)
	}

	fmt.Println("")
	emailOTP := &framework.Email{}
	o = framework.Otp{
		IOtp: emailOTP,
	}
	err = o.GenAndSendOTP(emailOTP,4)
	if err != nil {
		fmt.Printf("Error send Email OTP, error: %+v", err)
	}
}