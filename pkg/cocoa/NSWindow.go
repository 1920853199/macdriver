// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cocoa

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -lobjc -framework Foundation
#define OBJC_OLD_DISPATCH_PROTOTYPES 1
#include <objc/runtime.h>
#include <objc/message.h>


*/
import "C"

import (
	"github.com/progrium/macdriver/pkg/core"
	"github.com/progrium/macdriver/pkg/objc"
)

const (
	NSBorderlessWindowMask         = 0
	NSTitledWindowMask             = 1 << 0
	NSClosableWindowMask           = 1 << 1
	NSMiniaturizableWindowMask     = 1 << 2
	NSResizableWindowMask          = 1 << 3
	NSTexturedBackgroundWindowMask = 1 << 8

	NSWindowTitleVisible = 0
	NSWindowTitleHidden  = 1

	NSWindowAbove = 1
	NSWindowBelow = -1
	NSWindowOut   = 0

	NSBackingStoreRetained    = 0
	NSBackingStoreNonretained = 1
	NSBackingStoreBuffered    = 2
)

type NSBackingStoreType uintptr

type NSWindow struct {
	objc.Object
}

var NSWindow_ = objc.Get("NSWindow")

func NSWindow_New() NSWindow {
	return NSWindow{NSWindow_.Alloc().Init()}
}

func NSWindow_WithContentViewController(controller objc.Object) NSWindow {
	return NSWindow{NSWindow_.Send("windowWithContentViewController:", controller)}
}

func NSWindow_Init(rect core.NSRect, windowStyle core.NSUInteger, bufferingType NSBackingStoreType, deferCreation bool) NSWindow {
	obj := NSWindow_.Alloc().
		Send("initWithContentRect:styleMask:backing:defer:",
			rect, windowStyle, bufferingType, deferCreation)
	return NSWindow{obj}
}

func (w NSWindow) Display() {
	w.Send("display")
}

func (w NSWindow) Center() {
	w.Send("center")
}

func (w NSWindow) MakeKeyAndOrderFront(sender objc.Object) {
	w.Send("makeKeyAndOrderFront:", sender)
}

func (w NSWindow) SetTitle(title string) {
	w.Send("setTitle:", core.String(title))
}

func (w NSWindow) Title() string {
	return w.Send("title").String()
}

func (w NSWindow) SetContentView(view objc.Object) {
	w.Send("setContentView:", view)
}

func (w NSWindow) ContentView() objc.Object {
	return w.Send("contentView")
}

func (w NSWindow) IsVisible() bool {
	return w.Send("isVisible").Bool()
}

func (w NSWindow) Frame() (frame core.NSRect) {
	w.Send("frame", &frame)
	return frame
}

func (w NSWindow) ContentRectForFrameRect(frameRect core.NSRect) (rect core.NSRect) {
	w.Send("contentRectForFrameRect:", frameRect, &rect)
	return rect
}
