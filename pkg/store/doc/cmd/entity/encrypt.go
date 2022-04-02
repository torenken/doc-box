package entity

import (
	"io"
	"strings"

	"github.com/gtank/cryptopasta"

	"github.com/torenken/doc-box/pkg/store/doc/domain"
)

func Encrypt(doc domain.Document) (Document, error) {
	key, err := newEncryptionKey(doc.Id)
	if err != nil {
		return Document{}, err
	}

	encrypt, err := cryptopasta.Encrypt([]byte(doc.Name), key)
	if err != nil {
		return Document{}, err
	}

	return Document{
		Id:             doc.Id,
		Type:           doc.Type,
		Name:           encrypt,
		LifecycleState: doc.LifecycleState,
		CreationDate:   doc.CreationDate,
	}, nil
}

func Decrypt(doc Document) (domain.Document, error) {
	key, err := newEncryptionKey(doc.Id)
	if err != nil {
		return domain.Document{}, err
	}

	decrypt, err := cryptopasta.Decrypt(doc.Name, key)
	if err != nil {
		return domain.Document{}, err
	}

	return domain.Document{
		Id:             doc.Id,
		Type:           doc.Type,
		Name:           string(decrypt),
		LifecycleState: doc.LifecycleState,
		CreationDate:   doc.CreationDate,
	}, nil
}

func newEncryptionKey(value string) (*[32]byte, error) {
	key := [32]byte{}
	_, err := io.ReadFull(strings.NewReader(value), key[:])
	if err != nil {
		return nil, err
	}
	return &key, nil
}
