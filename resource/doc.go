/*
Package resource provides a generic way to write testable concourse-ci resources.

Usage

To create the resource, you simply have to implement the relevant interface and hand it over to a `Handler`.
See the example for more details.

Input validation

The resource-framework makes use of the the go-playground/validator ( https://godoc.org/github.com/go-playground/validator )
to validate the input.

Templating

If you require templates in user defined input ( for example to build a link to the current job ), you should use the
functions provided by the template package

 */
package resource
