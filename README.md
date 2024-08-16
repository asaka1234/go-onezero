## Installation

```bash
$ go get gitee.com/yaonuo_dev/aiguoclient_v2
```

## Usage
```bash
import "gitee.com/yaonuo_dev/aiguoclient_v2"
```
Create a new client, then use the exposed services to access different parts of the OPenPlatform API.

You can use your AppID/AppSecret/AppSecret to create a new client:
```bash
package main

import (
    "gitee.com/yaonuo_dev/aiguoclient_v2"
)

func main() {
    client := aiguoclient.NewClient(APP_ID, APP_KEY,BASE_URL_DEV)
}
```

## doc
https://www.showdoc.com.cn/2534208326234742/11263978442902776