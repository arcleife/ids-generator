# Use a Node.js image
FROM node:latest

# Set the working directory
WORKDIR /app

# Copy the package.json and package-lock.json files
COPY package*.json ./

# Install dependencies
RUN npm install
RUN npm install -g serve

# Copy everything to the container
COPY . .

# Build
RUN npm run build

# Expose port
EXPOSE 4001

CMD [ "serve", "-s", "dist", "-p", "4001"]
