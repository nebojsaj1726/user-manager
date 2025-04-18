FROM node:23-slim AS builder

WORKDIR /app
COPY . .
RUN yarn install
RUN yarn build

FROM nginx:stable-alpine
COPY --from=builder /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
