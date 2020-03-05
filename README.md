# eWorkBook

[![Go Report Card](https://goreportcard.com/badge/github.com/parithibang/e-work-book?style=flat-square)](https://goreportcard.com/report/github.com/parithibang/e-work-book)
[![GitHub contributors](https://img.shields.io/github/contributors/parithibang/e-work-book.svg?style=plastic&color=blue)](https://GitHub.com/parithibang/e-work-book/graphs/contributors/)
![Last Commit](https://img.shields.io/github/last-commit/parithibang/e-work-book.svg?style=plastic)

![search_module](/assets/search_module.png "Search Module")

eWorkBook is a project management tool which is used for resource allocation and visualize the project key metrics

## How Do I Install?

### Pre-requisite

If you are wanting to build and develop this, you will need the following items installed.

- Go version 1.12+

### Installation via Git

```bash
git clone git@github.com:parithibang/e-work-book.git $GOPATH/src/e-work-book
cd $GOPATH/src/e-work-book
make install
```

Make sure you set the `GOPATH` environment variable. [Resource](https://github.com/golang/go/wiki/SettingGOPATH) for setting the `GOPATH`.

## Benefits of this tool?

This is a resource allocation management tool which used to keep track of the projects assigned to the employees and how much percentage of work is been assigned to the employees on each project.

These are the core modules

- Users
- Pods
- Projects
- Teams
- Search

#### Users

Users are EHI employees. There are 2 user roles

- Pod Leads
- Employees

This application will be accessible only to Pod leads. Pod lead created or uploads employee data.

#### Pods

EMIS Health India Pods

#### Projects

Projects module used to list all the projects in the overall organisation.

#### Teams

Teams module is a group of employees who work in the projects

#### Search

Search module is used to filter out the user record and percentage of work assigned to each project.
