---
title: Use dict.get with default instead of if-else
---

Join multiple with statements into a single one. Rule [SIM401](https://github.com/MartinThoma/flake8-simplify/issues/72) from [flake8-simplify](https://github.com/MartinThoma/flake8-simplify).

Caveat: the transformation is not run if either `$key` or `$default` have a function call,
as they would be called a different number of times in the new code.

```grit
engine marzano(0.1)
language python

`
if $key in $dict:
    $var = $dict[$key]
else:
    $var = $default
` => `$var = $dict.get($key, $default)` where {
	! $key <: contains call(),
	! $default <: contains call()
}
```

## Replace if-else with dict.get()

```python
if "my_key" in example_dict:
    thing = example_dict["my_key"]
else:
    thing = "default_value"


# Left as is

if f() in example_dict:
    thing = example_dict[f()]
else:
    thing = "default_value"

if "my_key" in example_dict:
    thing = example_dict["my_key"]
else:
    thing = "default_value" + f()

if "name" in d:
    name = d[name]
else:
    name = "foo"

if "name" in d:
    name = d["name"]
else:
    surname = "foo"
```

```python
thing = example_dict.get("my_key", "default_value")


# Left as is

if f() in example_dict:
    thing = example_dict[f()]
else:
    thing = "default_value"

if "my_key" in example_dict:
    thing = example_dict["my_key"]
else:
    thing = "default_value" + f()

if "name" in d:
    name = d[name]
else:
    name = "foo"

if "name" in d:
    name = d["name"]
else:
    surname = "foo"
```
