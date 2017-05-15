# swagger-allof-test

This is to test swagger generating struct with embedded struct.

From the swagger.yaml file, an object B extending object A.
The generated Marshaler and Unmarshaler removed properties defined
in object B.

With the fix #847, the generated Mashaler and Unmarshaler look good.
But since https://github.com/go-swagger/go-swagger/commit/71225c4a7882ae49cd57e22292a58dff382908e8
the swagger was broken again.

