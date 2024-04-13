# Kafka_Messaging_System

## Overview

I recently got very interested in event streaming architecture and its wide adoption in banking, healthcare, government etc. Naturally, my first stop was learning about Apache Kafka. It's fascinating how fast, scalable and reliable this piece of software actually is. 

That's why I've decided to seek some lessons and tutorials to create my own (very amateurish and simple) event streaming platform. 

So, a very high level overview of the Kafka architecture is this:

![alt text](<assets/Screenshot from 2024-04-13 14-32-44.png>)

_Credit to: Finematics_

On this diagram you don't see Zookeeper but it's getting deprecated in Kafka 4.0.
I won't go into detail explaining the software - I am far, far, far from an expert.

## Personal Notes

* This was my first project with Docker Compose but it was already a build .yml file of Kafka and Zookeeper so it didn't take much effort.

* The screenshot above shows my producer (**on the left**), my consumer (**on the right**). The middle terminal is used to manually send a message to test the functionality of the whole thing. As you can see it works and on both sides of the frame there's a confirmation of message received and stored.

![alt text](<assets/Screenshot from 2024-04-13 02-06-04.png>)


