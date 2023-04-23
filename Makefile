api:
	cd server && make api

buildAdministratorServer:
	mkdir -p ./dist/server
	cd server/app/service/administrator && go mod tidy && make build
	mv server/app/service/administrator/bin/cmd ./dist/server/administrator.server

buildArticleServer:
	mkdir -p ./dist/server
	cd server/app/service/article && go mod tidy && make build
	mv server/app/service/article/bin/cmd ./dist/server/article.server

buildAdministratorInterface:
	mkdir -p ./dist/server
	cd server/app/interface/administrator && go mod tidy && make build
	mv server/app/interface/administrator/bin/cmd ./dist/server/administrator.interface.server

buildArticleInterface:
	mkdir -p ./dist/server
	cd server/app/interface/article && go mod tidy && make build
	mv server/app/interface/article/bin/cmd ./dist/server/article.interface.server

buildPassportInterface:
	mkdir -p ./dist/server
	cd server/app/interface/passport && go mod tidy && make build
	mv server/app/interface/passport/bin/cmd ./dist/server/passport.interface.server

buildAdminEntry:
	mkdir -p ./dist/admin && rm -rf ./dist/admin/entry
	cd admin/entry && npm install && npm run build
	mv admin/entry/dist ./dist/admin/entry

buildAdmiAdministrator:
	mkdir -p ./dist/admin && rm -rf ./dist/admin/administrator
	cd admin/administrator && npm install && npm run build && npm run manifest
	mv admin/administrator/dist ./dist/admin/administrator

buildAdminArticle:
	mkdir -p ./dist/admin && rm -rf ./dist/admin/article
	cd admin/article && npm install && npm run build && npm run manifest
	mv admin/article/dist ./dist/admin/article

buildClient:
	mkdir -p ./dist/client && rm -rf ./dist/client/**/*
	cd client && npm install && npm run build
	mv client/.output/* ./dist/client
