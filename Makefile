build:
	go build -o _build/main.exe cmd/release-server/main.go

server:
	air

ui:
	cd web && npm run dev
