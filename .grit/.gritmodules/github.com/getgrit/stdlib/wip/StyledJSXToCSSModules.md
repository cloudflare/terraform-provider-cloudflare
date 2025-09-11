---
title: Convert Styled JSX to CSS Modules
---

Extract all Styled JSX from a particular file and move it to CSS Module files.

- If there are multiple components in a given file, we create separate CSS Module file for each one.
- Styles defined as global are currently moved to the same CSS Module file with scope set to global.
- We use variable/component names for exporeted styles to create CSS Module files, and current filename for default exports.
- Currently given the limitation of processing the CSS code snippet, we don't touch styles that have conditions in them and that are evaluated on runtime. eg. `background-color: prop.active ? '#e7e7e7' : '#fff'`
- We use `import { default as {name} }` currently because we only have named imports in `ensureImportFrom`

tags: #good

```grit
pattern UpdateClassName($importModule) {
  JSXElement(children=$children, openingElement=JSXOpeningElement(attributes=[..., `className="$classesRaw"`, ...])) where {
    ensureImportFrom(Identifier(name=s"default as cn"), `'classnames'`),
    $children <: contains bubble($classesRaw, $importModule) `<style $_>{$styles}</style>` where {
      $styles <: contains bubble($classesRaw, $importModule) TemplateElement(value=RawCooked(raw=$css)) where {
        
        // NOTE: This is a workaround for now.
        // Idea is to select all classnames from the style.

        // Replace everything but the classnames.
        $classNames = replaceAll($css, r"(?:\\s*\\{[\\s\\S]*?\\})|(?:\\b(?!\\.)\\w+\\b(?!\\s*\\{))", ""),
        // Remove extra spaces and new lines.
        $classNames = replaceAll($classNames, r"[\\s\\n]+", ""),
        // Split classnames
        $classNames = split("\\.", $classNames),
        
        $classList = [],
        $classNames <: some bubble($classList, $importModule) $class where {
          if (! $class <: "") {
            $classList = [... $classList, raw(s"${importModule}.${class}") ]
          }
        },

        $classesRaw => `cn($classList)`
      }
    }
  }
}

predicate IsGlobalStyle($scope) {
  $scope <: contains `global`
}

predicate CreateCSSModule($styles, $cssFileName, $scope) {
  // Currently this only selects css with no expressions that needs to be evaluated on runtime.
  $styles <: or {
    bubble($scope, $cssFileName) TemplateLiteral(quasis=[$css]) where {
      if (IsGlobalStyle($scope)) {
        // Instead of bundling all globals in a file we scope them as global.
        $scopedCSS = raw(":global{\n" + unparse($css) + "\n}")
      } else {
        $scopedCSS = raw(unparse($css))
      },
      $newFiles = [File(name = $cssFileName, program = Program([$scopedCSS]))]
    },
    bubble($scope, $cssFileName) RawCooked(raw=$css) where {
      if (IsGlobalStyle($scope)) {
        $scopedCSS = raw(":global{\n" + $css + "\n}")
      } else {
        $scopedCSS = raw($css)
      },
      $newFiles = [File(name = $cssFileName, program = Program([$scopedCSS]))]
    }
  }
} 

pattern ExportStyles($cssFile, $importModule) {
  bubble($cssFile, $importModule) `<style $scope>{$styles}</style>` as $jsxMatch where {
    $jsxMatch => .,
    CreateCSSModule($styles, $cssFile, $scope),
    ensureImportFrom(Identifier(name=s"default as ${importModule}"), `$cssFile`)
  }
}

pattern RewriteNamedComponents() {
  bubble `const $compName = ($_) => $body` where {
    $file = join(".", [$compName, "module.css"]),
    $importModule = toLowerCase($compName),
    $body <: contains bubble($file, $compName, $importModule) ExportStyles($file, $importModule),
    $body <: maybe contains bubble($importModule) UpdateClassName($importModule)
  }
}

pattern RewriteDefaultComponents() {
  bubble `export default () => $body` where {
    $file = replaceAll($filename, r"\\.(tsx|js|jsx|ts)$", ".module.css"),
    $body <: contains bubble($file) ExportStyles($file, "styles")
  }
}

pattern RewriteNamedStyleExports() {
    `const $styleName = $body` where {
        $body <: bubble($body, $styleName) TaggedTemplateExpression(tag=$tag, quasi=$styles) where {
            $cssFileName = join(".", [$styleName, "module.css"]),
            ensureImportFrom(Identifier(name=s"default as ${styleName}Styles"), `$cssFileName`),
            CreateCSSModule($styles, $cssFileName, $tag),
            $body => raw(s"${styleName}Styles")
        }
    }
}

pattern RewriteDefaultStyleExports() {
    `export default $body` where {
        $body <: bubble($body) TaggedTemplateExpression(tag=$tag, quasi=$styles) where {
            $cssFileName = replaceAll($filename, r"\\.(tsx|js|jsx|ts)$", ".module.css"),
            ensureImportFrom(Identifier(name=s"default as defaultStyles"), `$cssFileName`),
            CreateCSSModule($styles, $cssFileName, $tag),
            $body => `defaultStyles`
        }
    }
}

or {
  RewriteNamedComponents(),
  RewriteDefaultComponents(),
  RewriteNamedStyleExports(),
  RewriteDefaultStyleExports()
}
```

