FROM node:alpine AS build

ARG timezome="Asia/Tokyo"
ARG PRODUCT_IP
WORKDIR /build
ADD ./tcx /build
ENV TZ=${timezome}
RUN echo "VUE_APP_SERVER_URL=http://$PRODUCT_IP:8888" > /build/.env.production
RUN npm install
RUN npm run product

FROM nginx:alpine
ENV TZ=${timezome}
COPY --from=build /build/dist /var/www/html
COPY ./product.conf /etc/nginx/nginx.conf
EXPOSE 80