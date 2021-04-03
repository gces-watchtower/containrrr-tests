# MQTT

## URL Format
*mqtt://__`host`__:__`port`__?topic=__`topic`__*

## Optional parameters

You can optionally specify the __`disableTLS`__, __`clientId`__, __`username`__ and __`password`__ parameters in the URL:  
*mqtt://__`host`__:__`port`__?topic=__`topic`__&disableTLS=true&clientId=__`clientId`__&username=__`username`__&password:__`password`__*

## Parameters Description

* __Host__ - MQTT broker server hostname or IP address  (**Required**)  
 Default: *empty*  
 Aliases: `host`  

* __Port__ - MQTT server port, common ones are 8883 for TCP/TLS and 1883 for TCP  (**Required**) 
 Default: `8883`


* __Topic__ - Topic where the message is sent (**Required**) 
 Default: *empty*  
 Aliases: `Topic`  

* __DisableTLS__ - disable TLS/SSL Configurations  
 Default: `false`

* __ClientID__ - The client identifier (ClientId) identifies each MQTT client that connects to an MQTT  
 Default: *empty*
 Aliases: `clientId`

* __Username__ - name of the sender to auth  
 Default: *empty*
 Aliases: `clientId`

* __Password__ - authentication password or hash  
 Default: *empty*  
 Aliases: `password`

## Certificates to use TCP/TLS

To use TCP/TLS connection, it is necessary the files:
- Cerficate Authority: ca.crt
- Client Certificate: client.crt
- Client Key: client.key


## Configure TLS in mosquitto

Generate the certificates [mosquitto-tls](https://mosquitto.org/man/mosquitto-tls-7.html).
