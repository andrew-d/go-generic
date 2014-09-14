# go-generic

This is a collection of generic-ish algorithms, datastructures and functions in
Go.  We achieve generics through code generation - specifically, using the
[gengen](https://github.com/joeshaw/gengen) library.  For example, to create a
version of `Merge` that is specialized for a given type, you can run:

		$ gengen chans/merge.go int | gofmt > my_merge.go

Which will give you versions of `Merge` and `MergeWithDone` that take and return
`int`s.
