package main

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	tree "github.com/xlab/treeprint"
	// "github.com/davecgh/go-spew/spew"
	"strings"
)

var (
	cwd              string
	verbose          bool
	directory        string
	filterDuplicates bool
)

type Certificate struct {
	*x509.Certificate
	Path string
}

func (c *Certificate) Sha256Sum() string {
	h := sha256.New()
	h.Write(c.Raw)
	return hex.EncodeToString(h.Sum(nil))[0:5]
}

func (c *Certificate) ID() string {
	return hex.EncodeToString(c.SubjectKeyId)
}

func (c *Certificate) AuthorityID() string {
	return hex.EncodeToString(c.AuthorityKeyId)
}

type Tree struct {
	tree.Tree
}

func (t *Tree) Format(certificate *Certificate) string {
	ca := ""
	if certificate.IsCA {
		ca = " [CA]"
	}

	return fmt.Sprintf(
		"%s%s %s -> %s",
		certificate.Sha256Sum(),
		ca,
		certificate.Subject.CommonName,
		ShortPath(certificate.Path),
	)
}

func (t *Tree) Add(certificate *Certificate) *Tree {
	return &Tree{t.Tree.AddBranch(t.Format(certificate))}
}

func ShortPath(path string) string {
	parts := strings.Split(path, string(filepath.Separator))
	if len(parts) > 3 {
		for k, part := range parts[:len(parts)-1] {
			parts[k] = part[:1]
		}
	}
	return strings.Join(parts, string(filepath.Separator))
}

func MakeCertificates(directory string) ([]*Certificate, error) {
	var (
		certificates = []*Certificate{}
		walker       = func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				buf, err := ioutil.ReadFile(path)
				if err != nil {
					print(err)
					return nil
				}

				relPath, err := filepath.Rel(directory, path)
				if err != nil {
					return err
				}

				block, _ := pem.Decode(buf)
				if block == nil {
					print(fmt.Errorf("could not decode pem from %s", path))
					return nil
				}
				certificate, err := x509.ParseCertificate(block.Bytes)
				if err != nil {
					print(err)
					return nil
				}
				certificates = append(certificates, &Certificate{
					Certificate: certificate,
					Path:        relPath,
				})
			}

			return nil
		}
	)
	return certificates, filepath.Walk(directory, walker)

}

func makeTree(t *Tree, certificates []*Certificate, authorityID string, processed map[string]struct{}) *Tree {
	for _, certificate := range certificates {
		if _, exists := processed[certificate.Path]; exists {
			continue
		}
		if certificate.AuthorityID() != authorityID {
			continue
		}

		if filterDuplicates {
			processed[certificate.Path] = struct{}{}
		}
		if certificate.IsCA {
			makeTree(t.Add(certificate), certificates, certificate.ID(), processed)
		} else {
			t.Add(certificate)
		}
	}

	return t
}

func MakeTree(t *Tree, certificates []*Certificate) *Tree {
	return makeTree(t, certificates, "", map[string]struct{}{})
}

func init() {
	var err error

	cwd, err = os.Getwd()
	if err != nil {
		fatal(err)
	}

	flag.StringVar(&directory, "d", cwd, "root directory where search certificates")
	flag.BoolVar(&verbose, "v", false, "verbose mode")
	flag.BoolVar(&filterDuplicates, "f", false, "filter duplicates while building a tree")
	flag.Parse()

	directory, err = filepath.Abs(directory)
	if err != nil {
		fatal(err)
	}
}

func print(v ...interface{}) {
	if verbose {
		log.Println(v...)
	}
}

func fatal(v ...interface{}) {
	log.Fatal(v...)
}

func main() {
	certificates, err := MakeCertificates(directory)
	if err != nil {
		fatal(err)
	}

	fmt.Println(MakeTree(&Tree{tree.New()}, certificates).String())
}
