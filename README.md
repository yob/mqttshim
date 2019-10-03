# MQTT Shim

How hard is it to accept all MQTT connections and print the received messages to stdout?

Turns out, not that hard.

## Running a server

Start an MQTT server like this (requires docker):

    ./auto/run

## Sending test messages

I'm on debian, and mosquitto is available by default:

    $ apt-get install mosquitto-clients

Then, send a message to localhost:

    $ mosquitto_pub -t foo -m 10 -q 0

## Helpful reading

* https://www.hivemq.com/blog/mqtt-essentials-part-1-introducing-mqtt/
* https://www.hivemq.com/blog/mqtt-essentials-part2-publish-subscribe/
* https://www.hivemq.com/blog/mqtt-essentials-part-3-client-broker-connection-establishment/
