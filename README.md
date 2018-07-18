# swagger-allof-test

This is to test swagger generating struct with embedded struct.

From the swagger.yaml file, an object B extending object A.
The generated Marshaler and Unmarshaler removed properties defined
in object B.

With the fix #847, the generated Mashaler and Unmarshaler look good.
But since https://github.com/go-swagger/go-swagger/commit/71225c4a7882ae49cd57e22292a58dff382908e8
the swagger was broken again.

See the [diff](https://github.com/yan-zhuang/swagger-allof-test/commit/9cc8549d59756eb813fac7fbe85c45412073d928) that comapre fix 847 with 7125c.

# Update
The issue was closed for the geerator works well when specify `allof` correctly.
(**See swagger.yaml file in *src/correct* folder**)

See https://github.com/go-swagger/go-swagger/issues/1042

# Version for testing
```sh
$ swagger version
version: 0.15.0
commit: 84da4f998602ef20ebd60a576756996162a1ca66
```

# Diff between incorrect and correct specs
```sh
diff --git a/src/incorrect/swagger.yaml b/src/correct/swagger.yaml
index 4983902..4aa4f7a 100644
--- a/src/incorrect/swagger.yaml
+++ b/src/correct/swagger.yaml
@@ -26,16 +26,16 @@ definitions:
         type: string
 
   B:
-    type: object
     allOf:
       - "$ref": "#/definitions/A"
-    required:
-      - f3
-      - f4
-    properties:
-      f3: 
-        type: string
-      f4: 
-        type: array
-        items:
-          type: string
+      - type: object
+        required:
+          - f3
+          - f4
+        properties:
+          f3: 
+            type: string
+          f4: 
+            type: array
+            items:
+              type: string
```

# Diff between generated files
```sh
diff --git a/src/incorrect/models/b.go b/src/correct/models/b.go
index a79e3db..b59ed0a 100644
--- a/src/incorrect/models/b.go
+++ b/src/correct/models/b.go
@@ -36,12 +36,26 @@ func (m *B) UnmarshalJSON(raw []byte) error {
 	}
 	m.A = aO0
 
+	// AO1
+	var dataAO1 struct {
+		F3 *string `json:"f3"`
+
+		F4 []string `json:"f4"`
+	}
+	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
+		return err
+	}
+
+	m.F3 = dataAO1.F3
+
+	m.F4 = dataAO1.F4
+
 	return nil
 }
 
 // MarshalJSON marshals this object to a JSON structure
 func (m B) MarshalJSON() ([]byte, error) {
-	_parts := make([][]byte, 0, 1)
+	_parts := make([][]byte, 0, 2)
 
 	aO0, err := swag.WriteJSON(m.A)
 	if err != nil {
@@ -49,6 +63,22 @@ func (m B) MarshalJSON() ([]byte, error) {
 	}
 	_parts = append(_parts, aO0)
 
+	var dataAO1 struct {
+		F3 *string `json:"f3"`
+
+		F4 []string `json:"f4"`
+	}
+
+	dataAO1.F3 = m.F3
+
+	dataAO1.F4 = m.F4
+
+	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
+	if errAO1 != nil {
+		return nil, errAO1
+	}
+	_parts = append(_parts, jsonDataAO1)
+
 	return swag.ConcatJSON(_parts...), nil
 }
 
```
