---
title: Upgrade Langfuse to v2
tags: [python, upgrade, langfuse, migration]
---

Upgrade the Langfuse SDK to v2 following [this guide](https://langfuse.com/docs/sdk/python#upgrading-from-v1xx-to-v2xx).

```grit
engine marzano(0.1)
language python

predicate imports_langfuse() {
	$program <: contains or {
		import_from_statement(),
		import_statement()
	} as $import where { $import <: contains `langfuse` }
}

pattern convert_snake_case() {
	maybe contains any {
		`traceId` => `trace_id`,
		`startTime` => `start_time`,
		`endTime` => `end_time`,
		`completionStartTime` => `completion_start_time`,
		`statusMessage` => `status_message`,
		`userId` => `user_id`,
		`sessionId` => `session_id`,
		`parentObservationId` => `parent_observation_id`,
		`modelParameters` => `model_parameters`
	}
}

pattern convert_pydantic_enum() {
	maybe some `level=$obs_level.$level` => `level="$level"` where {
		$obs_level <: `ObservationLevel`,
		$obs_level <: remove_from_imports()
	}
}

pattern rename_generation_params() {
	maybe some bubble keyword_argument($name) where {
		$name <: or {
			`prompt` => `input`,
			`completion` => `output`
		}
	}
}

pattern prune_langfuse_imports() {
	maybe contains bubble or {
		`InitialGeneration`,
		`CreateGeneration`,
		`InitialScore`,
		`InitialSpan`,
		`CreateScore`,
		`CreateTrace`,
		`CreateSpan`,
		`CreateEvent`,
		`UpdateGeneration`,
		`UpdateSpan`,
		`CreateDatasetItemRequest`,
		`CreateDatasetRequest`,
		`Usage`
	} as $deprecated where { $deprecated <: remove_from_imports() }
}

or {
	or {
		`$langfuse.generation(InitialGeneration($params))` => `$langfuse.generation($params)`,
		`$langfuse.generation(CreateGeneration($params))` => `$langfuse.generation($params)`,
		`$langfuse.generation($params)`
	} where { $params <: rename_generation_params() },
	`$langfuse.score(InitialScore($params))` => `$langfuse.score($params)`,
	`$langfuse.span(InitialSpan($params))` => `$langfuse.span($params)`,
	`$langfuse.score(CreateScore($params))` => `$langfuse.score($params)`,
	`$langfuse.trace(CreateTrace($params))` => `$langfuse.trace($params)`,
	`$langfuse.span(CreateSpan($params))` => `$langfuse.span($params)`,
	`$langfuse.event(CreateEvent($params))` => `$langfuse.event($params)`,
	`$generation.update(UpdateGeneration($params))` => `$generation.update($params)`,
	`$span.update(UpdateSpan($params))` => `$span.update($params)`,
	`$langfuse.create_dataset_item(CreateDatasetItemRequest($params))` => `$langfuse.create_dataset_item($params)`,
	`$langfuse.create_dataset(CreateDatasetRequest($params))` => `$langfuse.create_dataset($params)`,
	`usage=Usage($params)` as $usage where {
		$props = [],
		$params <: some bubble($props) keyword_argument($name, $value) where {
			$props += `"$name": $value`
		},
		$props = join($props, `, `),
		$usage => `usage={ $props }`
	},
	`$langfuse.score($params)`,
	`$langfuse.span($params)`,
	`$langfuse.trace($params)`,
	`$langfuse.event($params)`,
	`$generation.update($params)`,
	`$langfuse.create_dataset_item($params)`,
	`$langfuse.create_dataset($params)`
} as $lf_func where {
	$lf_func <: prune_langfuse_imports(),
	$params <: convert_snake_case(),
	$params <: convert_pydantic_enum(),
	imports_langfuse()
}
```

## Rewrites Pydantic interface argument

```python
from langfuse.model import InitialSpan

langfuse.span(
    InitialSpan(
        name="span",
        startTime=timestamp,
        endTime=timestamp,
        input={"key": "value"},
        output={"key": "value"},
    )
)
```

```python
langfuse.span(
    name="span",
    start_time=timestamp,
    end_time=timestamp,
    input={"key": "value"},
    output={"key": "value"},
)
```

## Renames `prompt` and `completion` to `input` and `output`

```python
from langfuse import *

client.generation(InitialGeneration(
        name="some_generation",
        startTime=start_time,
        endTime=end_time,
        prompt=prompt,
        completion=completion,
        metadata=metadata
    )
)
```

```python
from langfuse import *

client.generation(
    name="some_generation",
    start_time=start_time,
    end_time=end_time,
    input=prompt,
    output=completion,
    metadata=metadata
)
```

## Snake cases parameters without Pydantic

```python
import langfuse

generation = observation.generation(
    name='name',
    prompt=kwargs['messages'],
    startTime=dt.datetime.utcnow(),
)
```

```python
import langfuse

generation = observation.generation(
    name='name',
    input=kwargs['messages'],
    start_time=dt.datetime.utcnow(),
)
```

## Converts Pydantic enum

```python
from langfuse.model import InitialGeneration
from langfuse.api.resources.commons.types.observation_level import ObservationLevel
import langfuse

langfuse.generation(InitialGeneration(level=ObservationLevel.ERROR))
```

```python
import langfuse

langfuse.generation(level="ERROR")
```

## Rewrites nested Pydantic interface

```python
from langfuse.model import InitialGeneration, Usage
import langfuse as lf

generation = lf.generation(
    InitialGeneration(
        name="chatgpt-completion",
        startTime=generationStartTime,
        endTime=datetime.now(),
        model=self.model,
        modelParameters={"temperature": str(temperature)},
        prompt=history,
        completion=response["choices"][0]["message"]["content"],
        usage=Usage(
            promptTokens=50,
            completionTokens=50,
        ),
    )
)
```

```python
import langfuse as lf

generation = lf.generation(name="chatgpt-completion",
    start_time=generationStartTime,
    end_time=datetime.now(),
    model=self.model,
    model_parameters={"temperature": str(temperature)},
    input=history,
    output=response["choices"][0]["message"]["content"],
    usage={"promptTokens": 50, "completionTokens": 50},
)
```

## Does nothing without langfuse import

```python
model.event(
    CreateEvent(
        name="span",
        startTime=timestamp,
        endTime=timestamp,
    )
)
```
