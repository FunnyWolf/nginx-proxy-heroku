FROM nginx
COPY ssl.crt /etc/nginx/ssl.crt
COPY ssl.key /etc/nginx/ssl.key
ADD configure.sh /configure.sh
RUN chmod +x /configure.sh
CMD /configure.sh
