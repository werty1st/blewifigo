init:
	vagrant up
	ssh-keygen -R [localhost]:2222
	ssh-add .vagrant/machines/default/virtualbox/private_key


start:
	ssh-add .vagrant/machines/default/virtualbox/private_key
	vagrant reload

stop:
	vagrant halt

destroy:
	vagrant destroy