# -*- mode: ruby -*-
# vi: set ft=ruby :

#box = "bento/ubuntu-16.04"
box = "bento/centos-7.5"

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
  config.vm.box = box
  config.vm.network "private_network", ip: "192.168.33.10"
  config.vm.hostname = "isucon8-vagrant"
  config.vm.provider "virtualbox" do |vb|
    vb.memory = "2048"
    vb.name = "isucon8-vagrant"
  end

  #config.vm.synced_folder ".", "/home/isucon/isucon8-qualify", owner: "iscuon", group: "isucon"
end
