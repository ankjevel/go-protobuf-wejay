package types

import (
	"github.com/Iteam1337/go-protobuf-wejay/message"
	"github.com/golang/protobuf/proto"
)

// InputType …
type InputType byte

// InputTypeEnum …
const (
	IPing        InputType = 'p'
	IAction      InputType = 'a'
	INewUser     InputType = 'n'
	IUserExists  InputType = 'e'
	ICallbackURL InputType = 'c'
)

// Inv Response type
func (i InputType) Inv() (r ResponseType) {
	switch i {
	case IPing:
		r = RPong
	case IAction:
		r = RAction
	case INewUser:
		r = RNewUser
	case IUserExists:
		r = RUserExists
	case ICallbackURL:
		r = RCallbackURL
	}
	return
}

// Message protobuf message
func (i InputType) Message() (pb proto.Message) {
	switch i {
	case IPing:
		pb = &message.Ping{}
	case IAction:
		pb = &message.Action{}
	case INewUser:
		pb = &message.NewUser{}
	case IUserExists:
		pb = &message.UserExists{}
	case ICallbackURL:
		pb = &message.CallbackURL{}
	}
	return
}

// ResponseType …
type ResponseType byte

// ResponseTypeEnum …
const (
	RPong        ResponseType = 'P'
	RAction      ResponseType = 'A'
	RNewUser     ResponseType = 'N'
	RUserExists  ResponseType = 'E'
	RCallbackURL ResponseType = 'C'
)

// Inv Input type
func (r ResponseType) Inv() (i InputType) {
	switch r {
	case RPong:
		i = IPing
	case RAction:
		i = IAction
	case RNewUser:
		i = INewUser
	case RUserExists:
		i = IUserExists
	case RCallbackURL:
		i = ICallbackURL
	}
	return
}

// Message protobuf message
func (r ResponseType) Message() (pb proto.Message) {
	switch r {
	case RPong:
		pb = &message.Pong{}
	case RAction:
		pb = &message.ActionResponse{}
	case RNewUser:
		pb = &message.NewUserResponse{}
	case RUserExists:
		pb = &message.UserExistsResponse{}
	case RCallbackURL:
		pb = &message.CallbackURLResponse{}
	}
	return
}