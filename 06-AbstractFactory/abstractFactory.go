package main

import "fmt"

/*
Patron de diseño Abstract Factory

Este patron nos permite crear familias de objetos que son bastente similares
*/

/*

Notificaciones -> SMS, Push Notification, Email

La idea es manejar

El problema es que hay un programa que envia Notificaciones
pero no sabe de que tipo es puede ser un SMS, Push Notification
o Email y la idea es crear un software que los maneje todos
*/

// Notificaciones por SMS y Email

/*  Factory */
type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

/* END Factory */

/*  SMS */
type SMSNotification struct {
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification by SMS")
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "twillio"
}

/*  END SMS */

/*  Email */
type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending Notification by Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

/* END Email */

/*  LLamadas concreta - Factory */
func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("No Notificacion Type")
}

/*  END - LLamadas concreta - Factory */

// En esete instante no sabemos que notificacion se disparara
func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderChannel())
}

func main() {
	fmt.Println("============== Implementar  ==================")
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)

	_, err := getNotificationFactory("DUMMY")
	fmt.Println(err)
}
