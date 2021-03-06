VAGRANTFILE_API_VERSION = "2"

Vagrant.require_version ">= 1.8.4"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.communicator = "winrm"

  config.vm.synced_folder ".", "/vagrant", disabled: true
  config.vm.synced_folder ENV['HOME'], ENV['HOME']

  config.vm.define "windows-2016" do |cfg|
    cfg.vm.box     = "StefanScherer/windows_2016_docker"
    cfg.vm.provision "shell", path: "scripts/vagrant/windows/create-docker-machine.ps1", args: "-machineHome #{ENV['HOME']} -machineName windows-2016"
  end

  config.vm.provider "vmware_fusion" do |v, override|
    v.gui = false
    v.memory = 2048
    v.cpus = 2
    v.enable_vmrun_ip_lookup = false
    v.linked_clone = true
    v.vmx["vhv.enable"] = "TRUE"
  end

  config.vm.provider "virtualbox" do |v, override|
    v.gui = false
    v.memory = 2048
    v.cpus = 2
    v.linked_clone = true
    override.vm.network :private_network, ip: "192.168.99.90", gateway: "192.168.99.1"
  end
end
