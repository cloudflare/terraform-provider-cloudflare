---
title: Jest to Vitest
tags: [migration, js]
---

Convert Jest tests to Vitest


```grit
engine marzano(0.1)
language js

pattern adjust_imports_vitest() {
	$body where {
		$body <: or {
			`$sym.$_($_)`,
			`$sym($_)`
		} where {
			$program <: program($statements),
			$mocha = `'mocha'`,
			$sym <: or {
				`test`,
				`expect`,
				`describe`,
				`it`,
				`beforeEach`,
				`afterEach`,
				`beforeAll`,
				`afterAll`,
				`jest`,
				`vi`
			} where {
				// Filter out methods which are imported from mocha
				$sym <: not imported_from(from=$mocha),
				// Filter out methods which are imported as default from any module
				$statements <: not contains or { `import $sym from $_` }
			},
			$source = `'vitest'`,
			$vi = `vi`,
			if ($sym <: not `jest`) { $sym <: ensure_import_from($source) } else {
				$vi <: ensure_import_from($source)
			}
		}
	}
}

pattern main_jest_to_vitest_migration() {
	file($body) where {
		$body <: contains bubble or {
			`JEST_WORKER_ID` => `VITEST_POOL_ID`,
			// Imports
			`import $_ from 'jest'` => .,
			// MODLE MOCKS
			// default literal mocking
			`jest.mock($module, $mockImplementation)` where {
				$mockImplementation <: or {
					arrow_function($body) where {
						// still waiting for default mock confirmation
						$body <: literal_value() => `({
              default: $body
            })`
					},
					`function($_) { $body }` where {
						$body <: contains `return $return` where {
							$return <: literal_value() => `{ default: $return }`
						}
					}
				}
			} => `vi.mock($module, $mockImplementation)`,
			`...jest.requireActual($module)` => `...(await vi.importActual($module))`,
			`jest.requireActual($module)` => `await vi.importActual($module)`,
			`jest.$sym` => `vi.$sym`,
			// done callback
			or {
				`$sym.$_($description, $callback)`,
				`$sym($description, $callback)`
			} where {
				$sym <: or {
					`it`,
					`test`
				},
				$callback <: or {
					// For matching `xxx => { ... }` & `(xxx) => { ... }`
					`$parameter => { $_ }` where { $parameter <: not . },
					// For matching `function (xxx) { ... }`
					`function($parameters) { $_ }` where { $parameters <: [$one, ...] }
				} => `() => new Promise($callback)`
			},
			// beforeAll, beforeEach, afterAll, afterEach hooks
			`$sym($callback)` where {
				and {
					$sym <: or {
						`beforeAll`,
						`beforeEach`,
						`afterAll`,
						`afterEach`
					},
					$callback <: arrow_function($body) where {
						$body <: call_expression() => `{ $body }`
					}
					// when `callback` is already a function, we don't need to do the migration
				}
			}
		}
	}
}

sequential {
	maybe main_jest_to_vitest_migration(),
	maybe contains adjust_imports_vitest()
}
```

## Convert assertions

```javascript
import { runCLI, mock, requireActual } from 'jest';

jest.mock('./some-path', () => 'hello');
jest.mock('./some-path', function () {
  doSomeSetups();
  return 'hello';
});
jest.mock('./some-path', () => ({ default: 'value' }));
jest.mock(
  './some-path',
  () =>
    function () {
      return 'hello';
    },
);

const { cloneDeep } = jest.requireActual('lodash/cloneDeep');
const { map } = jest.requireActual('lodash/map');
const { filter } = jest.requireActual('lodash/filter');
const originalModule = {
  ...jest.requireActual('../foo-bar-baz'),
  test: 1,
};
const currentWorkerId = JEST_WORKER_ID;

test.skip('test.skip should be processed', () => {
  expect('value').toBe('value');
});
it('should complete asynchronously', (done) => {
  expect('value').toBe('value');
  done();
});
it('should complete asynchronously', (finish) => {
  expect('value').toBe('value');
  finish();
});
it('should complete asynchronously', (done) => {
  expect('value').toBe('value');
  done();
});
it('should complete asynchronously', function (done) {
  expect('value').toBe('value');
  done();
});
it('should complete asynchronously', function (finish) {
  expect('value').toBe('value');
  finish();
});
test.skip('test.skip with done should be processed', (done) => {
  expect('value').toBe('value');
  done();
});
it('should be ignored', () => {
  expect('value').toBe('value');
});
it('should be ignored', function () {
  expect('value').toBe('value');
});

beforeAll(() => setActivePinia(createTestingPinia()));
beforeEach(() => setActivePinia(createTestingPinia()));
afterAll(() => setActivePinia(createTestingPinia()));
afterEach(() => setActivePinia(createTestingPinia()));
beforeAll(async () => {
  await expect('1').toBe('1');
  await expect('2').toBe('2');
});
beforeEach(() => {
  initializeApp();
  setDefaultUser();
});
afterEach(function () {
  initializeApp();
  setDefaultUser();
});
```

