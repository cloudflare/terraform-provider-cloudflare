---
private: true
tags: [private]
---

# Test of the move import util

```grit
language js

`sanitizeFilePath` as $s where {
	move_import(`sanitizeFilePath`, `'@getgrit/universal'`)
}
```

## Base case

It already has the import, from the old package.

```js
import { posthog } from '../../services/flags';
import { InternalServiceAccount } from '../../services/auth/sa';
import type { MarzanoResolvedPattern } from '@getgrit/sdk';
import { marzanoResolvedPatternToResolvedGritPattern, sanitizeFilePath } from '@getgrit/sdk';
import path from 'path';
import { ApplicationFailure } from '@temporalio/workflow';
```

```js
import { posthog } from '../../services/flags';
import { sanitizeFilePath } from '@getgrit/universal';

import { InternalServiceAccount } from '../../services/auth/sa';
import type { MarzanoResolvedPattern } from '@getgrit/sdk';
import { marzanoResolvedPatternToResolvedGritPattern } from '@getgrit/sdk';
import path from 'path';
import { ApplicationFailure } from '@temporalio/workflow';
```

## Existing correct import

```js
import { posthog } from '../../services/flags';
import { InternalServiceAccount } from '../../services/auth/sa';
import type { MarzanoResolvedPattern } from '@getgrit/sdk';
import { sanitizeFilePath } from '@getgrit/universal';
import path from 'path';
import { ApplicationFailure } from '@temporalio/workflow';
```

```js
import { posthog } from '../../services/flags';
import { InternalServiceAccount } from '../../services/auth/sa';
import type { MarzanoResolvedPattern } from '@getgrit/sdk';
import { sanitizeFilePath } from '@getgrit/universal';
import path from 'path';
import { ApplicationFailure } from '@temporalio/workflow';
```


## Existing import append

```js
import { posthog } from '../../services/flags';
import { other } from '@getgrit/universal';
import type { MarzanoResolvedPattern } from '@getgrit/sdk';
import { marzanoResolvedPatternToResolvedGritPattern, sanitizeFilePath } from '@getgrit/sdk';
import path from 'path';
import { ApplicationFailure } from '@temporalio/workflow';
```

```js
import { posthog } from '../../services/flags';
import { other, sanitizeFilePath } from '@getgrit/universal';
import type { MarzanoResolvedPattern } from '@getgrit/sdk';
import { marzanoResolvedPatternToResolvedGritPattern } from '@getgrit/sdk';
import path from 'path';
import { ApplicationFailure } from '@temporalio/workflow';
```


## Type-only imports are not confused

If the existing import is type-only, it should not be used.



```js
import { posthog } from '../../services/flags';
import type { other } from '@getgrit/universal';
import type { MarzanoResolvedPattern } from '@getgrit/sdk';
import { marzanoResolvedPatternToResolvedGritPattern, sanitizeFilePath } from '@getgrit/sdk';
import path from 'path';
import { ApplicationFailure } from '@temporalio/workflow';
```

```js
import { posthog } from '../../services/flags';
import { sanitizeFilePath } from '@getgrit/universal';

import type { other } from '@getgrit/universal';
import type { MarzanoResolvedPattern } from '@getgrit/sdk';
import { marzanoResolvedPatternToResolvedGritPattern } from '@getgrit/sdk';
import path from 'path';
import { ApplicationFailure } from '@temporalio/workflow';
```
