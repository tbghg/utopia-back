
# Utopia-back

## 亮点介绍

+ [七牛云SDK使用](./doc/七牛云SDK使用.md)
  + 使用七牛云 kodo，采用 uploadToken + 回调 确保密钥安全性
  + 针对未上传封面的视频，异步截取视频首帧，回调替换视频封面
+ [热门视频设计方案](./doc/热门视频设计方案.md)
  + zset存储，异步刷取DB，获取一段时间内点赞量增长最高的视频，保持动态更新
  + version控制缓存版本，避免用户因热门视频突然更新而影响体验
+ 判断用户是否点赞缓存方案
  + 通过hash类型进行缓存，进行冷热数据分离，极大减少存储成本的同时也不会给DB带来过大负担
  + 根据用户维度存储减少查询成本，TTL合理续期
+ 自动部署方案
  + github webhook + hookdoo + 自定义脚本 实现持续部署，提高开发效率

## 使用说明

确保go版本大于等于1.21，通过`go version`命令检查

```shell
# 克隆项目
git clone git@github.com:VideoUtopia/utopia-back.git
cd utopia-back

# 修改配置文件，各字段含义见yaml文件字段备注
cp config/comfig_example.yaml config/congfig.yaml 
vim config/congfig.yaml

# linux 编译运行
go build -o server main.go 
./server

# windows
go build -o server.exe main.go
./server.exe
```

## 项目详细介绍

结构待安排

+ 数据库设计
+ 项目结构设计
+ 日志记录
+ 安全？IP限频、Token、salt
