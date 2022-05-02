# Data handling and processing

## `d.Set` calls for complex types

You should also include error handling for `d.Set` calls that are not using simple (string, number, boolean) values to catch any schema mismatches.

```go
func resourceExampleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    // ... snip
    // assuming `attr` is a map or list of types.


    if err := d.Set("my_attr", attr); err != nil {
      return fmt.Errorf("failed to set my_attr: %s", err)
    }

    return nil
}
```

## `expand*` style functions

Methods that start with `expand*` are used for taking a schema representation
and converting the data to a usable Golang struct.

This type of method will accept a `map[string]interface{}` or `[]interface{}`
and return a concrete type for use in API calls or interacting with the
underlying Go library.

```go
func expandThing(tfMap map[string]interface{}) *cloudflare.Thing {
    if tfMap == nil {
        return nil
    }

    apiObject := &cloudflare.Thing{}

    // ... nested attribute handling ...

    return apiObject
}

func expandThings(tfList []interface{}) []*cloudflare.Thing {
    if len(tfList) == 0 {
        return nil
    }

    var apiObjects []*cloudflare.Thing

    for _, tfMapRaw := range tfList {
        tfMap, ok := tfMapRaw.(map[string]interface{})

        if !ok {
            continue
        }

        apiObject := expandThing(tfMap)

        if apiObject == nil {
            continue
        }

        apiObjects = append(apiObjects, apiObject)
    }

    return apiObjects
}
```

## `flatten*` style functions

Methods that start with `flatten*` are used for taking a Golang struct and
converting the data to the Terraform schema representation.

This type of method will accept a Golang struct and usually returns a
`map[string]interface{}` or `[]interface{}` for use in `d.Set()` calls.

```go
func flattenThing(apiObject *cloudflare.Thing) map[string]interface{} {
    if apiObject == nil {
        return nil
    }

    tfMap := map[string]interface{}{}

    // ... nested attribute handling ...

    return tfMap
}

func flattenThings(apiObjects []*cloudflare.Thing) []interface{} {
    if len(apiObjects) == 0 {
        return nil
    }

    var tfList []interface{}

    for _, apiObject := range apiObjects {
        if apiObject == nil {
            continue
        }

        tfList = append(tfList, flattenThing(apiObject))
    }

    return tfList
}
```
