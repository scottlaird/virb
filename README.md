# Go library for Garmin Virb Cameras

This implements the [network
API](https://developer.garmin.com/downloads/virb/Camera_Network_Services_API_v0.5.pdf)
exposed by Garmin's Virb cameras with WiFi.  It has been tested with the Virb
Ultra 30 and Virb 360, but should work with the Virb X, XE, and several of
Garmin's Dashcam models as well.

It also includes a `virb` command-line tool that is mostly useful for testing
the library, although it can be used to start and stop recording on command.

This has only been tested with Go 1.12, but it probably compatible
with most recent Go versions.

### Using

This library generally follows Garmin's API as closely as possible.
Capitalization was changed to fit Go's requirements.  Each API call
has its own function in the `virb` module.  Each API function takes at
least one parameter, `hostname`, which is a string that contains the
hostname or IP address of the camera.  If the API call in
Garmin's doc takes any additional parameters, then they've been added in
order to the function.

Here's an example.

```go
package main

import (
	"fmt"
	"github.com/scottlaird/virb"
)

func main() {
     camera := "10.1.0.180"
     streamType := "rtp"
     maxResolution := 1080
     active := "1"
     
     resp, err := virb.LivePreview(camera, streamType, maxResolution, active)
     if err != nil { panic(err) }

     if resp.Result == 1 {
         fmt.Printf("The live stream URL is %s\n", resp.Url)
     }
}
```

See <https://scottstuff.net/posts/2019/05/07/virb/> for another
example or two.  The entire
[`virb/cmd`](https://github.com/scottlaird/virb/tree/master/virb/cmd)
directory in Github is basically just a giant example as well.

### Installing

With Go 1.12 (or possibly 1.11 with the new module system turned on),
simply importing `github.com/scottlaird/virb` is all that is needed;
go will automatically download and install this library and all of its
dependencies.

To install the command-line tool, make sure that you have Go 1.12
installed and run

```
   $ go get github.com/scottlaird/virb/virb
```

That will fetch all of the code and dependencies needed, compile them,
and install the resulting program into Go's binary directory; probably
`$HOME/go/bin` on most Unix systems.  It should work fine on Windows,
but hasn't been extensively tested.

### Disclaimer

This is not an officially supported Google product.
