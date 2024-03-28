FROM nginx

ENV TZ Asia/Shanghai

WORKDIR /app/dist

COPY ./dist  /app/dist/

COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80
