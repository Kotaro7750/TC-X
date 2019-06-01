ARG timezome="Asia/Tokyo"

FROM node:alpine AS build
WORKDIR /build
ADD ./tcx /build
RUN npm run product

FROM nginx:alpine
COPY --from=build /build/dist /var/www/html
COPY ./product.conf /etc/nginx/nginx.conf