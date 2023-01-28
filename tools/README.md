# 存放工程下的相关工具

## 初始化
通过kratos提供的初始化命令可以指定初始化使用的模板工程
```shell
kratos new projectName -r http:gitxxx.com/xxx -b branch
```
使用该模板工程初始化工程结构后，执行如下命令初始化项目
该脚本初始化过程利用kratos重新编译proto文件，确保`kratos`、`protoc`已安装
```shell
cd projectName/tools
chmod a+x init.sh
./init.sh -n projectName -c configs/xxx/xxx.yaml
```
该初始化脚本修改相关目录结构，通过指定配置文件直接运行该模板目  
***
`init.sh`脚本初始化命令如下，默认配置文件路径为`${workspaceFolder}/local/dev/config.yaml`
```shell
eg: ./init.sh [arg ...]
        Option:
         -n      project name
         -p      proto name
         -c      path of config file
```
