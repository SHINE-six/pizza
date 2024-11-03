# syntax=docker/dockerfile:1

# Comments are provided throughout this file to help you get started.
# If you need more help, visit the Dockerfile reference guide at
# https://docs.docker.com/go/dockerfile-reference/

# Want to help us make this template better? Share your feedback here: https://forms.gle/ybq9Krt8jtBL3iCk7

ARG NODE_VERSION=18.17.0

################################################################################
# Use node image for base image for all stages.
FROM node:${NODE_VERSION}-alpine AS base

# Set working directory for all build stages.
WORKDIR /usr/src/app


################################################################################
# Create a stage for installing production dependecies.
FROM base AS deps

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies.
RUN npm install

################################################################################
# Create a stage for building the application.
FROM deps AS final

# Copy the rest of the source files into the image.
COPY . .

# Install production dependencies
RUN npm install --only=production

# Expose the port that the application listens on.
EXPOSE 10070

# Run the application.
CMD ["npm", "start"]
