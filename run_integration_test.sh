echo -----Run containers in the background-----
docker-compose up -d
sleep 1

cd migration
echo -----Run db migration-----
goose mysql "root:test@(127.0.0.1:3309)/kbc?parseTime=true" up

docker-compose up
