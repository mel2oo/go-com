package com

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

var _ COM = (*ShellWindows)(nil)

type ShellWindows struct {
	appid string
}

func NewShellWindows() *ShellWindows {
	return &ShellWindows{
		appid: "{9BA05972-F6A8-11CF-A442-00A0C90A8F39}",
	}
}

func (sw *ShellWindows) AppID() string {
	return sw.appid
}

func (sw *ShellWindows) ShellExecute(cmd ...interface{}) error {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	unknown, err := oleutil.CreateObject(sw.appid)
	if err != nil {
		return err
	}
	defer unknown.Release()

	shell, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return err
	}
	defer shell.Release()

	item := oleutil.MustCallMethod(shell, "Item").ToIDispatch()
	defer item.Release()

	doc := oleutil.MustGetProperty(item, "Document").ToIDispatch()
	defer doc.Release()

	app := oleutil.MustGetProperty(doc, "Application").ToIDispatch()
	defer app.Release()

	res := oleutil.MustCallMethod(app, "ShellExecute", cmd...)
	if res.Value() == nil {
		return nil
	} else {
		return ErrCallMethod
	}
}
