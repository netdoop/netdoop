
set -e

# install swag tool: go get -u github.com/swaggo/swag
# generate swag docs
swag init --pd --pdl 3 -g ./server/server.go

# update swagger.json
cd ./app
cp ../docs/swagger.json ./api.json
rm -rf services
pnpm openapi
cd ../
