---
title: Replace React default imports with destructured named imports
tags: [import, react, global]
---

This pattern replaces React default import method references (e.g. `React.ReactNode`) with destructured named imports (`import { ReactNode } from 'react'`).
Running this will also make sure that `React` is imported.

```grit
engine marzano(0.1)
language js(jsx)

`React.$reactImport` where {
	$reactImport <: ensure_import_from(`"react"`)
} => `$reactImport`
```

## Replace method on React default import

Given the following interface with `React` global:

```typescript
import React from 'react';

interface MyComponentProps {
  children: React.ReactNode;
  anotherProp?: boolean;
}
```

The result should import the relevant module:

```typescript
import React from 'react';
import { ReactNode } from 'react';

interface MyComponentProps {
  children: ReactNode;
  anotherProp?: boolean;
}
```

Patterns such as `unused_imports` can be used to clean up the duplicate import.

## Replace React global

Given the following interface with `React` global:

```typescript
interface MyComponentProps {
  children: React.ReactNode;
  anotherProp?: boolean;
}
```

The result should import the relevant module:

```typescript
import { ReactNode } from 'react';

interface MyComponentProps {
  children: ReactNode;
  anotherProp?: boolean;
}
```
