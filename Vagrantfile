# -*- mode: ruby -*-
# vi: set ft=ruby :

#box = "bento/ubuntu-16.04"
box = "ubuntu/xenial64"

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|

  config.vm.box = box

  config.vm.network "private_network", ip: "192.168.33.10"

  config.vm.provider "virtualbox" do |vb|
    vb.memory = "2048"
  end

  config.vm.provision "shell", privileged: false, inline: <<-SHELL
    sudo apt-get update
    sudo apt-get install -y --no-install-recommends aptitude gcc git libc6-dev tzdata make mysql-server
    sudo useradd -m isucon
    sudo usermod -G sudo -a -s /bin/bash isucon
    sudo -u isucon cd ~ && git clone -b asai/add-vagrant https://github.com/care-1040/isucon8-qualify.git
    sudo -u isucon ~/isucon8-qualify/setup.sh
  SHELL
end
