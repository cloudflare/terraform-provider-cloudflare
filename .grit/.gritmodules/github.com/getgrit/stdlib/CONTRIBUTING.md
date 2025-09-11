# Contributor Guide

## Getting started

Make sure you have the Grit CLI installed:

```sh
npm install -g @getgrit/launcher
```

To propose changes, fork this repository and open a pull request.

## Adding a new pattern

All patterns require at least one sample validating the functionality. As a result, the best way to develop a
GritQL pattern is often by starting with before and after sample(s) of the code to be transformed. You can
iterate in the [GritQL studio](https://app.grit.io/studio) to develop your pattern.

Once you have a pattern, you can add it to the repository by creating a new `.md` file in the `.grit/patterns`
directory. The file name must be the snake-cased name of the pattern. Kebab case/dashes in `.md` files are not allowed by the GritQL parser. Keep the file name short and
descriptive, and add a concise description as well as any relevant tags. All PRs must contain at least
one sample of before/after code with a descriptive name.

## Testing

Samples can be tested locally using the Grit CLI.

```sh
grit patterns test
```

Once you have your sample(s) passing locally, creating a PR will automatically trigger a CI build running the same tests in a range of environments.

## GritQL conventions

When writing GritQL patterns, follow the conventions in the [GritQL conventions](./grit/patterns/gritql_conventions) directory.

You can confirm you're following the conventions by running the following command:

```sh
grit check
```
