package com

import "errors"

type COM interface {
	AppID() string
}

var (
	ErrCallMethod = errors.New("call method failed")
)
