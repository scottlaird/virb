Go library for Garmin Virb Cameras

This implements the [network
API](https://developer.garmin.com/downloads/virb/Camera_Network_Services_API_v0.5.pdf)
exposed by Garmin's Virb cameras with WiFi.  It has been tested with the Virb
Ultra 30 and Virb 360, but should work with the Virb X, XE, and several of
Garmin's Dashcam models as well.

It also includes a `virb` command-line tool that is mostly useful for testing
the library, although it can be used to start and stop recording on command.

This has only been tested with Go 1.12, but it probably compatible
with most recent Go versions.

To install the command-line tool, make sure that you have Go 1.12 installed, and run

```
   $ go get github.com/scottlaird/virb/virb
```

That will fetch all of the code and dependencies needed, compile them,
and install the resulting program into Go's binary directory; probably
`$HOME/go/bin` on most Unix systems.


This is not an officially supported Google product.
