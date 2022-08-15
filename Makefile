sqlstart : 
	sudo systemctl start postgresql

stmysql :
	sudo systemctl start mysqld.service
enmysql :
	sudo systemctl enable mysqld.service
push :
	git push
gadd : 
	git add .
gcom :
	git commit -m "default commit"
