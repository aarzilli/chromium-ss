Opening a new tab in chromium (or google chrome) using the chromium-browser (google-chrome) command sometimes takes up to 5 (FIVE) seconds. It's $CURRENT_YEAR and somehow this is considered acceptable.

Chromium-ss speeds up this process by not being stupid about it (or perhaps by *being* stupid).

Rename the binary from `chromium-ss` to `google-chrome-ss` to use with google chrome. Works only on linux.

Documentation about [singleton socket](https://chromium.googlesource.com/chromium/chromium/+/master/chrome/browser/process_singleton_linux.cc)