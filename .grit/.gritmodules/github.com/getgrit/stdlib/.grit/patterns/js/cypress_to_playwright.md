---
title: Convert Cypress to Playwright
tags: [migration]
---

Migrate from Cypress to Playwright.

```grit
engine marzano(0.1)
language js

pattern convert_cypress_assertions() {
	or {
		`expect($arg).to.not.be.null` => `expect($arg).not.toBeNull()`,
		`expect($arg).to.not.be.undefined` => `expect($arg).not.toBeUndefined()`,
		`expect($arg).to.include($item)` => `expect($arg).toContain($item)`,
		`expect($arg).to.eql($item)` => `expect($arg).toEqual($item)`,
		`$locator.should($cond1, $cond2)` as $should where {
			$pw_cond = "",
			$cond1 <: or {
				`'contain'` where { $pw_cond += `toContainText($cond2)` },
				`'have.attr'` where { $pw_cond += `toHaveAttribute($cond2)` }
			},
			$should => `await expect($locator).$pw_cond`
		},
		`$locator.should($condition)` as $should where {
			$condition <: bubble or {
				`'exist'` => `toBeAttached()`,
				`'not.exist'` => `not.toBeAttached()`
			},
			$should => `await expect($locator).$condition`
		}
	}
}

pattern convert_cypress_queries() {
	or {
		`cy.visit($loc)` => `await page.goto($loc)`,
		`cy.get($locator)` => `page.locator($locator)`,
		`cy.contains($text, $options)` => `await expect(page.getByText($text)).toBeVisible($options)`,
		`cy.get($locator).contains($text).$action()` => `await page.locator($locator, { hasText: $text }).$action()`,
		`cy.get($locator).contains($text)` => `page.locator($locator, { hasText: $text })`,
		`cy.contains($text)` => `await expect(page.getByText($text)).toBeVisible()`,
		`cy.log($log)` => `console.log($log)`,
		`cy.wait($timeout)` => `await page.waitForTimeout($timeout)`,
		`$locator.find($inner)` => `$locator.locator($inner)`,
		`$locator.eq($n)` => `$locator.nth($n)`,
		`$locator.click($opts)` => `await $locator.click($opts)`,
		`$locator.text()` => `await $locator.textContent()`,
		`Cypress.env('$var')` => `process.env.$var`,
		`cy.onlyOn($var === $cond)` => `if ($var !== $cond) {
  test.skip();
}`,
		`cy.$_($selector).each(($locator) => {
            $body
        })` as $loop where {
			$var = `$[locator]s`,
			$loop => `const $var = await page.locator($selector).all();
for (const $locator of $var) {
    $body
}`
		},
		`cy.request({ $opts })` as $req where {
			or {
				$opts <: contains pair(key=`method`, value=`"$method"`),
				$method = `get`
			},
			$opts <: contains pair(key=`url`, value=$url),
			$method = lowercase($method),
			$other_opts = [],
			$opts <: some bubble($other_opts) $opt where {
				$opt <: not contains or {
					`method`,
					`url`
				},
				$other_opts += $opt
			},
			$other_opts = join($other_opts, `,`),
			$req => `await request.$method($url, { $other_opts })`
		}
	}
}

pattern convert_cypress_test() {
	or {
		`describe($description, $suite)` => `test.describe($description, $suite)` where {
			$suite <: maybe contains bubble or {
				`before($hook)` => `test.beforeAll(async $hook)`,
				`beforeEach($hook)` => `test.beforeEach(async $hook)`,
				`after($hook)` => `test.afterAll(async $hook)`,
				`afterEach($hook)` => `test.afterEach(async $hook)`
			}
		},
		or {
			`it($description, () => { $body })`,
			`test($description, () => { $body })`
		} => `test($description, async ({ page, request }) => {
            $body
        })`
	}
}

contains bubble or {
	convert_cypress_assertions(),
	convert_cypress_queries()
} where {
	$program <: maybe contains bubble convert_cypress_test(),
	$expect = `expect`,
	$expect <: ensure_import_from(source=`"@playwright/test"`),
	$test = `test`,
	$test <: ensure_import_from(source=`"@playwright/test"`)
}
```

## Converts basic test

```js
describe('A mock test', () => {
  test('works', () => {
    cy.onlyOn(Cypress.env('ENVIRONMENT') === 'local');
    cy.visit('/');
    cy.get('.button').should('exist');
    cy.get('.button').should('contain', 'Hello world');
  });
});
```

```ts
import { expect, test } from '@playwright/test';

test.describe('A mock test', () => {
  test('works', async ({ page, request }) => {
    if (process.env.ENVIRONMENT !== 'local') {
      test.skip();
    }
    await page.goto('/');
    await expect(page.locator('.button')).toBeAttached();
    await expect(page.locator('.button')).toContainText('Hello world');
  });
});
```

## Converts requests

```js
cy.request({
  method: 'POST',
  url: '/submit',
  body: JSON.stringify({
    content: 'Hello world',
  }),
  failOnStatusCode: false,
});
cy.contains('Submitted', { timeout: 10000 });
```

```ts
import { expect, test } from '@playwright/test';

await request.post('/submit', {
  body: JSON.stringify({
    content: 'Hello world',
  }),
  failOnStatusCode: false,
});
await expect(page.getByText('Submitted')).toBeVisible({ timeout: 10000 });
```

## Converts hooks

```js
describe('Grouping', function () {
  before(function () {
    setup();
  });

  afterEach(function () {
    cy.wait(1000);
    teardown();
  });
});
```

```ts
import { expect, test } from '@playwright/test';

test.describe('Grouping', function () {
  test.beforeAll(async function () {
    setup();
  });

  test.afterEach(async function () {
    await page.waitForTimeout(1000);
    teardown();
  });
});
```

## Converts composite queries

```js
describe('Grouping', function () {
  it('my test', async () => {
    cy.get('.header').find('.button').eq(1).click({ force: true });
    cy.get('.sidebar').contains('Files').click();
    cy.get('.header').find('.button').eq(1).should('have.attr', 'disabled');
    cy.get('.button').each((button) => {
      expect(button.text()).to.eql('Submit');
    });
  });
});
```

```ts
import { expect, test } from '@playwright/test';

test.describe('Grouping', function () {
  test('my test', async ({ page, request }) => {
    await page.locator('.header').locator('.button').nth(1).click({ force: true });
    await page.locator('.sidebar', { hasText: 'Files' }).click();
    await expect(page.locator('.header').locator('.button').nth(1)).toHaveAttribute('disabled');
    const buttons = await page.locator('.button').all();
    for (const button of buttons) {
      expect(await button.textContent()).toEqual('Submit');
    }
  });
});
```
