all: manager page

manager: manager-frontend
	cd manager/backend && go build -o moyu-manager --ldflags="-w -s" .

manager-frontend:
	cd manager/frontend && npm run build

page: page-frontend
	cd page/backend && go build -o moyu-page --ldflags="-w -s" .

page-frontend:
	cd page/frontend && npm run build

clean:
	-rm -rf page/backend/dist page/backend/moyu-page manager/backend/dist manager/backend/moyu-manager

.PHONY: page-frontend page manager-frontend manager clean