```javascript
import { vi, test, expect, it, beforeAll, beforeEach, afterAll, afterEach } from 'vitest';

vi.mock('./some-path', () => ({
  default: 'hello',
}));
vi.mock('./some-path', function () {
  doSomeSetups();
  return { default: 'hello' };
});
vi.mock('./some-path', () => ({ default: 'value' }));
vi.mock(
  './some-path',
  () =>
    function () {
      return 'hello';
    },
);

const { cloneDeep } = await vi.importActual('lodash/cloneDeep');
const { map } = await vi.importActual('lodash/map');
const { filter } = await vi.importActual('lodash/filter');
const originalModule = {
  ...(await vi.importActual('../foo-bar-baz')),
  test: 1,
};
const currentWorkerId = VITEST_POOL_ID;

test.skip('test.skip should be processed', () => {
  expect('value').toBe('value');
});
it('should complete asynchronously', () =>
  new Promise((done) => {
    expect('value').toBe('value');
    done();
  }));
it('should complete asynchronously', () =>
  new Promise((finish) => {
    expect('value').toBe('value');
    finish();
  }));
it('should complete asynchronously', () =>
  new Promise((done) => {
    expect('value').toBe('value');
    done();
  }));
it('should complete asynchronously', () =>
  new Promise(function (done) {
    expect('value').toBe('value');
    done();
  }));
it('should complete asynchronously', () =>
  new Promise(function (finish) {
    expect('value').toBe('value');
    finish();
  }));
test.skip('test.skip with done should be processed', () =>
  new Promise((done) => {
    expect('value').toBe('value');
    done();
  }));
it('should be ignored', () => {
  expect('value').toBe('value');
});
it('should be ignored', function () {
  expect('value').toBe('value');
});

beforeAll(() => {
  setActivePinia(createTestingPinia());
});
beforeEach(() => {
  setActivePinia(createTestingPinia());
});
afterAll(() => {
  setActivePinia(createTestingPinia());
});
afterEach(() => {
  setActivePinia(createTestingPinia());
});
beforeAll(async () => {
  await expect('1').toBe('1');
  await expect('2').toBe('2');
});
beforeEach(() => {
  initializeApp();
  setDefaultUser();
});
afterEach(function () {
  initializeApp();
  setDefaultUser();
});
```

## Import vi even with just vi.mock

```javascript
jest.mock('./some-path', () => ({ default: 'value' }));
```

```javascript
import { vi } from 'vitest';

vi.mock('./some-path', () => ({ default: 'value' }));
```

## Do not import methods which are imported from mocha/expect

```javascript
import expect from 'expect';
import { describe, test, before, beforeEach, after } from 'mocha';

describe('Repository component', () => {
  beforeEach(() => {});
  it('should complete asynchronously', () => {
    expect('value').toBe('value');
  });
});
```

```javascript
import expect from 'expect';
import { it } from 'vitest';

import { describe, test, before, beforeEach, after } from 'mocha';

describe('Repository component', () => {
  beforeEach(() => {});
  it('should complete asynchronously', () => {
    expect('value').toBe('value');
  });
});
```
