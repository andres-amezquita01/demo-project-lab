#!/bin/bash
sudo yum install -y docker
sudo usermod -a -G docker ec2-user
sudo yum install -y python3-pip
sudo pip3 uninstall -y urllib3

sudo pip3 install 'urllib3<2.0'

sudo pip3 install docker-compose

sudo systemctl enable docker.service
sudo systemctl start docker.service
sudo mkdir ~/prometheus