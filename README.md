# reflectctl

An unofficial CLI for [Reflect](https://reflect.run/).

> NOTE: This project is very much under construction and therefore the code is unstable, largely untested and has high churn! Currently, it only implements a subset of the [Reflect API](https://reflect.run/docs/developer-api/documentation/). In most cases there is a clear path forward for adding more capabilities. PRs are welcome!

![CI](https://github.com/employeecycle/reflectctl/actions/workflows/ci.yaml/badge.svg)

# Installation
Clone the reponsitory:

```
git clone https://github.com/employeecycle/reflectctl
```

And then compile from source:

```
make install
```

Test the installation by running the binary:

```
reflectctl
```

Export your Reflect API token in an environment variable:

```
export REFLECT_KEY=<secret value>
```

Optionally, add a configuration file in your home directory:

```
touch ~/.reflectctl.yaml
```

# Usage
View full usage by running the help command:

```
reflectctl --help
```

A common use pattern is first executing all tests associated with a tag, in this case `regression`:

```
reflectctl execute tag regression
```

And then viewing the test status:
```
reflectctl executions status [test ID from execute command above]
```

You can chain these commands together by piping the outout of `execute tag` into the input for `executions status` (and watch the output live) like this:

```
reflectctl execute tag regression | reflectctl executions status -w
```
