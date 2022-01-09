package framework

type IOtp interface {
	GenRandomOTP(int) string
	SaveOTPCache(string)
	GetMessage(string) string
	SendNotification(string) error
	PublishMetric()
}


func (o *Otp) GenAndSendOTP(iOtp IOtp, otpLength int) error {
 otp := iOtp.GenRandomOTP(otpLength)
 iOtp.SaveOTPCache(otp)
 message := iOtp.GetMessage(otp)
 err := iOtp.SendNotification(message)
 if err != nil {
     return err
 }
 iOtp.PublishMetric()
 return nil
}

type Otp struct {
	IOtp IOtp
}

func (o *Otp) genAndSendOTP(otpLength int) error {
	otp := o.IOtp.GenRandomOTP(otpLength)
	o.IOtp.SaveOTPCache(otp)
	message := o.IOtp.GetMessage(otp)
	err := o.IOtp.SendNotification(message)
	if err != nil {
		return err
	}
	o.IOtp.PublishMetric()
	return nil
}