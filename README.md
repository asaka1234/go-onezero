## Installation

```bash
$ go get gitee.com/yaonuo_dev/aiguoclient
```

## Usage
```bash
import "gitee.com/yaonuo_dev/aiguoclient"
```
Create a new client, then use the exposed services to access different parts of the OPenPlatform API.

You can use your AppID/AppSecret/AppSecret to create a new client:
```bash
package main

import (
    "gitee.com/yaonuo_dev/aiguoclient"
)

func main() {
    client := aiguoclient.NewClient(APP_ID, APP_KEY, APP_SECRET,BASE_URL_DEV)
}
```