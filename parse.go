package mtree

import (
	"bufio"
	"io"
	"strings"
)

func ParseSpec(r io.Reader) (*DirectoryHierarchy, error) {
	s := bufio.NewScanner(r)
	i := int(0)
	dh := DirectoryHierarchy{}
	for s.Scan() {
		str := s.Text()
		e := Entry{Pos: i}
		switch {
		case strings.HasPrefix(str, "#"):
			e.Raw = str
			if strings.HasPrefix(str, "#mtree") {
				e.Type = SignatureType
			} else {
				e.Type = CommentType
				// from here, the comment could be "# key: value" metadata
				// or a relative path hint
			}
		case str == "":
			e.Type = BlankType
			// nothing else to do here
		case strings.HasPrefix(str, "/"):
			e.Type = SpecialType
			// collapse any escaped newlines
			for {
				if strings.HasSuffix(str, `\`) {
					str = str[:len(str)-1]
					s.Scan()
					str += s.Text()
				} else {
					break
				}
			}
			// parse the options
			f := strings.Fields(str)
			e.Name = f[0]
			e.Keywords = f[1:]
		case len(strings.Fields(str)) > 0 && strings.Fields(str)[0] == "..":
			e.Type = DotDotType
			e.Raw = str
			// nothing else to do here
		case len(strings.Fields(str)) > 0:
			// collapse any escaped newlines
			for {
				if strings.HasSuffix(str, `\`) {
					str = str[:len(str)-1]
					s.Scan()
					str += s.Text()
				} else {
					break
				}
			}
			// parse the options
			f := strings.Fields(str)
			if strings.Contains(str, "/") {
				e.Type = FullType
			} else {
				e.Type = RelativeType
			}
			e.Name = f[0]
			e.Keywords = f[1:]
		default:
			// TODO(vbatts) log a warning?
			continue
		}
		dh.Entries = append(dh.Entries, e)
		i++
	}
	return &dh, s.Err()
}