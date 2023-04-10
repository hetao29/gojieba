build:
	export GOPROXY=https://goproxy.cn && go build -o bin/gojieba .
start:	
	nohup /home/hetao/gojieba/bin/gojieba -b="0.0.0.0:8030" &
stop:
	killall gojieba
test:
	curl "http://127.0.0.1:8030/words?key=关于幼教体系组织结构调整等的通知"
	curl "http://127.0.0.1:8030/words?key=外国钱币硬币银铌世界纸钞爱藏"
	curl "http://127.0.0.1:8030/words?key=矮人火枪地狱兽残酷角斗士的军刺"
reload:
	curl "http://127.0.0.1:8030/reload"
docker-image:
	DOCKER_BUILDKIT=1 docker build -t hetao29/gojieba .
docker-image-nocache:
	DOCKER_BUILDKIT=1 docker build --no-cache -t hetao29/gojieba .
docker-push:
	docker push hetao29/gojieba:latest
