# Create a webserver on ubuntu

FROM ubuntu:18.04

RUN apt-get update
RUN apt-get istall -y apache2
RUN echo "Welcome to my website" > /var/www/html/index.html
EXPOSE 80
