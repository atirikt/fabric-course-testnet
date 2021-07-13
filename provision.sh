#!/bin/sh

# install git
install_git(){
    sudo apt install git -y
}
# install curl
install_curl(){
    sudo apt install curl -y
}

# install docker
install_docker(){
    sudo apt install \
        apt-transport-https \
        ca-certificates \
        gnupg-agent \
        software-properties-common -y
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
    sudo apt-key fingerprint 0EBFCD88
    sudo add-apt-repository \
        "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
        $(lsb_release -cs) \
        stable"
    sudo apt-get update
    sudo apt-get install docker-ce docker-ce-cli containerd.io docker-compose -y
    sudo usermod -aG docker $USER
    newgrp docker
    sudo chmod 666 /var/run/docker.sock
}
# install golang
install_golang(){
    curl -o "go.tar.gz" https://storage.googleapis.com/golang/go1.15.linux-amd64.tar.gz
    sudo chmod +x go.tar.gz
    sudo tar -C /usr/local -xzf "go.tar.gz"
}
# install node
install_node(){
    curl -sL https://deb.nodesource.com/setup_12.x | sudo -E bash -
    sudo apt install nodejs -y
}



install_git
install_curl
install_docker
install_golang
install_node

echo "export PATH=$PATH:/usr/local/go/bin" >> "$HOME/.bashrc"