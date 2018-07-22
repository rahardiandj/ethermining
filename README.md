# ethermining

This program is fetching transaction from Etherium.io and strore it to MongoDB. 

Basic functionality : 
* Show transactions
* Blacklist transactions
* Bookmak transactions

Cron will be running to fetch the data regulary and should be feature to delete from resository as well.

To run using docker just type comand `docker-compose up`

To run  without docker, run `main.go` in main folder and cron folder
