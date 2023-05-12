package main

import "fmt"

//SMS Email
type INotificacionFactory interface{
	SendNotification()
	GetSender() ISender
}
type ISender interface{
	GetSenderMethod() string
	GetSenderChannel() string
}



type SMSNotification struct{
	sender SMSNotificationSender
}
func (SMSNotification) SendNotification(){
	fmt.Println("Sending Notification via SMS")
}
func (s *SMSNotification) GetSender() ISender{
	return s.sender
}
type SMSNotificationSender struct{

}
func (SMSNotificationSender) GetSenderMethod()string{
	return "SMS"
}
func (SMSNotificationSender) GetSenderChannel()string{
	return "Twilio"
}



type EmailNotification struct{
	sender EmailNotificationSender
}
func (EmailNotification) SendNotification(){
	fmt.Println("Sending Notification via Email")
}
func (e *EmailNotification) GetSender()	ISender{
	return e.sender
}
type EmailNotificationSender struct{

}
func (EmailNotificationSender) GetSenderMethod()string{
	return "Email"
}
func (EmailNotificationSender) GetSenderChannel()string{
	return "SES"
}


func getNotificationFactory(notification string) (INotificacionFactory, error){
	if notification == "SMS"{
		return &SMSNotification{}, nil
	}
	if notification == "Email"{
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("didn't match with any notification type")
}
func sendNotification(f INotificacionFactory){
	f.SendNotification()
}
func getMethod(f INotificacionFactory){
	fmt.Println(f.GetSender().GetSenderMethod())
}
func Factory(){
	smsFactory, err1 := getNotificationFactory("SMS")
	emailFactory, err2 := getNotificationFactory("Email")
	if err1 != nil || err2 !=  nil{
		fmt.Println("Ha ocurrido un error :o")
	}else{
		fmt.Println("Con SMS")
		sendNotification(smsFactory)
		getMethod(smsFactory)
		fmt.Println("Con Email")
		sendNotification(emailFactory)
		getMethod(emailFactory)
	}
}