## Convert locally scoped styled JSX

```javascript
const Button = (props) => (
  <button className="cta large">
    {props.children}
    <style jsx>{`
      .cta {
        padding: 20px;
        background: #eee;
        color: #999;
      }
      .large {
        padding: 50px;
      }
    `}</style>
  </button>
);
```

```javascript
// @filename: test.js
import { default as button } from 'Button.module.css';
import { default as cn } from 'classnames';
const Button = (props) => (
  <button className={cn(button.cta, button.large)}>
    {props.children}

  </button>
);
// @filename: Button.module.css

      .cta {
        padding: 20px;
        background: #eee;
        color: #999;
      }
      .large {
        padding: 50px;
      }
    
```

## Convert globally scoped styled JSX

```javascript
export default () => (
  <div className="container">
    <style jsx global>{`
      body {
        background: red;
      }
      .container {
        background: red;
      }
    `}</style>
  </div>
);
```

```javascript
// @filename: test.js
import { default as styles } from 'test.module.css';
export default () => (
  <div className="container">

  </div>
);
// @filename: test.module.css
:global{

      body {
        background: red;
      }
      .container {
        background: red;
      }
    
}
```

## Convert exported styles

```javascript
import css from "styled-jsx/css";

export const button = css`
  .button {
    color: hotpink;
  }
`;
```

```javascript
// @filename: test.js
import css from "styled-jsx/css";

import { default as buttonStyles } from 'button.module.css';

export const button = buttonStyles;
// @filename: button.module.css

  .button {
    color: hotpink;
  }
```

## Multiple styled jsx in a single file

```javascript
const CTAButton = (props) => (
  <button className="button large">
    {props.children}
    <style jsx>{`
      .button {
        padding: 20px;
        background: #eee;
        color: #999;
      }
      .large {
        padding: 50px;
      }
    `}</style>
  </button>
);

const AppContainer = (props) => (
  <div className="main">
    {props.children}
    <style jsx>{`
      .main {
        padding: 20px;
        background: #eee;
        color: #999;
      }
    `}</style>
  </div>
);
```

```javascript
// @filename: test.js
import { default as ctabutton } from 'CTAButton.module.css';
import { default as cn } from 'classnames';
import { default as appcontainer } from 'AppContainer.module.css';
const CTAButton = (props) => (
  <button className={cn(ctabutton.button, ctabutton.large)}>
    {props.children}

  </button>
);

const AppContainer = (props) => (
  <div className={cn(appcontainer.main)}>
    {props.children}

  </div>
);
// @filename: AppContainer.module.css

      .main {
        padding: 20px;
        background: #eee;
        color: #999;
      }

// @filename: CTAButton.module.css

      .button {
        padding: 20px;
        background: #eee;
        color: #999;
      }
      .large {
        padding: 50px;
      }
    
```
