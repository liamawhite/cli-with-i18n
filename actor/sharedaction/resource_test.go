package sharedaction_test

import (
	"archive/zip"
	"io/ioutil"
	"os"
	"path/filepath"

	"code.cloudfoundry.org/ykk"
	. "github.com/liamawhite/cli-with-i18n/actor/sharedaction"
	"github.com/liamawhite/cli-with-i18n/actor/sharedaction/sharedactionfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Resource Actions", func() {
	var (
		fakeConfig *sharedactionfakes.FakeConfig
		actor      *Actor
		srcDir     string
	)

	BeforeEach(func() {
		fakeConfig = new(sharedactionfakes.FakeConfig)
		actor = NewActor(fakeConfig)

		var err error
		srcDir, err = ioutil.TempDir("", "resource-actions-test")
		Expect(err).ToNot(HaveOccurred())

		subDir := filepath.Join(srcDir, "level1", "level2")
		err = os.MkdirAll(subDir, 0777)
		Expect(err).ToNot(HaveOccurred())

		err = ioutil.WriteFile(filepath.Join(subDir, "tmpFile1"), []byte("why hello"), 0600)
		Expect(err).ToNot(HaveOccurred())

		err = ioutil.WriteFile(filepath.Join(srcDir, "tmpFile2"), []byte("Hello, Binky"), 0600)
		Expect(err).ToNot(HaveOccurred())

		err = ioutil.WriteFile(filepath.Join(srcDir, "tmpFile3"), []byte("Bananarama"), 0600)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		Expect(os.RemoveAll(srcDir)).ToNot(HaveOccurred())
	})

	Describe("GatherArchiveResources", func() {
		// tests are under resource_unix_test.go and resource_windows_test.go
	})

	Describe("GatherDirectoryResources", func() {
		// tests are under resource_unix_test.go and resource_windows_test.go
	})

	Describe("ZipArchiveResources", func() {
		var (
			archive    string
			resultZip  string
			resources  []Resource
			executeErr error
		)

		BeforeEach(func() {
			tmpfile, err := ioutil.TempFile("", "zip-archive-resources")
			Expect(err).ToNot(HaveOccurred())
			defer tmpfile.Close()
			archive = tmpfile.Name()

			err = zipit(srcDir, archive, "")
			Expect(err).ToNot(HaveOccurred())
		})

		JustBeforeEach(func() {
			resultZip, executeErr = actor.ZipArchiveResources(archive, resources)
		})

		AfterEach(func() {
			Expect(os.RemoveAll(archive)).ToNot(HaveOccurred())
			Expect(os.RemoveAll(resultZip)).ToNot(HaveOccurred())
		})

		Context("when the files have not been changed since scanning them", func() {
			BeforeEach(func() {
				resources = []Resource{
					{Filename: "/"},
					{Filename: "/level1/"},
					{Filename: "/level1/level2/"},
					{Filename: "/level1/level2/tmpFile1", SHA1: "9e36efec86d571de3a38389ea799a796fe4782f4"},
					{Filename: "/tmpFile2", SHA1: "e594bdc795bb293a0e55724137e53a36dc0d9e95"},
					// Explicitly skipping /tmpFile3
				}
			})

			It("zips the file and returns a populated resources list", func() {
				Expect(executeErr).ToNot(HaveOccurred())

				Expect(resultZip).ToNot(BeEmpty())
				zipFile, err := os.Open(resultZip)
				Expect(err).ToNot(HaveOccurred())
				defer zipFile.Close()

				zipInfo, err := zipFile.Stat()
				Expect(err).ToNot(HaveOccurred())

				reader, err := ykk.NewReader(zipFile, zipInfo.Size())
				Expect(err).ToNot(HaveOccurred())

				Expect(reader.File).To(HaveLen(5))
				Expect(reader.File[0].Name).To(Equal("/"))
				Expect(reader.File[1].Name).To(Equal("/level1/"))
				Expect(reader.File[2].Name).To(Equal("/level1/level2/"))
				Expect(reader.File[3].Name).To(Equal("/level1/level2/tmpFile1"))
				Expect(reader.File[4].Name).To(Equal("/tmpFile2"))

				expectFileContentsToEqual(reader.File[3], "why hello")
				expectFileContentsToEqual(reader.File[4], "Hello, Binky")

				for _, file := range reader.File {
					Expect(file.Method).To(Equal(zip.Deflate))
				}
			})
		})

		Context("when the files have changed since the scanning", func() {
			BeforeEach(func() {
				resources = []Resource{
					{Filename: "/"},
					{Filename: "/level1/"},
					{Filename: "/level1/level2/"},
					{Filename: "/level1/level2/tmpFile1", SHA1: "9e36efec86d571de3a38389ea799a796fe4782f4"},
					{Filename: "/tmpFile2", SHA1: "e594bdc795bb293a0e55724137e53a36dc0d9e95"},
					{Filename: "/tmpFile3", SHA1: "i dunno, 7?"},
				}
			})

			It("returns an FileChangedError", func() {
				Expect(executeErr).To(Equal(FileChangedError{Filename: "/tmpFile3"}))
			})
		})
	})

	Describe("ZipDirectoryResources", func() {
		var (
			resultZip  string
			resources  []Resource
			executeErr error
		)

		JustBeforeEach(func() {
			resultZip, executeErr = actor.ZipDirectoryResources(srcDir, resources)
		})

		AfterEach(func() {
			Expect(os.RemoveAll(resultZip)).ToNot(HaveOccurred())
		})

		Context("when the files have not been changed since scanning them", func() {
			BeforeEach(func() {
				resources = []Resource{
					{Filename: "level1"},
					{Filename: "level1/level2"},
					{Filename: "level1/level2/tmpFile1", SHA1: "9e36efec86d571de3a38389ea799a796fe4782f4"},
					{Filename: "tmpFile2", SHA1: "e594bdc795bb293a0e55724137e53a36dc0d9e95"},
					{Filename: "tmpFile3", SHA1: "f4c9ca85f3e084ffad3abbdabbd2a890c034c879"},
				}
			})

			It("zips the file and returns a populated resources list", func() {
				Expect(executeErr).ToNot(HaveOccurred())

				Expect(resultZip).ToNot(BeEmpty())
				zipFile, err := os.Open(resultZip)
				Expect(err).ToNot(HaveOccurred())
				defer zipFile.Close()

				zipInfo, err := zipFile.Stat()
				Expect(err).ToNot(HaveOccurred())

				reader, err := ykk.NewReader(zipFile, zipInfo.Size())
				Expect(err).ToNot(HaveOccurred())

				Expect(reader.File).To(HaveLen(5))
				Expect(reader.File[0].Name).To(Equal("level1/"))
				Expect(reader.File[1].Name).To(Equal("level1/level2/"))
				Expect(reader.File[2].Name).To(Equal("level1/level2/tmpFile1"))
				Expect(reader.File[3].Name).To(Equal("tmpFile2"))
				Expect(reader.File[4].Name).To(Equal("tmpFile3"))

				expectFileContentsToEqual(reader.File[2], "why hello")
				expectFileContentsToEqual(reader.File[3], "Hello, Binky")
				expectFileContentsToEqual(reader.File[4], "Bananarama")

				for _, file := range reader.File {
					Expect(file.Method).To(Equal(zip.Deflate))
				}
			})
		})

		Context("when the files have changed since the scanning", func() {
			BeforeEach(func() {
				resources = []Resource{
					{Filename: "level1"},
					{Filename: "level1/level2"},
					{Filename: "level1/level2/tmpFile1", SHA1: "9e36efec86d571de3a38389ea799a796fe4782f4"},
					{Filename: "tmpFile2", SHA1: "e594bdc795bb293a0e55724137e53a36dc0d9e95"},
					{Filename: "tmpFile3", SHA1: "i dunno, 7?"},
				}
			})

			It("returns an FileChangedError", func() {
				Expect(executeErr).To(Equal(FileChangedError{Filename: filepath.Join(srcDir, "tmpFile3")}))
			})
		})
	})
})

func expectFileContentsToEqual(file *zip.File, expectedContents string) {
	reader, err := file.Open()
	Expect(err).ToNot(HaveOccurred())
	defer reader.Close()

	body, err := ioutil.ReadAll(reader)
	Expect(err).ToNot(HaveOccurred())

	Expect(string(body)).To(Equal(expectedContents))
}
