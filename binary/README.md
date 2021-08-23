<h3>You don't have go-lang binary and want the executable:</h3>

    curl -OL https://github.com/iHawke/kontext/raw/main/binary/kontext-OS-ARCH

    Where OS-ARCH in:  darwin-amd64, darwin-arm64, windows-386.exe, windows-amd64.exe, linux-386, linux-amd64, linux-arm64, openbsd-386, openbsd-amd64

    So, for most Mac amd the command will be:
    curl -OL https://github.com/iHawke/kontext/raw/main/binary/kontext-darwin-amd64

    So, for Mac Applie Silicon the command will be:
    curl -OL https://github.com/iHawke/kontext/raw/main/binary/kontext-darwin-arm64

    For most linux amd the command will be:
    curl -OL https://github.com/iHawke/kontext/raw/main/binary/kontext-linux-amd64

    For most windows amd the command will be:
    curl -OL https://github.com/iHawke/kontext/raw/main/binary/kontext-windows-amd64.exe

<h3>What is the error "identity of the developer cannot be confirmed"?</h3>

When installing Kontext on Macs, you may run across the error message:

    "Kontext" can't be opened because the identity of the developer cannot be confirmed.

This error may occur during the initial execution of *kontext*, due to recent security changes in Mac OS X 10.10 and 10.11. Fortunately, the error is benign, and it is easy to workaround.

**To allow Kontext to run on your Mac computer:**

1. Click OK if you encounter the error message above.
2. Click on Security & Privacy in System Preferences:
   ![Alt text](../image/mac-security-and-privacy.png?raw=true "")
3. Click to Allow app downloads from App Store and identified developers
   ![Alt text](../image/mac-select-developer-option.png?raw=true "")
4. Re-run *kontext*