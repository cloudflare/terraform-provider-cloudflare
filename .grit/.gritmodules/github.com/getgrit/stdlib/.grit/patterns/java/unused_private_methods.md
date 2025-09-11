---
title: Unused private methods should be removed
tags: [java]
---

# Unused private methods should be removed

Unused private methods, excepting methods with annotations and special methods overriding Java's default behaviour, constitute dead code and should therefore be removed.


```grit
language java

class_body($declarations) where {
	$declarations <: contains bubble($declarations) {
		method_declaration(name=$unused_method, $modifiers) as $unused_decl where {
			$modifiers <: contains `private`,
			$modifiers <: not contains or {
				marker_annotation(),
				`native`
			},
			$unused_method <: not or {
				`writeObject`,
				`readObject`
			},
			$declarations <: not contains $unused_method until method_declaration($name) where {
				$name <: `$unused_method`
			},
			$unused_decl => .
		}
	}
}
```

## Removes unused non-constructor method

```java
public class Foo implements Serializable
{
  private Foo(){}
  public static void doSomething(){
    Foo foo = new Foo();
  }

  public void sayHi() {
    this.usedPrivateMethod();
  }

  private void usedPrivateMethod() { }

  private void unusedPrivateMethod() { }

  private void anotherUnusedPrivateMethod() { }
}
```

```java
public class Foo implements Serializable
{
  private Foo(){}
  public static void doSomething(){
    Foo foo = new Foo();
  }

  public void sayHi() {
    this.usedPrivateMethod();
  }

  private void usedPrivateMethod() { }


}
```

## Does not remove annotated and override methods

```java
public class Foo implements Serializable
{
  @Annotation
  private void annotatedMethod(){
    this.usedPrivateMethod();
  }
  private void writeObject(ObjectOutputStream s){ }
  private void readObject(ObjectInputStream in){ }
}
```
