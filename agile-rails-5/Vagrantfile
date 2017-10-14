# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box      = 'ubuntu/yakkety64' # 16.10
  config.vm.hostname = 'agile-rails-5-box'

  config.vm.network :forwarded_port, guest: 3000, host: 3000, host_ip: "127.0.0.1"
  config.vm.network :forwarded_port, guest: 3306, host: 3306, host_ip: "127.0.0.1"

  config.vm.provision :shell, path: 'bootstrap_rails.sh', keep_color: true

  config.vm.provider 'virtualbox' do |v|
    v.memory = 2048
    v.cpus = 2
  end

  config.vm.synced_folder "./", "/var/www/rails-app/"
end
