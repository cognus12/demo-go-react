##Режимы запуска

full - сервер статика+api
api - сервер предоставляет только api (для разработки, когда фронт запускается в dev режиме)


###Собрать и запустить (full)

`deno task serv`

###Запустить сервер в режиме full

`deno task go:full`

###Запустить сервер в режиме api

`deno task go:api`

###Запустить фронт в dev режиме

`deno task client:dev`

###Собрать фронт

`deno task client:build`