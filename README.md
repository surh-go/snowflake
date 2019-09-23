# Snowflake

提供简易的`Snowflake`算法

<br/>

## 快速上手

### 获取下一个ID

``` go
package main

import (
  "github.com/sirupsen/logrus"
  "github.com/surh-go/snowflake"
  
  _"github.com/surh-go/logrus_conf"
)

func main() {
  logrus.Info(snowflake.NextID().String())
}
```

<br/>

### 解析ID

``` go
package main

import (
  "github.com/sirupsen/logrus"
  "github.com/surh-go/snowflake"
  
  _"github.com/surh-go/logrus_conf"
)

func main() {
  var id int64 = 24074082846773248
  info := ParseID(&id)
  logrus.Info("Timestamp: ", info.Timestamp)
  logrus.Info("MachineID: ", info.MachineID)
}
```
