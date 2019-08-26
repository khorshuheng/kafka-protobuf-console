package reflection

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"
	"io/ioutil"
)

func LoadMessageDescriptor(fdPath string, msgName string) (*desc.MessageDescriptor, error) {
	bytes, err := ioutil.ReadFile(fdPath)
	if err != nil {
		return nil, err
	}

	var fileSet descriptor.FileDescriptorSet
	if err := proto.Unmarshal(bytes, &fileSet); err != nil {
		return nil, err
	}

	fds, err := desc.CreateFileDescriptorsFromSet(&fileSet)
	if err != nil {
		return nil, err
	}

	for _, fd := range fds {
		if fd := fd.FindMessage(msgName); fd != nil {
			return fd, nil
		} else {
			continue
		}
	}

	return nil, errors.New(fmt.Sprintf("Unable to find message named %s in file descriptor", msgName))
}