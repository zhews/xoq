FROM node:19.2-alpine AS build
WORKDIR /build
RUN wget -O- https://get.pnpm.io/v6.16.js | node - add --global pnpm
COPY . .
RUN pnpm install --frozen-lockfile
RUN pnpm build
FROM nginx:1.23-alpine AS serve
COPY default.conf.template /etc/nginx/templates/default.conf.template
COPY --from=build /build/dist /usr/share/nginx/html
