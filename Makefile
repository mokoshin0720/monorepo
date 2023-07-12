run:
	cd api; docker-compose up
	cd mobile; npx expo start --ios

ios:
	cd mobile; npx expo start --ios

server:
	cd api; docker-compose up