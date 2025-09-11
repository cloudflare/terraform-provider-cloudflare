---
tags: [hidden, linting, best-practice, recommended, ai, flaky]
---

# No Commented Out Code

Please [don't commit commented out code](https://kentcdodds.com/blog/please-dont-commit-commented-out-code).

```grit
engine marzano(0.1)
language js

file($body) where {
	$comments = [],
	$body <: contains bubble($comments) comment() as $comment where {
		$comments += $comment
	},
	$blocks = group_blocks(target=$comments),
	$blocks <: some bubble $block where {
		$joined = join($block, `\n`),
		$joined <: not or {
			includes "@ts-ignore",
			includes "@ts-expect-error",
			includes "eslint-disable",
			includes "eslint-disable-next-line",
			includes "eslint-disable-line",
			includes "biome-ignore",
			r"(.+)-ignore"
		},
		$joined <: ai_is("commented out code that is valid JavaScript, not a descriptive comment", examples=[
			"// console.log(name);",

			"// for (const name of names) { console.log(name); }",

			`// foo();
// bar();`,

			`/**
  * for (const item of items) {
  *   console.log(item);
  * }
  */`,
		], counter_examples=[
			"// Read the user's name from the database",

			`/**
 * This is a comment that describes the code below.
 * It is not commented out code.
*/`,
		]),
		// Remove the block
		$block <: some bubble $comment => .
	}
}
```

## Base case, single-line comments

```js
var increment = function (i) {
  console.log("hi")
  // const answer = 54;
  // const wow = 42;
  const answer = 42;
  return i + 1;
};

var remember = function (me) {
  // this is a comment, without a test
  this.you = me;
};

const blocks;

var sumToValue = function (x, y) {
  function Value(v) {
    this.value = v;
  }
  return new Value(x + y);
};

var times = (x, y) => {
  return x * y;
};
```

```js
var increment = function (i) {
  console.log("hi")
  const answer = 42;
  return i + 1;
};

var remember = function (me) {
  // this is a comment, without a test
  this.you = me;
};

const blocks;

var sumToValue = function (x, y) {
  function Value(v) {
    this.value = v;
  }
  return new Value(x + y);
};

var times = (x, y) => {
  return x * y;
};
```

## Doesn't match on valid files

```js
// Don't sample logging calls
if (name === 'grpc.google.logging.v2.LoggingServiceV2/WriteLogEntries') return RATE_DROP;
```

## Handles block comments too

Block comments don't currently parse correctly, see https://github.com/getgrit/rewriter/issues/7731.

```js
/** See sdk_proxy for how stdlib calls are intercepted and the workflow ID is injected. */
export const createSdkActivities = () => {
  /**
   * const stdlib = new Proxy({}, {});
   */
  return new Proxy(stdlib, {});
};
```

```js
/** See sdk_proxy for how stdlib calls are intercepted and the workflow ID is injected. */
export const createSdkActivities = () => {
  return new Proxy(stdlib, {});
};
```

## Doesn't remove eslint or typescript ignores

```js
// const foo = 9;
if (name === 'grpc.google.logging.v2.LoggingServiceV2/WriteLogEntries') return RATE_DROP;
// @ts-expect-error This is not useful
const foo: string = 9;
// @ts-ignore This is not useful
const foo: string = 9;
// eslint-disable-next-line no-use-before-define
const foo = 9;
```

```js
if (name === 'grpc.google.logging.v2.LoggingServiceV2/WriteLogEntries') return RATE_DROP;
// @ts-expect-error This is not useful
const foo: string = 9;
// @ts-ignore This is not useful
const foo: string = 9;
// eslint-disable-next-line no-use-before-define
const foo = 9;
```
