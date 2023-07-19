docker network create monorepo || true

docker-compose up -d mysql
docker-compose run --rm monorepo_dockerize -wait tcp://mysql:3306 -timeout 20s
