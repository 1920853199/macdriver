package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"time"

	"github.com/progrium/macdriver/pkg/cocoa"
	"github.com/progrium/macdriver/pkg/core"
	"github.com/progrium/macdriver/pkg/objc"
	"github.com/progrium/macdriver/pkg/webkit"
	"github.com/progrium/watcher"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	var err error

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dir := filepath.Join(usr.HomeDir, ".topframe")
	os.MkdirAll(dir, 0755)

	srv := http.Server{
		Handler: http.FileServer(http.Dir(dir)),
	}

	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}

	go srv.Serve(ln)

	fw := watcher.New()
	if err := fw.AddRecursive(dir); err != nil {
		log.Fatal(err)
	}

	go fw.Start(400 * time.Millisecond)

	app := cocoa.NSApp_WithDidLaunch(func(notification objc.Object) {
		config := webkit.WKWebViewConfiguration_New()
		config.Preferences().SetValueForKey(core.True, core.String("developerExtrasEnabled"))

		wv := webkit.WKWebView_Init(cocoa.NSScreen_Main().Frame(), config)
		wv.SetOpaque(false)
		wv.SetBackgroundColor(cocoa.NSColor_Clear())
		wv.SetValueForKey(core.False, core.String("drawsBackground"))

		url := core.URL(fmt.Sprintf("http://%s", ln.Addr().String()))
		req := core.NSURLRequest_Init(url)
		wv.LoadRequest(req)

		w := cocoa.NSWindow_Init(cocoa.NSScreen_Main().Frame(),
			cocoa.NSBorderlessWindowMask, cocoa.NSBackingStoreBuffered, false)
		w.SetContentView(wv)
		w.SetBackgroundColor(cocoa.NSColor_Clear())
		w.SetOpaque(false)
		w.SetTitleVisibility(cocoa.NSWindowTitleHidden)
		w.SetTitlebarAppearsTransparent(true)
		w.SetIgnoresMouseEvents(true)
		w.SetLevel(cocoa.NSMainMenuWindowLevel + 2)
		w.MakeKeyAndOrderFront(w)

		go func() {
			for {
				select {
				case event := <-fw.Event:
					if event.IsDir() {
						continue
					}
					wv.Reload(nil)
				case <-fw.Closed:
					return
				}
			}
		}()
	})
	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyAccessory)
	app.ActivateIgnoringOtherApps(true)

	log.Printf("topframe 0.1.0 by progrium\n")
	app.Run()
}
