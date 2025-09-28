# WebView SetUserAgent Implementation

This project implements the `SetUserAgent` functionality for the webview library on **macOS only**, with an architecture that makes it easy to add Linux and Windows support later.

## What This Does

Adds a `SetUserAgent()` method to webview so you can set a custom user agent string in your Go webview applications.

## Simple Usage

```go
package main

import "github.com/webview/webview_go"

func main() {
    w := webview.New(false)
    defer w.Destroy()
    w.SetTitle("User Agent Demo")
    w.SetSize(480, 320, webview.HintNone)

    // This is the new functionality:
    w.SetUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

    w.Navigate("https://www.whatismybrowser.com/detect/what-is-my-user-agent/")
    w.Run()
}
```

## How to Build & Run

```bash
./build.sh                 # Builds the demo app
./webview-useragent-demo    # Runs it - you'll see the custom user agent!
```

## Simple Implementation Overview

This implementation touches exactly **4 files**:

### 1. **C API Declaration** (`webview/core/include/webview/api.h`)
Just added one line:
```c
WEBVIEW_API webview_error_t webview_set_user_agent(webview_t w, const char *user_agent);
```

### 2. **C API Implementation** (`webview/core/include/webview/c_api_impl.hh`)
Simple wrapper that calls the C++ method:
```cpp
WEBVIEW_API webview_error_t webview_set_user_agent(webview_t w, const char *user_agent) {
  return api_filter([=] { return cast_to_webview(w)->set_user_agent(user_agent); });
}
```

### 3. **macOS Implementation** (`webview/core/include/webview/detail/backends/cocoa_webkit.hh`)
Uses macOS's native WebKit API:
```cpp
noresult set_user_agent_impl(const std::string &user_agent) override {
  objc::autoreleasepool arp;
  WKWebView_setCustomUserAgent(m_webview, NSString_stringWithUTF8String(user_agent));
  return {};
}
```

### 4. **Go Binding** (`webview_go/webview.go`)
Simple Go wrapper:
```go
func (w *webview) SetUserAgent(userAgent string) {
    s := C.CString(userAgent)
    defer C.free(unsafe.Pointer(s))
    C.webview_set_user_agent(w.w, s)
}
```

## Easy Future Scaling

**To add Linux support later**, you'd just need to uncomment one line in `gtk_webkitgtk.hh`:
```cpp
// Currently: return {}; (placeholder)
// Change to: webkit_settings_set_user_agent(settings, user_agent.c_str());
```

**To add Windows support later**, same thing in `win32_edge.hh`:
```cpp
// Currently: return {}; (placeholder)
// Change to: settings2->put_UserAgent(wuser_agent.c_str());
```

That's it! The architecture is already there, just needs the platform-specific calls.

## Dependencies

- **macOS**: None (WebKit built-in)
- **Future Linux**: `webkit2gtk-4.0` dev packages
- **Future Windows**: WebView2 runtime

## Project Structure

```
.
├── webview/                    # Core C++ library
├── webview_go/                 # Go bindings
├── cmd/main.go                 # Demo app
├── build.sh                    # Build script
└── README.md                   # This file
```

**Key insight**: This follows the exact same pattern as other webview functions like `SetTitle()` and `SetSize()`, so it's very consistent with the existing codebase.