all: manager page
	mkdir -p target/bin && mkdir -p target/db && mv manager/backend/moyu-manager target/bin/ && mv page/backend/moyu-page target/bin/ && cat manager/backend/database/init.sql | sqlite3 target/db/moyu_manager.db

manager: manager-frontend
	cd manager/backend && go build -o moyu-manager --ldflags="-w -s" .

manager-frontend:
	cd manager/frontend && npm install && npm run build

page: page-frontend
	cd page/backend && go build -o moyu-page --ldflags="-w -s" .

page-frontend:
	cd page/frontend && npm install && npm run build

clean:
	-rm -rf page/backend/dist page/backend/moyu-page manager/backend/dist manager/backend/moyu-manager target

.PHONY: page-frontend page manager-frontend manager clean
