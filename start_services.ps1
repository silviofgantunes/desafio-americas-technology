#./start_services.sh

# 1) Stop the containers
docker stop database-americas-technology auth-service-americas-technology crud-users-americas-technology order-service-americas-technology

# 2) Remove the containers
docker rm database-americas-technology auth-service-americas-technology crud-users-americas-technology order-service-americas-technology

# 3) Remove the images
docker rmi desafio-americas-technology-database desafio-americas-technology-auth-service desafio-americas-technology-crud-users desafio-americas-technology-order-service

# 4) Clean unused networks
docker network prune

# 5) Start the database service
docker-compose up -d database
sleep 3

# 6) Start the authentication service
docker-compose up -d auth-service
sleep 3

# 7) Start the users service
docker-compose up -d crud-users
sleep 3

# 8) Start the orders service
docker-compose up -d order-service
