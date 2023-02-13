# ob-example
Oceanbase基于gitpod建立了快速在线体验平台, 点击下面按钮一键体验(建议使用chrome浏览器):

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/oceanbase/ob-example)

进入后会自动部署一个OceanBase本地实例，请等待左边Oceanbase boot Success, 其余操作可在右边终端中操作
![示意图](./tools/scripts/gitpod1.png)

另外仓库中提供了不同语言和工具连接Oceanbase的示例，其中有run.sh的可直接在线体验，按下面步骤进行操作
```
//1. open 
cd xxxx
//2. prepare relative env
sh env.sh
//3. run to query
sh run.sh
```
这里以python3为例
```
cd python3-pymysql
sh env.sh
sh run.sh
```

关于更多Oceanbase的细节请参考 [OceanBase](https://open.oceanbase.com).

