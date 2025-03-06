FROM node:23-alpine3.19 AS builder

WORKDIR /app

COPY package.json /app/package.json

RUN npm install

COPY . .

RUN npm run build

FROM nginx:stable 

COPY --from=builder /app/dist /usr/share/nginx/html

COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80
EXPOSE 443

CMD ["nginx", "-g", "daemon off;"]
