# PobreBox

<img src="https://raw.githubusercontent.com/marcusrafael/distribution-platforms/master/assignment-01/images/pobrebox.png">

## Installation

### Server VM

Install Go:

```
sudo add-apt-repository ppa:longsleep/golang-backports
```
```
sudo apt update
```
```
sudo apt install golang-1.10-go
```

Install dependency:

```
go get github.com/streadway/amqp
```

Create the directory:

```
mkdir /home/ubuntu/pobrebox/
```

Open ports on security group (inbound):

```
TCP: 80 22 1337 5672
```

### Storage VM

Install dependency:

```
go get github.com/streadway/amqp
```

Install Apache Web Server:

```
sudo apt install apache2
```

Create a cron job:

```
sudo crontab -e
```
```
* * * * * /usr/bin/go run /home/ubuntu/distribution-platforms/assignment-01/pobrebox/socket-tcp-client.go >/dev/null 2>&1
```

Create the directory:

```
mkdir /home/ubuntu/pobrebox-storage/
```

Open ports on security group (inbound):

```
TCP: 80 22 1337 5672
```

### Message Queue VM

Install RabbitMQ 3.5.7:

```
sudo apt install rabbit-server
```

Create user on RabbitMQ:

```
sudo rabbitmqctl add_user cloud cloud
```

Set user as administrator:

```
sudo rabbitmqctl set_user_tags cloud administrator
```

Set user permissions:

```
sudo rabbitmqctl set_permissions -p / cloud ".*" ".*" ".*"
```

Open ports on security group (inbound):

```
TCP: 22 443 1337 5672 12672
```
