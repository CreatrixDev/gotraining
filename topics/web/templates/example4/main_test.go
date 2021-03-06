// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Test for the sample program to show how to use create, parse
// and execute a template with simple data processing. This
// example uses a struct type value and generates HTML markup.
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestExec(t *testing.T) {

	// Create a bytes Buffer for our Writer.
	var bb bytes.Buffer

	// Execute the template putting the output
	// into our bytes buffer.
	if err := Exec(&bb); err != nil {
		t.Fatal(err)
	}

	// Validate we received the correct version.
	got := strings.TrimSpace(bb.String())
	want := strings.TrimSpace(`
<a href="/foo?email=mark%40example.com">mark@example.com</a>
<script>
	window.user = {"name":"Mark","email":"mark@example.com"};
</script>`)
	if got != want {
		t.Log("Wanted:", want)
		t.Log("Got   :", got)
		t.Fatal("Mismatch")
	}
}
