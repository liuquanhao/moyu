all: frontend
	cd backend && go build --ldflags="-w -s" .

frontend:
	cd frontend && npm run build

clean:
	rm -rf backend/dist backend/moyu

.PHONY: frontend all clean