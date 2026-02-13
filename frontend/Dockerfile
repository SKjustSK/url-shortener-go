# Stage 1: Build the React application
FROM node:18-alpine as builder

WORKDIR /app

# Copy package files and install dependencies
COPY package.json package-lock.json ./
RUN npm ci

# Copy the rest of the source code
COPY . .

# Copy the .env file so Vite can "bake" the variables into the build
COPY .env .

# Build the app (outputs to /app/dist)
RUN npm run build

# Stage 2: Serve with Nginx
FROM nginx:alpine

# Copy the build output from Stage 1 to Nginx's html folder
COPY --from=builder /app/dist /usr/share/nginx/html

# Add a simple Nginx config to handle React Router (redirects 404s to index.html)
# This fixes the "refresh on a subpage breaks the app" issue
RUN echo 'server { \
    listen 80; \
    root /usr/share/nginx/html; \
    index index.html; \
    location / { \
        try_files $uri $uri/ /index.html; \
    } \
}' > /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]