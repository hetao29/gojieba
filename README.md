# 说明

基于[gojieba]开源的中文分词系统，特点：
1. 支持编译安装，纯golang代码
2. docker / docker swarm / k8s 部署
3. 支持多个自定义词库
4. rest api接口进行分词，也支持实时reload词库操作

## Docker & Swarm/Composer
```bash
docker pull hetao29/gojieba:latest
```

## 编译
```bash
make build
```

## 运行

```bash
make start
```

## 测试

```bash
make test
```
结果
```json
{"message":"pong","words":["关于","幼教","体系","组织","结构调整","结构","调整","等","的","通知"]}curl "http://127.0.0.1:8020/words?key=外国钱币硬币银铌世界纸钞爱藏"
{"message":"pong","words":["外国","钱币","硬币","银","铌","世界","纸钞","爱","藏"]}curl "http://127.0.0.1:8020/words?key=矮人火枪地狱兽残酷角斗士的军刺"
{"message":"pong","words":["矮人","火枪","地狱兽","地狱","残酷","角斗士","角斗","的","军","刺"]}
```
