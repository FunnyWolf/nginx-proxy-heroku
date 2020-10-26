FROM nginx
COPY nginx.conf /etc/nginx/nginx.conf
COPY ssl.crt /etc/nginx/ssl.crt
COPY ssl.key /etc/nginx/ssl.key

