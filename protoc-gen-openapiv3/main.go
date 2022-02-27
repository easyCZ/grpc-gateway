package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/internal/codegenerator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	f := os.Stdin
	req, err := codegenerator.ParseRequest(f)
	if err != nil {
		log.Fatalf("failed to read request")
	}

	ioutil.WriteFile("test.foo", []byte(fmt.Sprintf("%v", req)), 0644)

	emitResp(nil)
}

func emitFiles(out []*descriptor.ResponseFile) {
	files := make([]*pluginpb.CodeGeneratorResponse_File, len(out))
	for idx, item := range out {
		files[idx] = item.CodeGeneratorResponse_File
	}
	resp := &pluginpb.CodeGeneratorResponse{File: files}
	codegenerator.SetSupportedFeaturesOnCodeGeneratorResponse(resp)
	emitResp(resp)
}

func emitResp(resp *pluginpb.CodeGeneratorResponse) {
	buf, err := proto.Marshal(resp)
	if err != nil {
		glog.Fatal(err)
	}
	if _, err := os.Stdout.Write(buf); err != nil {
		glog.Fatal(err)
	}
}
