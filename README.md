# {BS_REPO_NAME}

{BS_README_SHORT}

## Запуск рутинных операций
Для запуска рутинных операций, такие как статические анализаторы, тесты, всякие генерации и т.п. существует скрипт 
`run.sh`, посмотрите его `help`, чтобы ознакомиться с его возможностями.

## Сборка Docker Image
```bash
docker build -t {BS_REPO_NAME} . 
```

## Аргументы и переменные окружения
Смотрите список всех аргументов и переменных окружения в файле [cmd/{BS_REPO_NAME}/config.go](cmd/{BS_REPO_NAME}/config.go)

## Примеры запуска
### Поднятие отдельных сервисов-зависимостей
```bash
docker compose -f ci/dev/docker-compose.yml up
```

### Запуск docker-контейнера
```bash
docker run \
  --rm \
  -it \
  -e CL_LOG_LEVEL=info \
  -e CL_HTTP_PRIVATE_LISTEN=localhost:8080 \
  --env-file=ci/dev/.env \
  ./cmd/{BS_REPO_NAME}
```

### Запуск скомпилированного исполняемого файла
```bash
(set -a && source ci/dev/.env && set +a && ./cmd/{BS_REPO_NAME}/{BS_REPO_NAME})
```

## Ссылки и адреса
### Логи
- [prod](https://kibana7.citilink.cloud/app/discover#/?_g=(filters:!(),refreshInterval:(pause:!t,value:0),time:(from:now-15m,to:now))&_a=(columns:!(level,msg,error),filters:!(('$state':(store:appState),meta:(alias:!n,disabled:!f,index:'62c054a0-af23-11eb-b6fd-c9722691841f',key:kubernetes.container.name,negate:!f,params:(query:{BS_REPO_NAME}),type:phrase),query:(match_phrase:(kubernetes.container.name:{BS_REPO_NAME})))),index:'62c054a0-af23-11eb-b6fd-c9722691841f',interval:auto,query:(language:kuery,query:''),sort:!()))
- [stage](https://kibana7.citilink.cloud/app/discover#/?_g=(filters:!(),refreshInterval:(pause:!t,value:0),time:(from:now-15m,to:now))&_a=(columns:!(level,msg,error),filters:!(('$state':(store:appState),meta:(alias:!n,disabled:!f,index:'62c054a0-af23-11eb-b6fd-c9722691841f',key:kubernetes.container.name,negate:!f,params:(query:{BS_REPO_NAME}),type:phrase),query:(match_phrase:(kubernetes.container.name:{BS_REPO_NAME})))),index:'43fe9a40-ae83-11eb-b6fd-c9722691841f',interval:auto,query:(language:kuery,query:''),sort:!()))

### Мониторинг
- [grafana](https://grafana.citilink.ru/d/1hJ9vqBWk/{BS_REPO_NAME})

### Список переменных helm-чарта
[helm-general/values.yaml](https://git.citilink.cloud/shared/helm-general/src/branch/master/helm-general/values.yaml)

### Адреса микросервиса для доступа вне kubernetes (например, со своего ПК)
- stage: `{BS_REPO_NAME}.svc.citilink.lt:9999`
- prod: `{BS_REPO_NAME}.svc.citilink.ee:9999`

### Адреса микросервиса для доступа внутри kubernetes (например, при подключении в других микросервисах)
- stage: `{BS_REPO_NAME}.services.svc:9999`
- prod: `{BS_REPO_NAME}.services.svc:9999`

## СХЕМА ДЕПЛОЯ
https://confluence.citilink.ru/pages/viewpage.action?pageId=62733120

## CONTRIBUTE
 * Напишите код
 * Запустите тесты: `./run.sh test`
 * Создайте pull-request и назначьте ревьювером мейнтейнера

## MAINTAINER
Команда {BS_MAINTAINER_NAME}