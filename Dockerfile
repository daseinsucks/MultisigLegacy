FROM node:alpine

WORKDIR /multisig

# Install build deps
RUN apk --no-cache --update-cache add alpine-sdk libusb-dev python3 py3-pip unzip
RUN npm install -g grunt

# Do the thing
COPY . .
# RUN npm ci; exit 0
WORKDIR /multisig/dapp
# RUN npm run postinstall
RUN unzip -o node1.zip && unzip -o node2.zip && unzip -o node3.zip
RUN unzip -o bundles.zip

EXPOSE 8282
# Run frontend
ENTRYPOINT [ "grunt" ]