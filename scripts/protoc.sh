GOPATH=$(go env GOPATH)
mkdir temp && cd temp
wget https://github.com/protocolbuffers/protobuf/releases/download/v29.3/protoc-29.3-linux-x86_64.zip
unzip protoc-29.3-linux-x86_64.zip

mv bin/protoc $GOPATH/bin
cd ../
rm -rf temp