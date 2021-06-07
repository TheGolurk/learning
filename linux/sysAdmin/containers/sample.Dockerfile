# Create a webserver on ubuntu. It is a very cool getting started example

FROM ubuntu:18.04

RUN apt-get update
RUN apt-get install -y apache2
RUN echo "Welcome to my website" > /var/www/html/index.html
EXPOSE 80
