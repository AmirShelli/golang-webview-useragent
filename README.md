# WebView SetUserAgent Implementation

This project implements the `SetUserAgent` functionality for the webview library on macOS, the architecture makes it easy to add Linux and Windows support later.
The `SetUserAgent()` method sets a custom user agent string in your Go webview applications.

## How to Use

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

Clone the repository:
```bash
git clone git@github.com:AmirShelli/golang-webview-useragent.git
cd golang-webview-useragent
```

Then build and run the demo:
```bash
chmod +x build.sh           # Make sure build script is executable
./build.sh                  # Builds the demo app
./webview-useragent-demo     # Runs it - opens webview with custom user agent
```

Or if you prefer to build manually:
```bash
go build -o webview-useragent-demo cmd/main.go
./webview-useragent-demo
```

**Requirements:**
- macOS (this implementation is macOS-only)
- Go 1.19+ installed
- No dependencies needed with WebKit built-in

## Implementation Overview

The implementation follows the same pattern as other webview functions to keep it consistent with the existing codebase.

### 1. **C API Declaration** (`webview/core/include/webview/api.h`)
```c
WEBVIEW_API webview_error_t webview_set_user_agent(webview_t w, const char *user_agent);
```

### 2. **C API Implementation** (`webview/core/include/webview/c_api_impl.hh`)
Wrapper that calls the C++ method:
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
Go wrapper over the C library:
```go
func (w *webview) SetUserAgent(userAgent string) {
    s := C.CString(userAgent)
    defer C.free(unsafe.Pointer(s))
    C.webview_set_user_agent(w.w, s)
}
```

## For Future Scaling

To add Linux support later, you'd just need to uncomment one line in `gtk_webkitgtk.hh`:
```cpp
// Currently: return {}; (placeholder)
// Change to: webkit_settings_set_user_agent(settings, user_agent.c_str());
```

To add Windows support later, same thing in `win32_edge.hh`:
```cpp
// Currently: return {}; (placeholder)
// Change to: settings2->put_UserAgent(wuser_agent.c_str());
```

The architecture is already there, just needs the platform-specific calls.

## Project Structure

```
.
├── webview/                    # Core C++ library
├── webview_go/                 # Go bindings
├── cmd/main.go                 # Demo app
├── build.sh                    # Build script
└── README.md                   # This file
```
