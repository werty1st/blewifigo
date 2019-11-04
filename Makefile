up:
	vagrant up
	ssh-keygen -R [localhost]:2222
	ssh-add .vagrant/machines/default/virtualbox/private_key


down:
	vagrant halt

destroy:
	vagrant destroy