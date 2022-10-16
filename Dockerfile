FROM node:alpine

WORKDIR /multisig

# Install build deps
RUN apk --no-cache --update-cache add alpine-sdk python3 py3-pip unzip
RUN npm install -g grunt

# Do the thing
COPY . .
RUN cd dapp && unzip node1.zip && unzip node2.zip && unzip node3.zip

ENTRYPOINT [ "npm", "start" ]