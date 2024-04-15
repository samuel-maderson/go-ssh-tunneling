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
        "host": "192.168.10.1",
        "port": 22
    },
    "Local": {
        "host": "127.0.0.1",
        "port": 33306
    },
    "bastion": {
        "host": "192.168.10.10"
    }
}

```
### Description ###

* User section: Your username on bastion host.
* Key section: Your ssh private key that is related to your pubkey stored on bastion host.
* Remote section: Change with the ip and port address you want to connect.
* Local section: host attribute remains the same, port change to a local port that's available.
* bastion section: don't change.


### Tunneling ###
Simple run the binary file as following:
```
# Starting tunnels
~/github.com/projects/go-ssh-tunneling (main)$ bin/go-ssh-tunneling-amd64-linux 
2024/04/15 15:55:35 [+] Starting Tunneling for host: db.example.com:3306
2024/04/15 15:55:35 [+] Listening on host: 127.0.0.1:53306
2024/04/15 15:55:35 
2024/04/15 15:55:35 [+] Starting Tunneling for host: 192.168.10.1:22
2024/04/15 15:55:35 [+] Listening on host: 127.0.0.1:33306
2024/04/15 15:55:35 


# Showing listening ports to illustrate
ss -t -n -l -p
State            Recv-Q           Send-Q                     Local Address:Port                      Peer Address:Port          Process                                  
LISTEN           0                128                            127.0.0.1:53306                          0.0.0.0:*              users:(("ssh",pid=5430,fd=5))           
LISTEN           0                128                            127.0.0.1:33306                          0.0.0.0:*              users:(("ssh",pid=5429,fd=5))           
LISTEN           0                4096                       127.0.0.53%lo:53                             0.0.0.0:*                                                      
LISTEN           0                5                              127.0.0.1:631                            0.0.0.0:*                                                      
LISTEN           0                128                                [::1]:53306                             [::]:*              users:(("ssh",pid=5430,fd=4))           
LISTEN           0                128                                [::1]:33306                             [::]:*              users:(("ssh",pid=5429,fd=4))           
LISTEN           0                5                                  [::1]:631                               [::]:*  
```