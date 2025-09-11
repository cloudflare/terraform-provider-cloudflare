---
title: Enzyme to React Testing Library
tags: [enzyme, react-testing-library, rtl, alpha, migration, hidden]
---

(Alpha) This pattern is a work in progress and is not yet ready for use.

Note: the full migration is packaged as a workflow. This is just a subcomponent.


```grit
language js

predicate rtl_import_render() {
	$render = `render`,
	$render <: ensure_import_from(source=`"@testing-library/react"`)
}
pattern mount() {
	or {
		`$mount($comp)` as $mountComp where {
			$mount <: or {
				`mount`,
				`shallow`
			},
			rtl_import_render(),
			$mountComp => `render($comp)`
		},
		`$_.render($_)` where { rtl_import_render() },
		`import { $imports } from 'enzyme'` => .
	}
}

pattern simulate_input() {
	`$inputFind.simulate($type, $value)` as $simulate where {
		$fire_event = `fireEvent`,
		$fire_event <: ensure_import_from(source=`"@testing-library/react"`),
		$simulate => `const selector = $inputFindfireEvent.$eventType(selector, { target: { value: $value } });`
	}
}

predicate is_rtl_query_selector($value) {
	$value <: r"^(?:[a-zA-Z0-9_-]*[#.])+[a-zA-Z0-9_.#-]*"
}

predicate rtl_selector_rewrite($value, $locator, $compVar, $selector) {
	if (is_rtl_query_selector($value)) {
		$locator => `querySelector`
	} else if ($value <: r"input\[name=([^\]]+)]"($formField)) {
		$selector => `["textbox", ObjectExpression(properties=[
            ObjectProperty(key=Identifier(name="name"), value=raw($formField))
        ])]`
	} else {
		$screen = `screen`,
		$screen <: ensure_import_from(source=`"@testing-library/react"`),
		$compVar => `screen`,
		$locator => `getByRole`,
		$selector <: or {
			`'h2'` => `'heading'`,
			`'span'` => `'heading'`,
			$_
			// TODO: AI fallback
			// $guessRole = guess(codePrefix="// fix role using HTML tag", fallback=unparse($selector), stop=["function"]),
			// $selector => $guessRole
		}
	}
}

pattern rewrite_selector() {
	`$compVar.$locator($selector)` where {
		$locator <: `find`,
		if ($selector <: string(fragment=$value)) {
			rtl_selector_rewrite($value, $locator, $compVar, $selector)
		} else {
			// If the variable used in the selector has a classname assigned rewrite it
			$program <: contains variable_declaration() as $var where {
				$var <: contains `$selector = $varSelector` where {
					$varSelector <: string(fragment=$value),
					rtl_selector_rewrite($value, $locator, $compVar, $selector)
				}
			}
		}
	}
}

pattern rtl_base_rewrite() {
	or {
		`$_.update()` => .,
		`$_.act()` => .,
		`$textFind.text()` => `$textFind.textContent`,
		`$inputFind.prop('value')` => `$inputFind.value`
	}
}

or {
	mount(),
	rewrite_selector(),
	simulate_input(),
	rtl_base_rewrite()
}
```

## Simple example

More examples: https://github.com/getgrit/js/blob/9abb21e0849cd220091a1f1ad44ed77a10f5d9d1/wip/EnzymeToRTL.md

```js
import { mount } from 'enzyme';
import TestModel from './modal';

describe('Modal', () => {
  describe('render', () => {
    it('should render', () => {
      testObject.render({ showModal: true });
      expect(testObject.component.find('h2').text()).toEqual('Test Modal');
    });

    it('renders header as the first child', () => {
      const header = testObject.component.find('span').at(0);
      expect(header.text()).toEqual('Hello, Header!');
    });
  });
});
```

```js

import { render, screen } from "@testing-library/react";

import TestModel from './modal';

describe('Modal', () => {
  describe('render', () => {
    it('should render', () => {
      testObject.render({ showModal: true });
      expect(screen.getByRole('heading').textContent).toEqual('Test Modal');
    });

    it('renders header as the first child', () => {
      const header = screen.getByRole('heading').at(0);
      expect(header.textContent).toEqual('Hello, Header!');
    });
  });
});
```
