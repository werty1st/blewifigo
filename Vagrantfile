#https://sonnguyen.ws/connect-usb-from-virtual-machine-using-vagrant-and-virtualbox/




# UUID:               05df4a26-40b7-412a-a00e-b2964169ad98
# VendorId:           0x2357 (2357)
# ProductId:          0x0109 (0109)
# Revision:           2.0 (0200)
# Port:               1
# USB version/speed:  0/High
# Manufacturer:       Realtek
# Product:            802.11n NIC
# SerialNumber:       00e04c000001
# Address:            p=0x0109;v=0x2357;s=0x00000579e7e22baf;l=0x14210000
# Current State:      Available

# UUID:               5d753cc9-7448-4f13-818b-9ae96276ceed
# VendorId:           0x050d (050D)
# ProductId:          0x065a (065A)
# Revision:           1.18 (0118)
# Port:               2
# USB version/speed:  0/Full
# Manufacturer:       Broadcom Corp
# Product:            BCM20702A0
# SerialNumber:       000272DA209C
# Address:            p=0x065a;v=0x050d;s=0x00000579ea9104ea;l=0x14220000
# Current State:      Available


#source: https://github.com/hashicorp/vagrant/issues/5774
def usbfilter_exists(vendor_id, product_id)
    #
    # Determine if a usbfilter with the provided Vendor/Product ID combination
    # already exists on this VM.
    #
    # TODO: Use a more reliable way of retrieving this information.
    #
    # NOTE: The "machinereadable" output for usbfilters is more
    #       complicated to work with (due to variable names including
    #       the numeric filter index) so we don't use it here.
    #
    machine_id_filepath = ".vagrant/machines/default/virtualbox/id"

    if not File.exists? machine_id_filepath
      # VM hasn't been created yet.
      return false
    end

    vm_info = `VBoxManage showvminfo $(<#{machine_id_filepath})`
    filter_match = "VendorId:         #{vendor_id}\nProductId:        #{product_id}\n"
    return vm_info.include? filter_match
end

def better_usbfilter_add(vb, vendor_id, product_id, filter_name)
    #
    # This is a workaround for the fact VirtualBox doesn't provide
    # a way for preventing duplicate USB filters from being added.
    #
    # TODO: Implement this in a way that it doesn't get run multiple
    #       times on each Vagrantfile parsing.
    #
    if not usbfilter_exists(vendor_id, product_id)
        vb.customize ["usbfilter", "add", "0",
                    "--target", :id,
                    "--name", filter_name,
                    "--vendorid", vendor_id,
                    "--productid", product_id
                    ]
    end
end


Vagrant.configure("2") do | config |
    config.vm.box = "ubuntu/bionic64"

    # not needed for VSC ssh remote if vagrant key is added to agent
    # ssh_pub_key = File.readlines("#{Dir.home}/.ssh/id_rsa.pub").first.strip
    # config.vm.provision 'shell', inline: 'mkdir -p /root/.ssh'
    # config.vm.provision 'shell', inline: "echo #{ssh_pub_key} >> /root/.ssh/authorized_keys"
    # config.vm.provision 'shell', inline: "echo #{ssh_pub_key} >> /home/vagrant/.ssh/authorized_keys", privileged: false

    config.vm.provider "virtualbox" do |vb|
        vb.memory = "2404"
        vb.cpus = "2"
        vb.name = "ble-dev"
        vb.customize ["modifyvm", :id, "--usb", "on"]
        vb.customize ["modifyvm", :id, "--usbehci", "on"]
        better_usbfilter_add(vb, "0x2357", "0x0109", "wifi")
        better_usbfilter_add(vb, "0x050d", "0x065a", "blue")

    end    

    #ordner
    #/vegrant

end

