---
title: Fix `for` counter direction
tags: [bug, fix, good]
---

If a `for` counter moves in the wrong direction the loop will run infinitely. Mostly, an infinite `for` loop is a typo and causes a bug.


```grit
engine marzano(0.1)
language js

for_statement($condition, $increment, $initializer) where {
	or {
		and {
			$condition <: contains or {
				`$x < $_`,
				`$x <= $_`,
				`$_ > $x`,
				`$_ >= $x`
			},
			$increment <: contains { `$x--` => `$x++` }
		},
		and {
			$condition <: contains or {
				`$x > $_`,
				`$x >= $_`,
				`$_ < $x`,
				`$_ <= $x`
			},
			$increment <: contains { `$x++` => `$x--` }
		}
	}
}
```

## Transform `for` counter for `<`/`<=` directions

```javascript
for (var i = 0; i < 10; i--) {
  doSomething(i);
}
```

```typescript
for (var i = 0; i < 10; i++) {
  doSomething(i);
}
```

## Transform `for` counter for `>`/`>=` directions

```javascript
for (var i = 10; i >= 0; i++) {
  doSomething(i);
}
```

```typescript
for (var i = 10; i >= 0; i--) {
  doSomething(i);
}
```

## Transform counter for `<`/`<=` directions

```javascript
for (var i = 0; 10 > i; i--) {
  doSomething(i);
}
```

```typescript
for (var i = 0; 10 > i; i++) {
  doSomething(i);
}
```

## Do not change `for` counter

```javascript
for (var i = 0; i < 10; i++) {}
```
