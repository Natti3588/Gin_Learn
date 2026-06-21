FROM nginx:1.30-alpine

WORKDIR /app

COPY conf.d/ .

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]