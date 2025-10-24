FROM node:18-alpine

WORKDIR /app

# Copy package files terlebih dahulu
COPY package*.json ./
RUN npm install

# Copy seluruh source code
COPY . .

EXPOSE 3000

CMD ["npm", "run", "dev", "--", "--host"]