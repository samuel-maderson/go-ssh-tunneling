# GO-SSH-TUNNELING #

## HOW-TO ##

### Installation ###
Clone the project, then change the config.json according the type of host you will try to access.
For example, I'll connect to a ssh host, so my config.json will look like this.
```
{
    "user": "my-user",
    "key": "~/.ssh/my-ssh-key",
    "Remote": {
        "host": "172.31.29.131",
        "port": 22
    },
    "Local": {
        "host": "127.0.0.1",
        "port": 33306
    },
    "bastion": {
        "host": "172.31.75.252"
    }
}

```
### Description ###

* User section: Your username on bastion host.
* Key section: Your ssh private key that is related to your pubkey stored on bastion host.
* Remote section: Change with the ip and port address you want to connect.
* Local section: host attribute remains the same, port change to a local port that's available.
* bastion section: don't change.